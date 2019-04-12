package lago

import (
	"math"
	"testing"
)

func TestVectorString(t *testing.T) {
	v := NewVector([]float64{1, 2, 3, 4, 5})
	res := "[ 1 2 3 4 5 ]"
	str := v.String()
	if str != res {
		t.Errorf("Error converting to string: '%s' != '%s'", str, res)
	}
	t.Logf("Vector string: %s", str)
}

func TestVectorSlice(t *testing.T) {
	v := NewVector([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	res := []float64{1, 4}
	slc, err := v.Slice(1, 5, 3)
	if err != nil {
		t.Errorf("Error slicing: %v", err)
	}
	str := slc.Result()
	if notEqual64(str, res) {
		t.Errorf("Sliced string: %v != %v", str, res)
	}
	t.Logf("Sliced vector: %s", slc)
}

func TestVectorDot(t *testing.T) {
	v := NewVector([]float64{1, 2, 3})
	w := NewVector([]float64{4, 5, 6})
	res := 32.0
	resl := v.Dot(w)
	if math.Abs(resl-res) > errprecision {
		t.Errorf("Dot: Expected %f, got %f", res, resl)
	}
	t.Logf("Dot Result: %f", resl)
}

func TestVectorCopy(t *testing.T) {
	v := NewVector([]float64{1, 2, 3, 4, 5})
	w := v.Copy()
	if notEqual64(v.data, w.data) {
		t.Errorf("Copy: %v != %v", v, w)
	}
}
func TestVectorScale(t *testing.T) {
	v := NewVector([]float64{1, 2, 3})
	res := []float64{1.2, 2.4, 3.6}
	v.Scale(1.2)
	if notEqual64(res, v.data) {
		t.Errorf("Scaling: Expected %v, got %v", res, v)
	}
	t.Logf("Scale Result: %v", v)
}

func TestVectorSum(t *testing.T) {
	v := NewVector([]float64{1, 2, 3})
	res := 6.0
	resl := v.Sum()

	if res != resl {
		t.Errorf("Sum: Expected %f, got %f", res, resl)
	}
	t.Logf("Sum Result: %f", resl)

}

func TestVectorNorm(t *testing.T) {
	v := NewVector([]float64{1, 2, 3})
	res := 3.741657387
	resl := v.Norm()

	if notEqual(res, resl) {
		t.Errorf("Norm: Expected %f, got %f", res, resl)
	}
	t.Logf("Norm Result: %f", resl)

}

func TestVectorMax(t *testing.T) {
	v := NewVector([]float64{1, 2, 3})
	res := 2
	resl := v.Max()

	if res != resl {
		t.Errorf("Max: Expected %d, got %d", res, resl)
	}
	t.Logf("Max Result: %d", resl)

}

func TestVectorPlusAX(t *testing.T) {
	v := NewVector([]float64{1, 2, 3})
	x := NewVector([]float64{2, 3, 4})
	a := 2.4

	res := NewVector([]float64{5.8, 9.2, 12.6})
	v.PlusAX(a, x)
	if notEqual64(res.data, v.data) {
		t.Errorf("PlusAX: Expected %v, got %v", res, v)
	}
	t.Logf("PlusAX Result: %v", v)
}
