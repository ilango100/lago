package lago

import (
	"testing"
)

func TestMatrixString(t *testing.T) {
	mat, _ := NewMatrix([]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}, 3, 3)
	t.Log(mat.String())
}

func TestMatrixRowCol(t *testing.T) {
	mat, _ := NewMatrix([]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}, 3, 3)
	v := mat.Row(1).Result()
	row := []float64{4, 5, 6}
	if notEqual64(row, v) {
		t.Errorf("Row(1): Expected %v, got %v", row, v)
	}
	t.Logf("Row(1): %v", v)
	v = mat.Col(1).Result()
	col := []float64{2, 5, 8}
	if notEqual64(col, v) {
		t.Errorf("Col(1): Expected %v, got %v", row, v)
	}
	t.Logf("Col(1): %v", v)
}

func TestI(t *testing.T) {
	res := []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	i3 := I(3)
	if notEqual64(res, i3.data) {
		t.Errorf("I(3): Expected %v, got %v", res, i3.data)
	}
	t.Logf("I(3): %v", i3)
}

func TestMatrixPlusAMM(t *testing.T) {
	res := []float64{107.0, 273.8, 440.6, 69.2, 173.9, 278.6, 31.4, 74.0, 116.6}
	a, _ := NewMatrix([]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}, 3, 3)
	b, _ := NewMatrix([]float64{
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
	}, 3, 3)
	c, _ := NewMatrix([]float64{
		1, 4, 7,
		2, 5, 8,
		3, 6, 9,
	}, 3, 3)
	a.PlusAMM(1.2, b, c, 2.3)
	if notEqual64(res, a.data) {
		t.Errorf("PlusAMM: %v != %v", a, res)
	}
	t.Logf("PlusAMM Result:\n%v", a)
}

func TestMatrixPlusAVV(t *testing.T) {
	res := []float64{11.8, 45.2, 78.6, 13.6, 43.4, 73.2, 15.4, 41.6, 67.8}
	a, _ := NewMatrix([]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}, 3, 3)
	b := NewVector([]float64{
		9, 8, 7,
	})
	c := NewVector([]float64{
		1, 4, 7,
	})
	a.PlusAVV(1.2, b, c)
	if notEqual64(res, a.data) {
		t.Errorf("PlusAVV: %v != %v", a, res)
	}
	t.Logf("PlusAVV Result:\n%v", a)
}

func TestMatrixSubMatrix(t *testing.T) {
	res := []float64{5, 6, 8, 9}
	a, _ := NewMatrix([]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}, 3, 3)
	b := a.SubMatrix(1, 2, 1, 2).Copy()
	if notEqual64(res, b.data) {
		t.Errorf("SubMatrix %v != %v", b, res)
	}
	t.Logf("SubMatrix Result:\n%v", b)
}
