package M4

import (
	"fmt"
	"math"
)

type V2 [2]float32
type V3 [3]float32
type V4 [4]float32
type M3 [9]float32
type M4 [16]float32
type Q4 [4]float32

var IdentityM3 = M3{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1,
}

var IdentityM4 = M4{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
}

var IdentityQ4 = Q4{0, 0, 0, 1}

func (v V2) String() string {
	return fmt.Sprintf("[\n\t%0.2f\n\t%0.2f\n]", v[0], v[1])
}

func (v V2) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1])))
}

func (v V2) Negate() V2 {
	return V2{-v[0], -v[1]}
}

func (v V2) Normalize() V2 {
	return v.DivideScalar(v.Length())
}

func (vl V2) AddScalar(s float32) V2 {
	return V2{vl[0] + s, vl[1] + s}
}

func (vl V2) SubtractScalar(s float32) V2 {
	return V2{vl[0] - s, vl[1] - s}
}

func (vl V2) MultiplyScalar(s float32) V2 {
	return V2{vl[0] * s, vl[1] * s}
}

func (vl V2) DivideScalar(s float32) V2 {
	return V2{vl[0] / s, vl[1] / s}
}

func (vl V2) Distance(vr V2) float32 {
	return vr.Subtract(vl).Length()
}

func (vl V2) Add(vr V2) V2 {
	return V2{vl[0] + vr[0], vl[1] + vr[1]}
}

func (vl V2) Subtract(vr V2) V2 {
	return V2{vl[0] - vr[0], vl[1] - vr[1]}
}

func (vl V2) Multiply(vr V2) V2 {
	return V2{vl[0] * vr[0], vl[1] * vr[1]}
}

func (vl V2) Divide(vr V2) V2 {
	return V2{vl[0] / vr[0], vl[1] / vr[1]}
}

func (vl V2) DotProduct(vr V2) float32 {
	return vl[0]*vr[0] + vl[1]*vr[1]
}

func (start V2) Lerp(end V2, t float32) V2 {
	return V2{
		start[0]*(1.0-t) + end[0]*t,
		start[1]*(1.0-t) + end[1]*t,
	}
}

func (vl V2) Project(vr V2) V2 {
	unit := vr.Normalize()
	magnitude := vl.DotProduct(unit)
	return unit.MultiplyScalar(magnitude)
}

func (vl V2) Maximum(vr V2) V2 {
	v := V2{}

	if vl[0] > vr[0] {
		v[0] = vl[0]
	} else {
		v[0] = vr[0]
	}

	if vl[1] > vr[1] {
		v[1] = vl[1]
	} else {
		v[1] = vr[1]
	}

	return v
}

func (vl V2) Minimum(vr V2) V2 {
	v := V2{}

	if vl[0] < vr[0] {
		v[0] = vl[0]
	} else {
		v[0] = vr[0]
	}

	if vl[1] < vr[1] {
		v[1] = vl[1]
	} else {
		v[1] = vr[1]
	}

	return v
}

func (v V3) String() string {
	return fmt.Sprintf("[\n\t%0.2f\n\t%0.2f\n\t%0.2f\n\n]", v[0], v[1], v[2])
}

func (v V3) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
}

func (v V3) Negate() V3 {
	return V3{-v[0], -v[1], -v[2]}
}

func (v V3) Normalize() V3 {
	return v.DivideScalar(v.Length())
}

func (vl V3) AddScalar(s float32) V3 {
	return V3{vl[0] + s, vl[1] + s, vl[2] + s}
}

func (vl V3) SubtractScalar(s float32) V3 {
	return V3{vl[0] - s, vl[1] - s, vl[2] - s}
}

func (vl V3) MultiplyScalar(s float32) V3 {
	return V3{vl[0] * s, vl[1] * s, vl[2] * s}
}

func (vl V3) DivideScalar(s float32) V3 {
	return V3{vl[0] / s, vl[1] / s, vl[2] / s}
}

