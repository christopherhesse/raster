package obj

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"path"
	"strconv"
	"strings"

	"image/jpeg"

	"github.com/ftrvxmtrx/tga"

	. "matrix"
)

type Face struct {
	Vertices      []V4
	TextureCoords []V4
	Normals       []V4
}

type Object struct {
	Faces    []Face
	Material Material
}

type Material struct {
	Ns      float32
	Ni      float32
	D       float32
	Tr      float32
	Tf      V4
	Illum   int
	Ka      V4
	Kd      V4
	Ks      V4
	Ke      V4
	MapKa   image.Image
	MapKd   image.Image
	MapBump image.Image
	Bump    image.Image
}

func parseVector(vals []string) (V4, error) {
	vec := V4{}
	for i, v := range vals {
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return vec, err
		}
		vec[i] = float32(f)
	}
	return vec, nil
}

func parseInts(vals []string) ([]int, error) {
	result := make([]int, len(vals))
	for i, v := range vals {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result[i] = n
	}
	return result, nil
}

func loadMtl(mtlPath string) (map[string]Material, error) {
	materials := map[string]Material{}

	mtlFile, err := os.Open(mtlPath)
	if err != nil {
		return nil, err
	}

	name := ""
	m := Material{}

	scanner := bufio.NewScanner(mtlFile)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		cmd := parts[0]
		args := parts[1:]

		switch cmd {
		case "newmtl":
			if name != "" {
				materials[name] = m
			}
			m = Material{}
			name = parts[1]
		case "Tf", "Ka", "Kd", "Ks", "Ke":
			v, err := parseVector(args)
			if err != nil {
				return nil, err
			}
			switch cmd {
			case "Tf":
				m.Tf = v
			case "Ka":
				m.Ka = v
			case "Kd":
				m.Kd = v
			case "Ks":
				m.Ks = v
			case "Ke":
				m.Ke = v
			}
		case "Ns", "Ni", "d", "Tr":
			f64, err := strconv.ParseFloat(args[0], 32)
			if err != nil {
				return nil, err
			}
			f := float32(f64)
			switch cmd {
			case "Ns":
				m.Ns = f
			case "Ni":
				m.Ni = f
			case "d":
				m.D = f
			case "Tr":
				m.Tr = f
			}
		case "map_Ka", "map_Kd", "map_bump", "bump":
			texPath := path.Join(path.Dir(mtlPath), args[0])
			texFile, err := os.Open(texPath)
			if err != nil {
				return nil, err
			}
			var img image.Image
			if strings.HasSuffix(strings.ToLower(texPath), ".jpg") {
				img, err = jpeg.Decode(texFile)
			} else if strings.HasSuffix(strings.ToLower(texPath), ".tga") {
				img, err = tga.Decode(texFile)
			} else {
				panic("unsupported image format: " + texPath)
			}
			if err != nil {
				return nil, err
			}
			switch cmd {
			case "map_Ka":
				m.MapKa = img
			case "map_Kd":
				m.MapKd = img
			case "map_bump":
				m.MapBump = img
			case "bump":
				m.Bump = img
			}
		case "illum":
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return nil, err
			}
			m.Illum = i
		default:
			fmt.Println("unrecognized line:", line)
		}
	}

	if name != "" {
		materials[name] = m
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return materials, nil
}

func Load(objPath string) ([]Object, error) {
	objFile, err := os.Open(objPath)
	if err != nil {
		return nil, err
	}

	vertices := []V4{}
	textureCoords := []V4{}
	normals := []V4{}
	faces := []Face{}
	materials := map[string]Material{}
	objects := []Object{}
	o := Object{}

	scanner := bufio.NewScanner(objFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		cmd := parts[0]
		args := parts[1:]

		switch cmd {
		case "v", "vt", "vn":
			v, err := parseVector(args)
			if err != nil {
				return nil, err
			}
			switch cmd {
			case "v":
				vertices = append(vertices, v)
			case "vt":
				textureCoords = append(textureCoords, v)
			case "vn":
				normals = append(normals, v)
			}
		case "f":
			f := Face{}
			for _, p := range args {
				idx, err := parseInts(strings.Split(p, "/"))
				if err != nil {
					return nil, err
				}

				// 1-based indexing
				f.Vertices = append(f.Vertices, vertices[idx[0]-1])
				if len(idx) > 1 {
					f.TextureCoords = append(f.TextureCoords, textureCoords[idx[1]-1])
				}
				if len(idx) > 2 {
					f.Normals = append(f.Normals, normals[idx[2]-1])
				}
			}
			faces = append(faces, f)
		case "g":
			if len(faces) > 0 {
				o.Faces = faces
				objects = append(objects, o)
				faces = []Face{}
			}
			o = Object{}
		case "mtllib":
			mtlPath := path.Join(path.Dir(objPath), args[0])
			materials, err = loadMtl(mtlPath)
			if err != nil {
				return nil, err
			}
		case "usemtl":
			o.Material = materials[args[0]]
		default:
			fmt.Println("unrecognized line:", line)
		}
	}

	if len(faces) > 0 {
		o.Faces = faces
		objects = append(objects, o)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return objects, nil
}
