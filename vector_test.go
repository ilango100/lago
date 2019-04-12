package lago

import "testing"

func TestVectorString(t *testing.T) {
	v := NewVector([]float64{1, 2, 3, 4, 5})
	res := "[ 1, 2, 3, 4, 5, \b\b ]"
	str := v.String()
	if str != res {
		t.Errorf("Error converting to string: '%s' %s '%s'", str, "!=", res)
	}
	t.Logf("Vector string: %s", str)
}