func (vl V3) Distance(vr V3) float32 {
	return vr.Subtract(vl).Length()
}

func (vl V3) Add(vr V3) V3 {
	return V3{vl[0] + vr[0], vl[1] + vr[1], vl[2] + vr[2]}
}

func (vl V3) Subtract(vr V3) V3 {
	return V3{vl[0] - vr[0], vl[1] - vr[1], vl[2] - vr[2]}
}

func (vl V3) Multiply(vr V3) V3 {
	return V3{vl[0] * vr[0], vl[1] * vr[1], vl[2] * vr[2]}
}

func (vl V3) Divide(vr V3) V3 {
	return V3{vl[0] / vr[0], vl[1] / vr[1], vl[2] / vr[2]}
}

func (vl V3) DotProduct(vr V3) float32 {
	return vl[0]*vr[0] + vl[1]*vr[1] + vl[2]*vr[2]
}

func (vl V3) CrossProduct(vr V3) V3 {
	return V3{vl[1]*vr[2] - vl[2]*vr[1], vl[2]*vr[0] - vl[0]*vr[2], vl[0]*vr[1] - vl[1]*vr[0]}
}

func (start V3) Lerp(end V3, t float32) V3 {
	return V3{
		start[0]*(1.0-t) + end[0]*t,
		start[1]*(1.0-t) + end[1]*t,
		start[2]*(1.0-t) + end[2]*t,
	}
}

func (vl V3) Project(vr V3) V3 {
	unit := vr.Normalize()
	magnitude := vl.DotProduct(unit)
	return unit.MultiplyScalar(magnitude)
}

func (vl V3) Maximum(vr V3) V3 {
	v := V3{}

	if vl[0] > vr[0] {
		v[0] = vl[0]
	} else {
		v[0] = vr[0]
	}

	if vl[1] > vr[1] {
		v[1] = vl[1]
	} else {
		v[1] = vr[1]
	}

	if vl[2] > vr[2] {
		v[2] = vl[2]
	} else {
		v[2] = vr[2]
	}

	return v
}

func (vl V3) Minimum(vr V3) V3 {
	v := V3{}

	if vl[0] < vr[0] {
		v[0] = vl[0]
	} else {
		v[0] = vr[0]
	}

	if vl[1] < vr[1] {
		v[1] = vl[1]
	} else {
		v[1] = vr[1]
	}

	if vl[2] < vr[2] {
		v[2] = vl[2]
	} else {
		v[2] = vr[2]
	}

	return v
}

func (v V4) String() string {
	return fmt.Sprintf("[\n\t%0.2f\n\t%0.2f\n\t%0.2f\n\t%0.2f\n]", v[0], v[1], v[2], v[3])
}

func (v V4) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])))
}

func (v V4) Negate() V4 {
	return V4{-v[0], -v[1], -v[2], -v[3]}
}

func (v V4) Normalize() V4 {
	return v.DivideScalar(v.Length())
}

func (vl V4) AddScalar(s float32) V4 {
	return V4{vl[0] + s, vl[1] + s, vl[2] + s, vl[3] + s}
}

func (vl V4) SubtractScalar(s float32) V4 {
	return V4{vl[0] - s, vl[1] - s, vl[2] - s, vl[3] - s}
}

func (vl V4) MultiplyScalar(s float32) V4 {
	return V4{vl[0] * s, vl[1] * s, vl[2] * s, vl[3] * s}
}

func (vl V4) DivideScalar(s float32) V4 {
	return V4{vl[0] / s, vl[1] / s, vl[2] / s, vl[3] / s}
}

func (vl V4) Distance(vr V4) float32 {
	return vr.Subtract(vl).Length()
}

func (vl V4) Add(vr V4) V4 {
	return V4{vl[0] + vr[0], vl[1] + vr[1], vl[2] + vr[2], vl[3] + vr[3]}
}

