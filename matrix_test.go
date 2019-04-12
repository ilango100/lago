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
	v := mat.Row(1)
	row := []float64{4, 5, 6}
	if notEqual64(row, v.data) {
		t.Errorf("Row(1): Expected %v, got %v", row, v)
	}
	t.Logf("Row(1): %v", v)
	v = mat.Col(1)
	col := []float64{2, 5, 8}
	if notEqual64(col, v.data) {
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
