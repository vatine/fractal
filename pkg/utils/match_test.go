package utils

import (
	"testing"
)

func TestMatch(t *testing.T) {
	// Test what root is returned for a given Z and the
	// roots 1, -1, 1+i, 1-1, -1+i, +1-i

	roots := []complex128{1, complex(1, 1), complex(-1, 1), -1, complex(-1, -1), complex(-1, 1)}

	cases := []struct{
		z complex128
		d float64
		r int
	}{
		{1, 0.001, 0}, {-1, 0.001, 3}, {complex(-1, -1), 0.0001, 4},
	}

	for ix, c := range cases {
		want := c.r
		got := Match(c.z, roots, c.d)
		if got != want {
			t.Errorf("Case #%d, got %v, want %v", ix, got, want)
		}
	}
}
