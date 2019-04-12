package lago

import (
	"fmt"

	"github.com/ilango100/blasgo"
)

//Matrix denotes 2d matrix.
type Matrix struct {
	data     []float64
	m, n, ld int
}

//NewMatrix creates a new matrix.
func NewMatrix(f []float64, m, n int) (mat Matrix, err error) {
	if m*n != len(f) {
		return mat, fmt.Errorf("m*n != len(f)")
	}
	if len(f) <= 0 {
		return mat, fmt.Errorf("len(f) <= 0")
	}
	mat.data = f
	mat.m = m
	mat.n = n
	mat.ld = n
	return mat, nil
}

//I creates identity matrix of order m x m.
func I(m int) Matrix {
	mat, _ := NewMatrix(make([]float64, m*m), m, m)
	for i := 0; i < m; i++ {
		mat.data[i*m+i] = 1
	}
	return mat
}

func (mat Matrix) String() string {
	s := ""
	ind := 0
	for i := 0; i < mat.m; i++ {
		s += "| "
		for j := 0; j < mat.n; j++ {
			s += fmt.Sprint(mat.data[ind], " ")
			ind++
		}
		ind += mat.ld - mat.n
		s += "|\n"

	}
	return s
}

//Row returns i-th row.
func (mat Matrix) Row(i int) (v Vector) {
	v.data = make([]float64, mat.n)
	blasgo.DCOPY(mat.n, mat.data[i*mat.ld:], 1, v.data, 1)
	v.inc = 1
	v.n = mat.n
	return v
}

//Col returns i-th column.
func (mat Matrix) Col(i int) (v Vector) {
	v.data = make([]float64, mat.m)
	blasgo.DCOPY(mat.m, mat.data[i:], mat.ld, v.data, 1)
	v.inc = 1
	v.n = mat.m
	return v
}
