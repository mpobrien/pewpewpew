package pewpewpew

import "testing"

func TestVectors(t *testing.T) {
	expectedSum := Vector{2, 6, 10}
	v := &Vector{3, 4, 5}
	r := v.Add(Vector{-1, 2, 5})
	if !r.Equals(expectedSum) {
		t.Errorf("Expected %s got %s", expectedSum, v)
	}

	a, b := Vector{2, 3, 4}, Vector{5, 6, 7}
	c := Cross(a, b)
	if !c.Equals(Vector{-3, 6, -3}) {
		t.Errorf("got %s", c)
	}

	d := Dot(a, b)
	if d != 56.0 {
		t.Errorf("expected 56, got %v", d)
	}

}
