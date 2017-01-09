package obj

import (
	"bytes"
	. "matrix"
	"testing"
)

var cube = Model{
	Faces: []Face{
		// bottom
		Face{
			Vertices: [3]Vector{
				{-1.0, -1.0, -1.0},
				{1.0, -1.0, -1.0},
				{-1.0, -1.0, 1.0},
			},
			TextureCoords: [3]Vector{
				{0.0, 0.0},
				{1.0, 0.0},
				{0.0, 1.0},
			},
			Normals: [3]Vector{
				{0, -1, 0},
				{0, -1, 0},
				{0, -1, 0},
			},
		},

		Face{
			Vertices: [3]Vector{
				{1.0, -1.0, -1.0},
				{1.0, -1.0, 1.0},
				{-1.0, -1.0, 1.0},
			},
			TextureCoords: [3]Vector{
				{1.0, 0.0},
				{1.0, 1.0},
				{0.0, 1.0},
			},
			Normals: [3]Vector{
				{0, -1, 0},
				{0, -1, 0},
				{0, -1, 0},
			},
		},

		// top
		Face{
			Vertices: [3]Vector{
				{-1.0, 1.0, -1.0},
				{-1.0, 1.0, 1.0},
				{1.0, 1.0, -1.0},
			},
			TextureCoords: [3]Vector{
				{0.0, 0.0},
				{0.0, 1.0},
				{1.0, 0.0},
			},
			Normals: [3]Vector{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
		},

		Face{
			Vertices: [3]Vector{
				{1.0, 1.0, -1.0},
				{-1.0, 1.0, 1.0},
				{1.0, 1.0, 1.0},
			},
			TextureCoords: [3]Vector{
				{1.0, 0.0},
				{0.0, 1.0},
				{1.0, 1.0},
			},
			Normals: [3]Vector{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
		},

		// front
		Face{
			Vertices: [3]Vector{
				{-1.0, -1.0, 1.0},
				{1.0, -1.0, 1.0},
				{-1.0, 1.0, 1.0},
			},
			TextureCoords: [3]Vector{
				{1.0, 0.0},
				{0.0, 0.0},
				{1.0, 1.0},
			},
			Normals: [3]Vector{
				{0, 0, 1},
				{0, 0, 1},
				{0, 0, 1},
			},
		},

		Face{
			Vertices: [3]Vector{
				{1.0, -1.0, 1.0},
				{1.0, 1.0, 1.0},
				{-1.0, 1.0, 1.0},
			},
			TextureCoords: [3]Vector{
				{0.0, 0.0},
				{0.0, 1.0},
				{1.0, 1.0},
			},
			Normals: [3]Vector{
				{0, 0, 1},
				{0, 0, 1},
				{0, 0, 1},
			},
		},

		// back
		Face{
			Vertices: [3]Vector{
				{-1.0, -1.0, -1.0},
				{-1.0, 1.0, -1.0},
				{1.0, -1.0, -1.0},
			},
			TextureCoords: [3]Vector{
				{0.0, 0.0},
				{0.0, 1.0},
				{1.0, 0.0},
			},
			Normals: [3]Vector{
				{0, 0, -1},
				{0, 0, -1},
				{0, 0, -1},
			},
		},

		Face{
			Vertices: [3]Vector{
				{1.0, -1.0, -1.0},
				{-1.0, 1.0, -1.0},
				{1.0, 1.0, -1.0},
			},
			TextureCoords: [3]Vector{
				{1.0, 0.0},
				{0.0, 1.0},
				{1.0, 1.0},
			},
			Normals: [3]Vector{
				{0, 0, -1},
				{0, 0, -1},
				{0, 0, -1},
			},
		},

		// left
		Face{
			Vertices: [3]Vector{
				{-1.0, -1.0, 1.0},
				{-1.0, 1.0, -1.0},
				{-1.0, -1.0, -1.0},
			},
			TextureCoords: [3]Vector{
				{0.0, 1.0},
				{1.0, 0.0},
				{0.0, 0.0},
			},
			Normals: [3]Vector{
				{-1, 0, 0},
				{-1, 0, 0},
				{-1, 0, 0},
			},
		},

		Face{
			Vertices: [3]Vector{
				{-1.0, -1.0, 1.0},
				{-1.0, 1.0, 1.0},
				{-1.0, 1.0, -1.0},
			},
			TextureCoords: [3]Vector{
				{0.0, 1.0},
				{1.0, 1.0},
				{1.0, 0.0},
			},
			Normals: [3]Vector{
				{-1, 0, 0},
				{-1, 0, 0},
				{-1, 0, 0},
			},
		},

		// right
		Face{
			Vertices: [3]Vector{
				{1.0, -1.0, 1.0},
				{1.0, -1.0, -1.0},
				{1.0, 1.0, -1.0},
			},
			TextureCoords: [3]Vector{
				{1.0, 1.0},
				{1.0, 0.0},
				{0.0, 0.0},
			},
			Normals: [3]Vector{
				{1, 0, 0},
				{1, 0, 0},
				{1, 0, 0},
			},
		},

		Face{
			Vertices: [3]Vector{
				{1.0, -1.0, 1.0},
				{1.0, 1.0, -1.0},
				{1.0, 1.0, 1.0},
			},
			TextureCoords: [3]Vector{
				{1.0, 1.0},
				{0.0, 0.0},
				{0.0, 1.0},
			},
			Normals: [3]Vector{
				{1, 0, 0},
				{1, 0, 0},
				{1, 0, 0},
			},
		},
	},
}

func TestLoadStore(t *testing.T) {
	b1, err := cube.MarshalBinary()
	if err != nil {
		t.Errorf("failed to marshal err=%v", err)
	}

	m := Model{}
	if err := m.UnmarshalBinary(b1); err != nil {
		t.Errorf("failed to unmarshal err=%v", err)
	}

	b2, err := m.MarshalBinary()
	if err != nil {
		t.Errorf("failed to marshal err=%v", err)
	}

	if !bytes.Equal(b1, b2) {
		t.Errorf("mismatch")
	}
}
