package main

import (
	"image"
	"math"
	"obj"

	"image/color"

	. "matrix"

	"github.com/go-gl/glfw/v3.1/glfw"
)

var (
	triangles []Triangle
)

const (
	interpCount = 5
)

type Triangle struct {
	Vertices      [3]V4
	TextureCoords [3]V4
	Normals       [3]V4
	Texture       *image.NRGBA
}

// find which side of a line a point is on using cross product
func side(x0, y0, x1, y1, px, py float32) bool {
	return (x1-x0)*(py-y0)-(y1-y0)*(px-x0) > 0
}

func clip(v V4) bool {
	return v[0] < -1 || v[0] > 1 || v[1] < -1 || v[1] > 1 || v[2] < -1 || v[2] > 1
}

func HSVToRGB(h, s, v float64) color.NRGBA {
	rgbToColor := func(r, g, b float64) color.NRGBA {
		return color.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), 255}
	}
	var f, p, q, t float64
	if s == 0 {
		return rgbToColor(v, v, v)
	}
	h *= 6 // normally h is 0-360
	i, f := math.Modf(h)
	p = v * (1 - s)
	q = v * (1 - s*f)
	t = v * (1 - s*(1-f))
	switch i {
	case 0:
		return rgbToColor(v, t, p)
	case 1:
		return rgbToColor(q, v, p)
	case 2:
		return rgbToColor(p, v, t)
	case 3:
		return rgbToColor(p, q, v)
	case 4:
		return rgbToColor(t, p, v)
	case 5:
		return rgbToColor(v, p, q)
	}
	return color.NRGBA{255, 255, 255, 255}
}

func min(a, b float32) float32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b float32) float32 {
	if a > b {
		return a
	} else {
		return b
	}
}

var cameraPosition = V4{0, 0, 5}
var cameraYRotation float32 = 0.0
var cameraXRotation float32 = 0.0

var rotation float32 = 0.0

func setup() error {
	// textureFile, err := os.Open("data/cat_diff.tga")
	// if err != nil {
	// 	return err
	// }
	//
	// textureImage, err := tga.Decode(textureFile)
	// if err != nil {
	// 	return err
	// }

	// textureFile, err := os.Open("data/square.png")
	// if err != nil {
	// 	return err
	// }
	//
	// textureImage, err := png.Decode(textureFile)
	// if err != nil {
	// 	return err
	// }
	//
	// texture = textureImage.(*image.NRGBA)

	objPath := "data/cat.obj"
	// objPath := "data/dabrovic-sponza/sponza.obj"

	objects, err := obj.Load(objPath)
	if err != nil {
		return err
	}

	convertedTextures := map[image.Image]*image.NRGBA{nil: nil}
	for _, obj := range objects {
		src := obj.Material.MapKd
		if src == nil {
			continue
		}
		dst := image.NewNRGBA(src.Bounds())

		for x := 0; x < src.Bounds().Max.X; x++ {
			for y := 0; y < src.Bounds().Max.Y; y++ {
				oldColor := src.At(x, y)
				newColor := dst.ColorModel().Convert(oldColor)
				dst.Set(x, y, newColor)
			}
		}

		convertedTextures[src] = dst
	}

	for _, obj := range objects {
		for _, f := range obj.Faces {
			normals := f.Normals[:]
			if len(normals) == 0 {
				// if normals are missing, fill them in
				normal := f.Vertices[1].Subtract(f.Vertices[0]).CrossProduct(f.Vertices[2].Subtract(f.Vertices[0])).Normalize()
				for range f.Vertices {
					normals = append(normals, normal)
				}
			}

			// generate indices to generate triangles from polygons
			// https://www.siggraph.org/education/materials/HyperGraph/scanline/outprims/polygon1.htm
			for i := 0; i < len(f.Vertices)-2; i++ {
				triangle := Triangle{
					Vertices:      [3]V4{f.Vertices[0], f.Vertices[i+1], f.Vertices[i+2]},
					TextureCoords: [3]V4{f.TextureCoords[0], f.TextureCoords[i+1], f.TextureCoords[i+2]},
					Normals:       [3]V4{normals[0], normals[i+1], normals[i+2]},
					Texture:       convertedTextures[obj.Material.MapKd],
				}
				triangles = append(triangles, triangle)
			}
		}
	}

	return nil
}

var frame = 0

