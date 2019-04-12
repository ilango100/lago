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