func (vl V4) Subtract(vr V4) V4 {
	return V4{vl[0] - vr[0], vl[1] - vr[1], vl[2] - vr[2], vl[3] - vr[3]}
}

func (vl V4) Multiply(vr V4) V4 {
	return V4{vl[0] * vr[0], vl[1] * vr[1], vl[2] * vr[2], vl[3] * vr[3]}
}

func (vl V4) Divide(vr V4) V4 {
	return V4{vl[0] / vr[0], vl[1] / vr[1], vl[2] / vr[2], vl[3] / vr[3]}
}

func (vl V4) DotProduct(vr V4) float32 {
	return vl[0]*vr[0] + vl[1]*vr[1] + vl[2]*vr[2] + vl[3]*vr[3]
}

func (vl V4) CrossProduct(vr V4) V4 {
	return V4{vl[1]*vr[2] - vl[2]*vr[1], vl[2]*vr[0] - vl[0]*vr[2], vl[0]*vr[1] - vl[1]*vr[0], 0}
}

func (start V4) Lerp(end V4, t float32) V4 {
	return V4{
		start[0]*(1.0-t) + end[0]*t,
		start[1]*(1.0-t) + end[1]*t,
		start[2]*(1.0-t) + end[2]*t,
		start[3]*(1.0-t) + end[3]*t,
	}
}

func (vl V4) Project(vr V4) V4 {
	unit := vr.Normalize()
	magnitude := vl.DotProduct(unit)
	return unit.MultiplyScalar(magnitude)
}

func (vl V4) Maximum(vr V4) V4 {
	v := V4{}

	if vl[0] > vr[0] {
		v[0] = vl[0]
	} else {
		v[0] = vr[0]
	}

	if vl[1] > vr[1] {
		v[1] = vl[1]
	} else {
		v[1] = vr[1]
	}

	if vl[2] > vr[2] {
		v[2] = vl[2]
	} else {
		v[2] = vr[2]
	}

	if vl[3] > vr[3] {
		v[3] = vl[3]
	} else {
		v[3] = vr[3]
	}

	return v
}

func (vl V4) Minimum(vr V4) V4 {
	v := V4{}

	if vl[0] < vr[0] {
		v[0] = vl[0]
	} else {
		v[0] = vr[0]
	}

	if vl[1] < vr[1] {
		v[1] = vl[1]
	} else {
		v[1] = vr[1]
	}

	if vl[2] < vr[2] {
		v[2] = vl[2]
	} else {
		v[2] = vr[2]
	}

	if vl[3] < vr[3] {
		v[3] = vl[3]
	} else {
		v[3] = vr[3]
	}

	return v
}

func (ml M3) String() string {
	return fmt.Sprintf(`[
	%0.2f, %0.2f, %0.2f,
	%0.2f, %0.2f, %0.2f,
	%0.2f, %0.2f, %0.2f,
]`, ml[0], ml[3], ml[6],
		ml[1], ml[4], ml[7],
		ml[2], ml[5], ml[8])
}

func (m M3) Inverse() (M3, bool) {
	tmp := M3{
		m[4]*m[8] - m[7]*m[5],
		m[2]*m[7] - m[8]*m[1],
		m[1]*m[5] - m[4]*m[2],
		m[5]*m[6] - m[8]*m[3],
		m[0]*m[8] - m[6]*m[2],
		m[2]*m[3] - m[5]*m[0],
		m[3]*m[7] - m[6]*m[4],
		m[1]*m[6] - m[7]*m[0],
		m[0]*m[4] - m[3]*m[1],
	}

	det := m[0]*tmp[0] + m[1]*tmp[3] + m[2]*tmp[6]

	if det == 0 {
		return M3{}, false
	}

	for i := 0; i < 9; i++ {
		tmp[i] = tmp[i] * 1.0 / det
	}

	return tmp, true
}

func (m M3) Transpose() M3 {
	return M3{
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8],
	}
}

