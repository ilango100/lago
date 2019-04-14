package lago

import (
	"fmt"

	"github.com/ilango100/blasgo"
)

//Matrix denotes 2d matrix.
type Matrix struct {
	data     []float64
	m, n, ld int
	trans    blasgo.Transpose
}

//NewMatrix creates a new matrix.
func NewMatrix(f []float64, m, n int) (mat Matrix, err error) {
	if m*n <= 0 {
		return mat, fmt.Errorf("m*n <= 0")
	}
	if f == nil {
		f = make([]float64, m*n)
	}
	if m*n != len(f) {
		return mat, fmt.Errorf("m*n != len(f)")
	}
	mat.data = f
	mat.m = m
	mat.n = n
	mat.ld = n
	mat.trans = blasgo.NoTrans
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
	for i := 0; i < mat.m; i++ {
		s += fmt.Sprintln(mat.Row(i))
	}
	return s
}

//Row returns i-th row.
func (mat Matrix) Row(i int) (v Vector) {
	// v.data = make([]float64, mat.n)
	// blasgo.DCOPY(mat.n, mat.data[i*mat.ld:], 1, v.data, 1)
	// v.inc = 1

	if mat.trans == blasgo.Trans {
		mat.trans = blasgo.NoTrans
		return mat.Col(i)
	}

	if i < 0 {
		i = 0
	} else if i >= mat.m {
		i = mat.m - 1
	}

	v.data = mat.data[i*mat.ld:]
	v.inc = 1
	v.n = mat.n

	return v
}

//Col returns i-th column.
func (mat Matrix) Col(i int) (v Vector) {
	// v.data = make([]float64, mat.m)
	// blasgo.DCOPY(mat.m, mat.data[i:], mat.ld, v.data, 1)
	// v.inc = 1

	if mat.trans == blasgo.Trans {
		mat.trans = blasgo.NoTrans
		return mat.Row(i)
	}

	if i < 0 {
		i = 0
	} else if i >= mat.n {
		i = mat.n - 1
	}

	v.data = mat.data[i:]
	v.inc = mat.ld
	v.n = mat.m

	return v
}

//T transposes the matrix
func (mat Matrix) T() Matrix {
	if mat.trans == blasgo.Trans {
		mat.trans = blasgo.NoTrans
		return mat
	}
	mat.trans = blasgo.Trans
	return mat
}

//PlusAMM => mat = b*mat + a*M1*M2
func (mat Matrix) PlusAMM(b float64, m1 Matrix, m2 Matrix, a float64) Matrix {
	if mat.data == nil {
		mat, _ = NewMatrix(nil, m1.m, m2.n)
	}
	if m1.n != m2.m {
		return mat
	}
	blasgo.DGEMM(blasgo.RowMajor, m1.trans, m2.trans, m1.m, m2.n, m1.n, a, m1.data, m1.ld, m2.data, m2.ld, b, mat.data, mat.ld)
	return mat
}

//PlusAVV => mat = mat + a*v*vT
func (mat Matrix) PlusAVV(a float64, v1 Vector, v2 Vector) Matrix {
	if mat.data == nil {
		mat, _ = NewMatrix(nil, v1.n, v2.n)
	}
	if mat.m != v1.n || mat.n != v2.n {
		return mat
	}
	blasgo.DGER(blasgo.RowMajor, mat.m, mat.n, a, v1.data, v1.inc, v2.data, v2.inc, mat.data, mat.ld)
	return mat
}
