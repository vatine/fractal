package poly

// A package for implementing various operations on polynomials.

import (
	"fmt"
	"strings"
)

type Poly struct {
	coeffs []complex128
}

// Generate a new polynomial with the given roots.
func New(roots ...complex128) Poly {
	rv := Poly{[]complex128{1}}
	for _, r := range roots {
		t := Poly{[]complex128{-r, 1}}
		rv = rv.Mul(t)
	}

	return rv
}

// Format a polynomial as a string
func (p Poly) String() string {
	var parts []string
	for d := len(p.coeffs)-1; d >= 0; d-- {
		if p.coeffs[d] != 0 {
			parts = append(parts, fmt.Sprintf("(%g)z^%d", p.coeffs[d], d))
		}
	}

	return strings.Join(parts, " + ")
}

// Add two polynomials
func (p1 Poly) Add(p2 Poly) Poly {
	l1 := len(p1.coeffs)
	l2 := len(p2.coeffs)

	var c []complex128
	max := l1
	if max < l2 {
		max = l2
	}

	for i := 0; i < max; i++ {
		switch {
		case i < l1 && i < l2:
			c = append(c, p1.coeffs[i] + p2.coeffs[i])
		case i < l1:
			c = append(c, p1.coeffs[i])
		case i < l2:
			c = append(c, p2.coeffs[i])
		}
	}

	return Poly{coeffs: c}
}

// Multiply two polynomials
func (p1 Poly) Mul(p2 Poly) Poly {
	l1 := len(p1.coeffs)
	l2 := len(p2.coeffs)
	c := make([]complex128, l1+l2)

	for i1, c1 := range p1.coeffs {
		for i2, c2 := range p2.coeffs {
			c[i1+i2] += c1 * c2
		}
	}

	return Poly{coeffs: c}
}

// Take the derivative of a polynomial
func (p Poly) Deriv() Poly {
	c := make([]complex128, len(p.coeffs)-1)

	for d := 1; d < len(p.coeffs); d++ {
		c[d-1] = p.coeffs[d] * complex(float64(d), 0)
	}

	return Poly{coeffs: c}
}

// Return true if two polynomials have the same coefficients
func (p1 Poly) Equal(p2 Poly) bool {
	l1 := len(p1.coeffs)
	l2 := len(p2.coeffs)

	min := l1
	if (l2 < l1) {
		min = l2
	}

	for d := 0; d < min; d++ {
		if p1.coeffs[d] != p2.coeffs[d] {
			return false
		}
	}

	for d := min; d < l1; d++ {
		if p1.coeffs[d] != 0 {
			return false
		}
	}
	for d := min; d < l2; d++ {
		if p2.coeffs[d] != 0 {
			return false
		}
	}

	return true
}

// Evaluate a polynomial at a specific point in the complex plane
func (p Poly) Eval(z complex128) complex128 {
	var rv complex128
	for d := len(p.coeffs)-1; d >= 0; d-- {
		rv = rv*z + p.coeffs[d]
	}

	return rv
}
