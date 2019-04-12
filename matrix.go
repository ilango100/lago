package lago

import (
	"fmt"

	"github.com/ilango100/blasgo"
)

//Matrix denotes 2d matrix
type Matrix struct {
	data     []float64
	m, n, ld int
}

//NewMatrix creates a new matrix
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

func (mat Matrix) String() string {
	s := ""
	ind := 0
	for i := 0; i < mat.m; i++ {
		s += "| "
		for j := 0; j < mat.n; j++ {
			s += fmt.Sprintf("%f ", mat.data[ind])
			ind++
		}
		ind += mat.ld - mat.n
		s += "|\n"

	}
	return s
}

//Row returns i-th row
func (mat Matrix) Row(i int) (v Vector) {
	v.data = make([]float64, mat.n)
	blasgo.DCOPY(mat.n, mat.data, 1, v.data, 1)
	v.inc = 1
	v.n = mat.n
	return v
}

//Column returns i-th column
func (mat Matrix) Column(i int) (v Vector) {
	v.data = make([]float64, mat.m)
	blasgo.DCOPY(mat.m, mat.data, mat.ld, v.data, 1)
	v.inc = 1
	v.n = mat.m
	return v
}