func render(img *image.NRGBA, elapsed float64) error {
	size := img.Bounds().Max.X
	fsize := float32(size)

	dx, dy := window.GetCursorPos()
	window.SetCursorPos(0, 0)
	cameraYRotation += float32(dx / 100)
	cameraXRotation += float32(dy / 100)

	if cameraXRotation > math.Pi/2 {
		cameraXRotation = math.Pi / 2
	}
	if cameraXRotation < -math.Pi/2 {
		cameraXRotation = -math.Pi / 2
	}

	delta := V4{}
	if keys[glfw.KeyD] {
		delta = V4{1, 0, 0, 0}
	}

	if keys[glfw.KeyA] {
		delta = V4{-1, 0, 0, 0}
	}

	if keys[glfw.KeyW] {
		delta = V4{0, 0, -1, 0}
	}

	if keys[glfw.KeyS] {
		delta = V4{0, 0, 1, 0}
	}

	if keys[glfw.KeyE] {
		rotation += 0.1
	}

	if keys[glfw.KeyQ] {
		rotation -= 0.1
	}

	if delta != (V4{}) {
		cameraPosition = cameraPosition.Add(IdentityM4.RotateY(-cameraYRotation).RotateX(-cameraXRotation).MultiplyV4(delta.MultiplyScalar(float32(10 * elapsed))))
	}

	// clear the image
	for py := 0; py < size; py++ {
		for px := 0; px < size; px++ {
			offset := py*img.Stride + px*4
			img.Pix[offset] = 127
			img.Pix[offset+1] = 127
			img.Pix[offset+2] = 127
			img.Pix[offset+3] = 255
		}
	}

	// process the vertex data
	projection := IdentityM4.ProjectPerspective(70.0/180.0*math.Pi, 1, 1, 150)
	// projection := Identity.ProjectOrthographic(-5, 5, -5, 5, 5, 15)
	pos := cameraPosition.Negate()
	view := IdentityM4.RotateX(cameraXRotation).RotateY(cameraYRotation).Translate(V3{pos[0], pos[1], pos[2]})
	if record {
		rotation = float32(frame) / 100 * math.Pi
		frame++
	}
	model := IdentityM4.Scale(V3{3, 3, 3}).RotateY(rotation).RotateX(rotation).Translate(V3{-0.1, -0.5, -0.5})
	modelView := view.Multiply(model)
	modelViewProjection := projection.Multiply(modelView)

	normalTransform, ok := modelView.InverseTranspose()
	if !ok {
		panic("failed to invert transform")
	}

	type Datum struct {
		Vertices      [3]V4
		TextureCoords [3]V4
		Normals       [3]V4
		Texture       *image.NRGBA
		Interps       [3][interpCount]float32
	}

	data := make([]Datum, len(triangles))
	for i, t := range triangles {
		data[i].Vertices = t.Vertices
		data[i].TextureCoords = t.TextureCoords
		data[i].Normals = t.Normals
		data[i].Texture = t.Texture
	}

	for i, t := range data {
		for i := 0; i < 3; i++ {
			pos := t.Vertices[i]
			position := V4{pos[0], pos[1], pos[2], 1}
			t.Vertices[i] = modelViewProjection.MultiplyV4(position)

			tex := t.TextureCoords[i]

			norm := t.Normals[i]
			normal := V4{norm[0], norm[1], norm[2], 0}

			// calculate color (interpolated across triangle)
			eye4 := normalTransform.MultiplyV4(normal)
			eye := V4{eye4[0], eye4[1], eye4[2], 0}.Normalize()
			// this is the light direction, not position
			light := view.MultiplyV4(V4{1, 1, 1, 0})
			dotProduct := max(0, eye.DotProduct(light.Normalize()))
			diffuseColor := V3{0.4, 0.4, 1}
			c := diffuseColor.MultiplyScalar(dotProduct)

			t.Interps[i] = [interpCount]float32{c[0], c[1], c[2], tex[0], tex[1]}
		}
		data[i] = t
	}

	// depth buffer so that we can draw triangles in any order and they don't overlap incorrectly
	depthBuf := make([]float32, size*size)
	for i := range depthBuf {
		depthBuf[i] = 1
	}

	for _, d := range data {
		ra := d.Vertices[0]
		rb := d.Vertices[1]
		rc := d.Vertices[2]

		interpA := d.Interps[0]
		interpB := d.Interps[1]
		interpC := d.Interps[2]

		a := ra.MultiplyScalar(1.0 / ra[3])
		b := rb.MultiplyScalar(1.0 / rb[3])
		c := rc.MultiplyScalar(1.0 / rc[3])

		// only show front facing triangles (CCW)
		// https://www.opengl.org/registry/doc/glspec44.core.pdf p.426
		area := a[0]*b[1] - b[0]*a[1] + b[0]*c[1] - c[0]*b[1] + c[0]*a[1] - a[0]*c[1]
		if area <= 0 {
			continue
		}

		// if all points are outside clip-space, skip this triangle
		if clip(a) && clip(b) && clip(c) {
			continue
		}

		minX := min(a[0], min(b[0], c[0]))
		maxX := max(a[0], max(b[0], c[0]))
		minY := min(a[1], min(b[1], c[1]))
		maxY := max(a[1], max(b[1], c[1]))

		// create bounding boxes for triangles
		// it's really slow without bounding boxes
		minPx := int(math.Floor((float64(minX)+1.0)/2.0*float64(size) - 0.5))
		if minPx < 0 {
			minPx = 0
		}
		minPy := int(math.Floor((float64(minY)+1.0)/2.0*float64(size) - 0.5))
		if minPy < 0 {
			minPy = 0
		}
		maxPx := int(math.Ceil((float64(maxX)+1.0)/2.0*float64(size) - 0.5))
		if maxPx >= size {
			maxPx = size - 1
		}
		maxPy := int(math.Ceil((float64(maxY)+1.0)/2.0*float64(size) - 0.5))
		if maxPy >= size {
			maxPy = size - 1
		}

		// generate all pixels that fall into this box
		for py := minPy; py < maxPy; py++ {
			for px := minPx; px < maxPx; px++ {
				// check which pixels have their center inside the triangle
				wx := float32(px) + 0.5
				wy := float32(py) + 0.5

				x := wx/fsize*2.0 - 1.0
				y := wy/fsize*2.0 - 1.0

				s0 := side(a[0], a[1], b[0], b[1], x, y)
				s1 := side(b[0], b[1], c[0], c[1], x, y)
				s2 := side(c[0], c[1], a[0], a[1], x, y)

				if s0 == s1 && s1 == s2 {
					// calculate depth at x,y on the surface of the triangle
					// intersection between line and plane to get depth

					// calculate barycentric coordinates http://en.wikipedia.org/wiki/Barycentric_coordinate_system
					bdenom := (b[1]-c[1])*(a[0]-c[0]) + (c[0]-b[0])*(a[1]-c[1])
					ba := ((b[1]-c[1])*(x-c[0]) + (c[0]-b[0])*(y-c[1])) / bdenom
					bb := ((c[1]-a[1])*(x-c[0]) + (a[0]-c[0])*(y-c[1])) / bdenom
					bc := 1 - ba - bb

					// https://www.opengl.org/registry/doc/glspec44.core.pdf p.427
					depth := ba*a[2] + bb*b[2] + bc*c[2]
					if depth >= -1 && depth <= depthBuf[py*size+px] {
						depthBuf[py*size+px] = depth

						ia := ba / ra[3]
						ib := bb / rb[3]
						ic := bc / rc[3]
						idenom := ia + ib + ic

						interp := [interpCount]float32{}
						for i := range interp {
							interp[i] = (interpA[i]*ia + interpB[i]*ib + interpC[i]*ic) / idenom
						}

						var c color.Color
						if d.Texture == nil {
							c = color.NRGBA{
								uint8(interp[0] * 255),
								uint8(interp[1] * 255),
								uint8(interp[2] * 255),
								255,
							}
						} else {
							// wrap to 0-1
							_, u := math.Modf(float64(interp[3]))
							_, v := math.Modf(float64(interp[4]))
							if u < 0 {
								u = 1 + u
							}
							if v < 0 {
								v = 1 + v
							}
							tx := int(float32(u) * float32(d.Texture.Bounds().Max.X))
							ty := int(float32(v) * float32(d.Texture.Bounds().Max.Y))
							c = d.Texture.At(tx, d.Texture.Bounds().Max.Y-ty)
						}

						img.Set(px, size-py, c) // origin is bottom-left
					}
				}
			}
		}
	}

	// for py := 0; py < size; py++ {
	// 	for px := 0; px < size; px++ {
	// 		d := (depthBuf[py*size+px] + 1.0) / 2.0
	// 		img.Set(px, size-py, HSVToRGB(float64(d), 0.8, 1.0))
	// 		// c := color.NRGBA{
	// 		// 	uint8(d * 255),
	// 		// 	uint8(d * 255),
	// 		// 	uint8(d * 255),
	// 		// 	255,
	// 		// }
	// 		// img.Set(px, py, c)
	// 	}
	// }

	return nil
}
