package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

var keys = map[glfw.Key]bool{}
var window *glfw.Window
var size = 512

const record = true

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	window, err = glfw.CreateWindow(size, size, "3D", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window.MakeContextCurrent()

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			keys[key] = true
		case glfw.Release:
			keys[key] = false
		}
	})

	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		log.Fatal(err)
	}
	gl.UseProgram(program)

	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(quadData)*4, gl.Ptr(quadData), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	if err := setup(); err != nil {
		panic(err)
	}

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	img := image.NewNRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{size, size}})
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(img.Rect.Size().X),
		int32(img.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix))

	lastFrame := glfw.GetTime()

	images := []*image.NRGBA{}
	frame := 0
	for !window.ShouldClose() {
		currentFrame := glfw.GetTime()

		start := glfw.GetTime()

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(program)

		gl.BindVertexArray(vao)

		if err := render(img, currentFrame-lastFrame); err != nil {
			log.Fatal(err)
		}

		if record {
			imgCopy := image.NewNRGBA(img.Bounds())
			draw.Draw(imgCopy, imgCopy.Bounds(), img, image.ZP, draw.Over)
			images = append(images, imgCopy)
			frame++
			if frame == 200 {
				break
			}
		}

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture)
		gl.TexSubImage2D(
			gl.TEXTURE_2D,
			0,
			0,
			0,
			int32(img.Rect.Size().X),
			int32(img.Rect.Size().Y),
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(img.Pix))

		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		fmt.Println("ms:", (glfw.GetTime()-start)*1000)

		window.SwapBuffers()
		glfw.PollEvents()

		lastFrame = currentFrame
	}

	reduce := func(input map[color.NRGBA]bool, bits int) map[color.NRGBA]bool {
		result := map[color.NRGBA]bool{}
		mask := uint8(^(1<<uint(bits) - 1))
		for c := range input {
			c.R = c.R & mask
			c.G = c.G & mask
			c.B = c.B & mask
			c.A = c.A
			result[c] = true
		}
		return result
	}

	g := gif.GIF{}
	for _, img := range images {
		colors := map[color.NRGBA]bool{}
		for x := 0; x < img.Bounds().Max.X; x++ {
			for y := 0; y < img.Bounds().Max.Y; y++ {
				colors[img.NRGBAAt(x, y)] = true
			}
		}
		pal := []color.Color{color.NRGBA{127, 127, 127, 255}}
		for c := range reduce(colors, 4) {
			pal = append(pal, c)
		}
		pimg := image.NewPaletted(img.Bounds(), pal)
		draw.Draw(pimg, img.Bounds(), img, img.Bounds().Min, draw.Over)
		g.Image = append(g.Image, pimg)
		g.Delay = append(g.Delay, 1)
	}

	f, err := os.Create("out.gif")
	if err != nil {
		log.Fatal(err)
	}
	if err := gif.EncodeAll(f, &g); err != nil {
		log.Fatal(err)
	}
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, errors.New(fmt.Sprintf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

var vertexShader string = `
#version 330

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = vec4(vert, 1);
}
`

var fragmentShader = `
#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}
`

var quadData = []float32{
	//  X, Y, Z, U, V
	-1.0, -1.0, 0, 0.0, 1.0,
	1.0, -1.0, 0, 1.0, 1.0,
	1.0, 1.0, 0, 1.0, 0.0,

	-1.0, -1.0, 0, 0.0, 1.0,
	-1.0, 1.0, 0, 0.0, 0.0,
	1.0, 1.0, 0, 1.0, 0.0,
}
