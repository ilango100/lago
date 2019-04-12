package lago

import (
	"fmt"
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
		s += fmt.Sprint(v.data[i*v.inc], ", ")
	}
	s += "\b\b ]"
	return s
}
