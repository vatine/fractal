package poly

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		p1, p2, pe Poly
	}{
		{Poly{[]complex128{complex(1, 0)}}, Poly{[]complex128{complex(0, 1), complex(2, 1)}}, Poly{[]complex128{complex(1, 1), complex(2, 1)}}},
	}

	for ix, c := range cases {
		seen := c.p1.Add(c.p2)
		if !seen.Equal(c.pe) {
			t.Errorf("Case #%d, got %s, want %s", ix, seen, c.pe)
		}
	}
}

func TestMul(t *testing.T) {
	cases := []struct {
		p1, p2, pe Poly
	}{
		{Poly{[]complex128{1}}, Poly{[]complex128{complex(0, 1)}}, Poly{[]complex128{complex(0, 1)}}},
		{Poly{[]complex128{0, 1}}, Poly{[]complex128{0, 2}}, Poly{[]complex128{0, 0, 2}}},
	}

	for ix, c := range cases {
		seen := c.p1.Mul(c.p2)
		if !seen.Equal(c.pe) {
			t.Errorf("Case #%d, got %s, want %s", ix, seen, c.pe)
		}
	}
}	

func TestNewPoly(t *testing.T) {
	cases := []struct{
		roots []complex128
		exp   Poly
	}{
		{[]complex128{1, -1}, Poly{[]complex128{-1, 0, 1}}},
	}
	for ix, c := range cases {
		seen := NewPoly(c.roots...)
		if !seen.Equal(c.exp) {
			t.Errorf("Case #%d, got %s, want %s", ix, seen, c.exp)
		}
	}
}