func (m M3) InverseTranspose() (M3, bool) {
	inverse, ok := m.Inverse()
	if !ok {
		return M3{}, false
	}
	return inverse.Transpose(), true
}

func (ml M3) Multiply(mr M3) M3 {
	return M3{
		ml[0]*mr[0] + ml[3]*mr[1] + ml[6]*mr[2],
		ml[1]*mr[0] + ml[4]*mr[1] + ml[7]*mr[2],
		ml[2]*mr[0] + ml[5]*mr[1] + ml[8]*mr[2],
		ml[0]*mr[3] + ml[3]*mr[4] + ml[6]*mr[5],
		ml[1]*mr[3] + ml[4]*mr[4] + ml[7]*mr[5],
		ml[2]*mr[3] + ml[5]*mr[4] + ml[8]*mr[5],
		ml[0]*mr[6] + ml[3]*mr[7] + ml[6]*mr[8],
		ml[1]*mr[6] + ml[4]*mr[7] + ml[7]*mr[8],
		ml[2]*mr[6] + ml[5]*mr[7] + ml[8]*mr[8],
	}
}

func (m M3) Translate(v V2) M3 {
	return m.Multiply(M3{
		1, 0, 0,
		0, 1, 0,
		v[0], v[1], 1,
	})
}

func (m M3) Rotate(radians float32) M3 {
	return m.Multiply(M3{
		float32(math.Cos(float64(radians))), float32(math.Sin(float64(radians))), 0,
		float32(-math.Sin(float64(radians))), float32(math.Cos(float64(radians))), 0,
		0, 0, 1,
	})
}

func (m M3) Scale(v V2) M3 {
	return m.Multiply(
		M3{
			v[0], 0, 0,
			0, v[1], 0,
			0, 0, 1,
		})
}

func (ml M3) Add(mr M3) M3 {
	return M3{
		ml[0] + mr[0],
		ml[1] + mr[1],
		ml[2] + mr[2],
		ml[3] + mr[3],
		ml[4] + mr[4],
		ml[5] + mr[5],
		ml[6] + mr[6],
		ml[7] + mr[7],
		ml[8] + mr[8],
	}
}

func (ml M3) Subtract(mr M3) M3 {
	return M3{
		ml[0] - mr[0],
		ml[1] - mr[1],
		ml[2] - mr[2],
		ml[3] - mr[3],
		ml[4] - mr[4],
		ml[5] - mr[5],
		ml[6] - mr[6],
		ml[7] - mr[7],
		ml[8] - mr[8],
	}
}

func (ml M3) MultiplyV2(vr V2) V2 {
	return V2{
		vr[0]*ml[0] + vr[1]*ml[3] + 1*ml[6],
		vr[0]*ml[1] + vr[1]*ml[4] + 1*ml[7],
	}
}

func (ml M3) MultiplyV3(vr V3) V3 {
	return V3{
		vr[0]*ml[0] + vr[1]*ml[3] + vr[2]*ml[6],
		vr[0]*ml[1] + vr[1]*ml[4] + vr[2]*ml[7],
		vr[0]*ml[2] + vr[1]*ml[5] + vr[2]*ml[8],
	}
}

func (ml M4) String() string {
	return fmt.Sprintf(`[
	%0.2f, %0.2f, %0.2f, %0.2f,
	%0.2f, %0.2f, %0.2f, %0.2f,
	%0.2f, %0.2f, %0.2f, %0.2f,
	%0.2f, %0.2f, %0.2f, %0.2f,
]`, ml[0], ml[4], ml[8], ml[12],
		ml[1], ml[5], ml[9], ml[13],
		ml[2], ml[6], ml[10], ml[14],
		ml[3], ml[7], ml[11], ml[15])
}

