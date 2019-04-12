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
