package lago

import "math"

const (
	errprecision = 0.0001
)

func notEqual64(a, b []float64) bool {
	if len(a) != len(b) {
		return true
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > errprecision {
			return true
		}
	}
	return false
}