func (m M4) Inverse() (M4, bool) {
	// http://stackoverflow.com/questions/1148309/inverting-a-4x4-M4
	tmp := M4{}

	tmp[0] = m[5]*m[10]*m[15] -
		m[5]*m[11]*m[14] -
		m[9]*m[6]*m[15] +
		m[9]*m[7]*m[14] +
		m[13]*m[6]*m[11] -
		m[13]*m[7]*m[10]

	tmp[4] = -m[4]*m[10]*m[15] +
		m[4]*m[11]*m[14] +
		m[8]*m[6]*m[15] -
		m[8]*m[7]*m[14] -
		m[12]*m[6]*m[11] +
		m[12]*m[7]*m[10]

	tmp[8] = m[4]*m[9]*m[15] -
		m[4]*m[11]*m[13] -
		m[8]*m[5]*m[15] +
		m[8]*m[7]*m[13] +
		m[12]*m[5]*m[11] -
		m[12]*m[7]*m[9]

	tmp[12] = -m[4]*m[9]*m[14] +
		m[4]*m[10]*m[13] +
		m[8]*m[5]*m[14] -
		m[8]*m[6]*m[13] -
		m[12]*m[5]*m[10] +
		m[12]*m[6]*m[9]

	tmp[1] = -m[1]*m[10]*m[15] +
		m[1]*m[11]*m[14] +
		m[9]*m[2]*m[15] -
		m[9]*m[3]*m[14] -
		m[13]*m[2]*m[11] +
		m[13]*m[3]*m[10]

	tmp[5] = m[0]*m[10]*m[15] -
		m[0]*m[11]*m[14] -
		m[8]*m[2]*m[15] +
		m[8]*m[3]*m[14] +
		m[12]*m[2]*m[11] -
		m[12]*m[3]*m[10]

	tmp[9] = -m[0]*m[9]*m[15] +
		m[0]*m[11]*m[13] +
		m[8]*m[1]*m[15] -
		m[8]*m[3]*m[13] -
		m[12]*m[1]*m[11] +
		m[12]*m[3]*m[9]

	tmp[13] = m[0]*m[9]*m[14] -
		m[0]*m[10]*m[13] -
		m[8]*m[1]*m[14] +
		m[8]*m[2]*m[13] +
		m[12]*m[1]*m[10] -
		m[12]*m[2]*m[9]

	tmp[2] = m[1]*m[6]*m[15] -
		m[1]*m[7]*m[14] -
		m[5]*m[2]*m[15] +
		m[5]*m[3]*m[14] +
		m[13]*m[2]*m[7] -
		m[13]*m[3]*m[6]

	tmp[6] = -m[0]*m[6]*m[15] +
		m[0]*m[7]*m[14] +
		m[4]*m[2]*m[15] -
		m[4]*m[3]*m[14] -
		m[12]*m[2]*m[7] +
		m[12]*m[3]*m[6]

	tmp[10] = m[0]*m[5]*m[15] -
		m[0]*m[7]*m[13] -
		m[4]*m[1]*m[15] +
		m[4]*m[3]*m[13] +
		m[12]*m[1]*m[7] -
		m[12]*m[3]*m[5]

	tmp[14] = -m[0]*m[5]*m[14] +
		m[0]*m[6]*m[13] +
		m[4]*m[1]*m[14] -
		m[4]*m[2]*m[13] -
		m[12]*m[1]*m[6] +
		m[12]*m[2]*m[5]

	tmp[3] = -m[1]*m[6]*m[11] +
		m[1]*m[7]*m[10] +
		m[5]*m[2]*m[11] -
		m[5]*m[3]*m[10] -
		m[9]*m[2]*m[7] +
		m[9]*m[3]*m[6]

	tmp[7] = m[0]*m[6]*m[11] -
		m[0]*m[7]*m[10] -
		m[4]*m[2]*m[11] +
		m[4]*m[3]*m[10] +
		m[8]*m[2]*m[7] -
		m[8]*m[3]*m[6]

	tmp[11] = -m[0]*m[5]*m[11] +
		m[0]*m[7]*m[9] +
		m[4]*m[1]*m[11] -
		m[4]*m[3]*m[9] -
		m[8]*m[1]*m[7] +
		m[8]*m[3]*m[5]

	tmp[15] = m[0]*m[5]*m[10] -
		m[0]*m[6]*m[9] -
		m[4]*m[1]*m[10] +
		m[4]*m[2]*m[9] +
		m[8]*m[1]*m[6] -
		m[8]*m[2]*m[5]

	det := m[0]*tmp[0] + m[1]*tmp[4] + m[2]*tmp[8] + m[3]*tmp[12]

	if det == 0 {
		return M4{}, false
	}

	for i := 0; i < 16; i++ {
		tmp[i] = tmp[i] * 1.0 / det
	}

	return tmp, true
}

