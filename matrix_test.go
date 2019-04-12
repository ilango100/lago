package lago

import "testing"

func TestMatrixString(t *testing.T) {
	mat, _ := NewMatrix([]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}, 3, 3)
	t.Log(mat.String())
}

func TestMatrixColRow(t *testing.T) {
	mat, _ := NewMatrix([]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}, 3, 3)
	v := mat.Row(1)
	t.Log(v.String())
	v = mat.Column(1)
	t.Log(v.String())

}
