package M4

import (
	"math"
	"strconv"
	"testing"
)

const epsilon = 1e-7

func TestV4(t *testing.T) {
	scalarPairs := []struct {
		Calculated float32
		Desired    float32
	}{
		{V4{1, 1}.Length(), float32(math.Sqrt(2))},
		{V4{1, 1}.Normalize().Length(), 1.0},
		{V4{1, 1}.Distance(V4{2, 2}), float32(math.Sqrt(2))},
		{V4{2, 3}.DotProduct(V4{3, 4}), 2*3 + 3*4},
	}

	for i, pair := range scalarPairs {
		difference := float64(pair.Desired - pair.Calculated)
		if math.Abs(difference) > epsilon {
			t.Errorf("incorrect result for scalarPair #%d: calculated=%f desired=%f difference=%s", i, pair.Calculated, pair.Desired, strconv.FormatFloat(difference, 'e', -1, 32))
		}
	}

	V4Pairs := []struct {
		Calculated V4
		Desired    V4
	}{
		{V4{1, 1}, V4{1, 1}},
		{V4{1, 1}.MultiplyScalar(3.5), V4{3.5, 3.5}},
		{V4{1, 2}.Multiply(V4{2, 3}), V4{2, 6}},
		{V4{1, 1}.Lerp(V4{2, 2}, 0), V4{1, 1}},
		{V4{1, 1}.Lerp(V4{2, 2}, 0.5), V4{1.5, 1.5}},
		{V4{1, 1}.Lerp(V4{2, 2}, 1), V4{2, 2}},
		{V4{2, 1}.Maximum(V4{1, 2}), V4{2, 2}},
		{V4{2, 1}.Minimum(V4{1, 2}), V4{1, 1}},
	}

	for i, pair := range V4Pairs {
		difference := float64(pair.Desired.Subtract(pair.Calculated).Length())
		if math.Abs(difference) > epsilon {
			t.Errorf("incorrect result for V4Pairs #%d: calculated=%v desired=%v difference=%s", i, pair.Calculated, pair.Desired, strconv.FormatFloat(difference, 'e', -1, 32))
		}
	}

	if (V4{3, -3, 1}).CrossProduct(V4{4, 9, 2}) != (V4{-15, -2, 39}) {
		t.Errorf("cross product fail")
	}
}

func TestM4(t *testing.T) {
	m := M4{
		2, 8, 12, 4,
		11, 3, 5, 9,
		6, 13, 14, 10,
		7, 15, 16, 17,
	}
	if m.Transpose() != (M4{2, 11, 6, 7, 8, 3, 13, 15, 12, 5, 14, 16, 4, 9, 10, 17}) {
		t.Errorf("transpose fail")
	}

	inverse, ok := m.Inverse()
	if !ok {
		t.Errorf("not invertible")
	}

	transposeInverse := inverse.Transpose()

	inverseTranspose, ok := m.Transpose().Inverse()
	if !ok {
		t.Errorf("not invertible")
	}

	// inverse transpose should equal transpose inverse
	if inverseTranspose != transposeInverse {
		t.Errorf("inverse transpose fail")
	}
}

func TestQ4(t *testing.T) {
	basicallyZero := func(v float32) bool {
		return -0.0000001 < v && v < 0.0000001
	}

	{
		angle := float32(math.Pi / 2.0)
		axis := V3{1, 1, 1}.Normalize()
		q := NewQ4(angle, axis)

		if !basicallyZero(q.Length() - 1) {
			t.Errorf("non-unit")
		}

		if q.Angle() != angle {
			t.Errorf("angle fail")
		}
		if q.Axis() != axis {
			t.Errorf("axis fail")
		}

		if IdentityQ4.Multiply(q) != q {
			t.Errorf("identity fail")
		}

		r := q.Multiply(q.Inverse())
		if r[3] != 1 || r.Length() != 1 {
			t.Errorf("inverse fail")
		}
	}

	{
		q := NewQ4(float32(math.Pi/2.0), V3{0, 0, 1})
		v := V4{1, 0, 0}

		if !basicallyZero(q.Rotate(v).Subtract(V4{0, 1, 0}).Length()) {
			t.Errorf("rotate fail")
		}
	}
}