func (m M4) Transpose() M4 {
	return M4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

func (m M4) InverseTranspose() (M4, bool) {
	inverse, ok := m.Inverse()
	if !ok {
		return M4{}, false
	}
	return inverse.Transpose(), true
}

func (m M4) ProjectOrthographic(left, right, bottom, top, near, far float32) M4 {
	return m.Scale(V3{2 / (right - left), 2 / (top - bottom), -2 / (far - near)}).Translate(V3{-(right + left) / 2, -(top + bottom) / 2, (far + near) / 2})
}

func (m M4) ProjectPerspective(fovRadians, aspect, near, far float32) M4 {
	scale := float32(1.0 / math.Tan(float64(fovRadians)/2.0))
	return m.Multiply(M4{
		scale / aspect, 0, 0, 0,
		0, scale, 0, 0,
		0, 0, -(far + near) / (far - near), -1,
		0, 0, -2 * near * far / (far - near), 0,
	})
}

func (ml M4) Multiply(mr M4) M4 {
	return M4{
		mr[0]*ml[0] + mr[1]*ml[4] + mr[2]*ml[8] + mr[3]*ml[12],
		mr[0]*ml[1] + mr[1]*ml[5] + mr[2]*ml[9] + mr[3]*ml[13],
		mr[0]*ml[2] + mr[1]*ml[6] + mr[2]*ml[10] + mr[3]*ml[14],
		mr[0]*ml[3] + mr[1]*ml[7] + mr[2]*ml[11] + mr[3]*ml[15],
		mr[4]*ml[0] + mr[5]*ml[4] + mr[6]*ml[8] + mr[7]*ml[12],
		mr[4]*ml[1] + mr[5]*ml[5] + mr[6]*ml[9] + mr[7]*ml[13],
		mr[4]*ml[2] + mr[5]*ml[6] + mr[6]*ml[10] + mr[7]*ml[14],
		mr[4]*ml[3] + mr[5]*ml[7] + mr[6]*ml[11] + mr[7]*ml[15],
		mr[8]*ml[0] + mr[9]*ml[4] + mr[10]*ml[8] + mr[11]*ml[12],
		mr[8]*ml[1] + mr[9]*ml[5] + mr[10]*ml[9] + mr[11]*ml[13],
		mr[8]*ml[2] + mr[9]*ml[6] + mr[10]*ml[10] + mr[11]*ml[14],
		mr[8]*ml[3] + mr[9]*ml[7] + mr[10]*ml[11] + mr[11]*ml[15],
		mr[12]*ml[0] + mr[13]*ml[4] + mr[14]*ml[8] + mr[15]*ml[12],
		mr[12]*ml[1] + mr[13]*ml[5] + mr[14]*ml[9] + mr[15]*ml[13],
		mr[12]*ml[2] + mr[13]*ml[6] + mr[14]*ml[10] + mr[15]*ml[14],
		mr[12]*ml[3] + mr[13]*ml[7] + mr[14]*ml[11] + mr[15]*ml[15],
	}
}

func (m M4) Translate(v V3) M4 {
	return m.Multiply(M4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		v[0], v[1], v[2], 1,
	})
}

