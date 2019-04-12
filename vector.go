package lago

import (
	"fmt"

	"github.com/ilango100/blasgo"
)

//Vector defines column vector
type Vector struct {
	data   []float64
	inc, n int
}

//NewVector creates a new vector
func NewVector(f []float64) (v Vector) {
	v.data = f
	v.n = len(f)
	v.inc = 1
	return v
}

//String satisfied Stringer interface
func (v Vector) String() (s string) {
	s = "[ "
	for i := 0; i < v.n; i++ {
		s += fmt.Sprint(v.data[i*v.inc], " ")
	}
	s += "]"
	return s
}

//Slice slices vector from i to j
func (v Vector) Slice(arr ...int) (Vector, error) {
	i, j, k := 0, v.n, 1
	switch len(arr) {
	case 0:
	case 1:
		j = arr[0]
	default:
		k = arr[2]
		fallthrough
	case 2:
		i, j = arr[0], arr[1]
	}
	if i >= j {
		return v, fmt.Errorf("i >= j")
	}
	if i > v.n || j > v.n {
		return v, fmt.Errorf("index > v.n")
	}
	v.data = v.data[i*v.inc:]
	//v.inc = v.inc*step + step + v.inc
	v.inc = v.inc * k
	v.n = (j - i) / k
	if (j-i)%k > 0 {
		v.n++
	}
	return v, nil
}

// Assign assigns the value to the vector if dimensions match.
// Note that the Vectors extracted from Matrix also change the value of the Matrix.
func (v Vector) Assign(value []float64) error {
	if len(value) != v.n {
		return fmt.Errorf("Dimensions do not match")
	}

	for i := 0; i < v.n; i++ {
		v.data[i*v.inc] = value[i]
	}

	return nil
}

//Result returns the final result.
func (v Vector) Result() (res []float64) {
	res = make([]float64, v.n)
	for i := range res {
		res[i] = v.data[i*v.inc]
	}
	return res
}

//Dot calculates the dot product. Returns 0 when dimensions don't match.
func (v Vector) Dot(w Vector) (d float64) {
	if v.n != w.n {
		return d
	}
	return blasgo.DDOT(v.n, v.data, v.inc, w.data, w.inc)
}

//Scale scales the vector with scalar.
func (v Vector) Scale(s float64) {
	blasgo.DSCAL(v.n, s, v.data, v.inc)
}

//Copy creates a new copy of vector
func (v Vector) Copy() (w Vector) {
	w.data = make([]float64, v.n)
	w.n = v.n
	w.inc = 1
	blasgo.DCOPY(v.n, v.data, v.inc, w.data, w.inc)
	return w
}

//Sum calculates the sum of elements.
func (v Vector) Sum() float64 {
	return blasgo.DASUM(v.n, v.data, v.inc)
}

//Max returns the max index of elements.
func (v Vector) Max() int {
	return blasgo.IDAMAX(v.n, v.data, v.inc)
}

//Norm calculates the norm of the vector.
func (v Vector) Norm() float64 {
	return blasgo.DNRM2(v.n, v.data, v.inc)
}