func (m M4) Rotate(radians float32, axis V3) M4 {
	x := axis[0]
	y := axis[1]
	z := axis[2]

	c := float32(math.Cos(float64(radians)))
	s := float32(math.Sin(float64(radians)))

	m00 := c + x*x*(1-c)
	m01 := x*y*(1-c) - z*s
	m02 := x*z*(1-c) + y*s
	ml0 := y*x*(1-c) + z*s
	ml1 := c + y*y*(1-c)
	ml2 := y*z*(1-c) - x*s
	mr0 := z*x*(1-c) - y*s
	mr1 := z*y*(1-c) + x*s
	mr2 := c + z*z*(1-c)

	return m.Multiply(M4{
		m00, ml0, mr0, 0,
		m01, ml1, mr1, 0,
		m02, ml2, mr2, 0,
		0, 0, 0, 1,
	})
}

func (m M4) RotateX(radians float32) M4 {
	return m.Rotate(radians, V3{1, 0, 0})
}

func (m M4) RotateY(radians float32) M4 {
	return m.Rotate(radians, V3{0, 1, 0})
}

func (m M4) RotateZ(radians float32) M4 {
	return m.Rotate(radians, V3{0, 0, 1})
}

func (m M4) Scale(v V3) M4 {
	return m.Multiply(
		M4{
			v[0], 0, 0, 0,
			0, v[1], 0, 0,
			0, 0, v[2], 0,
			0, 0, 0, 1,
		})
}

func (ml M4) Add(mr M4) M4 {
	return M4{
		ml[0] + mr[0],
		ml[1] + mr[1],
		ml[2] + mr[2],
		ml[3] + mr[3],
		ml[4] + mr[4],
		ml[5] + mr[5],
		ml[6] + mr[6],
		ml[7] + mr[7],
		ml[8] + mr[8],
		ml[9] + mr[9],
		ml[10] + mr[10],
		ml[11] + mr[11],
		ml[12] + mr[12],
		ml[13] + mr[13],
		ml[14] + mr[14],
		ml[15] + mr[15],
	}
}

func (ml M4) Subtract(mr M4) M4 {
	return M4{
		ml[0] - mr[0],
		ml[1] - mr[1],
		ml[2] - mr[2],
		ml[3] - mr[3],
		ml[4] - mr[4],
		ml[5] - mr[5],
		ml[6] - mr[6],
		ml[7] - mr[7],
		ml[8] - mr[8],
		ml[9] - mr[9],
		ml[10] - mr[10],
		ml[11] - mr[11],
		ml[12] - mr[12],
		ml[13] - mr[13],
		ml[14] - mr[14],
		ml[15] - mr[15],
	}
}

func (ml M4) MultiplyV4(vl V4) V4 {
	return V4{
		vl[0]*ml[0] + vl[1]*ml[4] + vl[2]*ml[8] + vl[3]*ml[12],
		vl[0]*ml[1] + vl[1]*ml[5] + vl[2]*ml[9] + vl[3]*ml[13],
		vl[0]*ml[2] + vl[1]*ml[6] + vl[2]*ml[10] + vl[3]*ml[14],
		vl[0]*ml[3] + vl[1]*ml[7] + vl[2]*ml[11] + vl[3]*ml[15],
	}
}

func NewQ4(radians float32, axis V3) Q4 {
	halfAngle := float64(radians) / 2.0
	scale := float32(math.Sin(halfAngle))
	return Q4{
		axis[0] * scale,
		axis[1] * scale,
		axis[2] * scale,
		float32(math.Cos(halfAngle)),
	}
}

func (v Q4) String() string {
	return fmt.Sprintf("[%0.2fi %0.2fj %0.2fk %0.2f]", v[0], v[1], v[2], v[3])
}

func (q Q4) Length() float32 {
	return float32(math.Sqrt(float64(q[0]*q[0] + q[1]*q[1] + q[2]*q[2] + q[3]*q[3])))
}

func (q Q4) Normalize() Q4 {
	scale := 1.0 / float32(math.Sqrt(float64(q[0]*q[0]+q[1]*q[1]+q[2]*q[2]+q[3]*q[3])))
	return Q4{
		q[0] * scale,
		q[1] * scale,
		q[2] * scale,
		q[3] * scale,
	}
}

func (q Q4) Conjugate() Q4 {
	return Q4{
		-q[0],
		-q[1],
		-q[2],
		q[3],
	}
}

func (q Q4) Inverse() Q4 {
	scale := 1.0 / (q[0]*q[0] + q[1]*q[1] + q[2]*q[2] + q[3]*q[3])
	return Q4{
		-q[0] * scale,
		-q[1] * scale,
		-q[2] * scale,
		q[3] * scale,
	}
}

func (q Q4) Angle() float32 {
	return float32(math.Acos(float64(q[3])) * 2.0)
}

func (q Q4) Axis() V3 {
	halfAngle := math.Acos(float64(q[3]))
	scale := float32(1.0 / math.Sin(halfAngle))
	return V3{
		q[0] * scale,
		q[1] * scale,
		q[2] * scale,
	}
}

func (ql Q4) Multiply(qr Q4) Q4 {
	return Q4{
		ql[3]*qr[0] + ql[0]*qr[3] + ql[1]*qr[2] - ql[2]*qr[1],
		ql[3]*qr[1] - ql[0]*qr[2] + ql[1]*qr[3] + ql[2]*qr[0],
		ql[3]*qr[2] + ql[0]*qr[1] - ql[1]*qr[0] + ql[2]*qr[3],
		ql[3]*qr[3] - ql[0]*qr[0] - ql[1]*qr[1] - ql[2]*qr[2],
	}
}

func (ql Q4) Add(qr Q4) Q4 {
	return Q4{
		ql[0] + qr[0],
		ql[1] + qr[1],
		ql[2] + qr[2],
		ql[3] + qr[3],
	}
}

func (ql Q4) Subtract(qr Q4) Q4 {
	return Q4{
		ql[0] - qr[0],
		ql[1] - qr[1],
		ql[2] - qr[2],
		ql[3] - qr[3],
	}
}

func (q Q4) Rotate(v V4) V4 {
	return V4{
		(1-2*q[1]*q[1]-2*q[2]*q[2])*v[0] + (2*q[0]*q[1]-2*q[3]*q[2])*v[1] + (2*q[0]*q[2]+2*q[3]*q[1])*v[2],
		(2*q[0]*q[1]+2*q[3]*q[2])*v[0] + (1-2*q[0]*q[0]-2*q[2]*q[2])*v[1] + (2*q[1]*q[2]-2*q[3]*q[0])*v[2],
		(2*q[0]*q[2]-2*q[3]*q[1])*v[0] + (2*q[1]*q[2]+2*q[3]*q[0])*v[1] + (1-2*q[0]*q[0]-2*q[1]*q[1])*v[2],
		v[3],
	}
}

func (start Q4) Lerp(end Q4, t float32) Q4 {
	return Q4{
		start[0]*(1.0-t) + end[0]*t,
		start[1]*(1.0-t) + end[1]*t,
		start[2]*(1.0-t) + end[2]*t,
		start[3]*(1.0-t) + end[3]*t,
	}
}

func (start Q4) Nlerp(end Q4, t float32) Q4 {
	return start.Lerp(end, t).Normalize()
}

func (start Q4) Slerp(end Q4, t float32) Q4 {
	angle := math.Acos(float64(start[0]*end[0] + start[1]*end[1] + start[2]*end[2] + start[3]*end[3]))
	startScale := float32(math.Sin(angle*(1.0-float64(t))) / math.Sin(angle))
	endScale := float32(math.Sin(angle*float64(t)) / math.Sin(angle))
	return Q4{
		start[0]*startScale + end[0]*endScale,
		start[1]*startScale + end[1]*endScale,
		start[2]*startScale + end[2]*endScale,
		start[3]*startScale + end[3]*endScale,
	}
}
