package utils

import (
	"math"
)

type HSL struct {
	h, s, l float64
}

func NewHSL(h, s, l float64) HSL {
	return HSL{h, s, l}
}

func (col HSL) RGBA() (r, g, b, a uint32) {
	c := (1 - math.Abs(2 * col.s - 1.0)) * col.l
	hprim := col.h / 60.0
	x := c * (1 - math.Abs(math.Mod(hprim, 2.0) - 1))

	xint := uint32(math.Round(0xffff * x))
	cint := uint32(math.Round(0xffff * c))
	m := uint32(math.Round(0xffff * (col.l - c / 2.0)))
	
	switch {
	case (0.0 <= hprim ) && (hprim <= 1.0):
 		return cint + m, xint + m, m, 0xffff
	case (1.0 <= hprim ) && (hprim <= 2.0):
		return xint + m, cint + m, m, 0xffff
	case (2.0 <= hprim ) && (hprim <= 3.0):
		return m, cint + m, xint + m, 0xffff
	case (3.0 <= hprim ) && (hprim <= 4.0):
		return m, xint + m, cint + m, 0xffff
	case (4.0 <= hprim ) && (hprim <= 5.0):
		return xint + m, m, cint + m, 0xffff
	case (5.0 <= hprim ) && (hprim <= 6.0):
		return cint + m, m, xint + m, 0xffff
	default:
		return m, m, m, 0xffff		
	}
	return 0, 0, 0, 0xffff
}


type HSV struct {
	h float64
	s float64
	v float64
}

func NewHSV(h, s, v float64) HSV {
	return HSV{h, s, v}
}

func (hsv HSV) RGBA() (r, g, b, a uint32) {
	c := hsv.s * hsv.v
	hprim := hsv.h / 60.0
	x := c * (1 - math.Abs(math.Mod(hprim, 2.0) - 1))

	xint := uint32(math.Round(0xffff * x))
	cint := uint32(math.Round(0xffff * c))
	m := uint32(math.Round(0xffff * (hsv.v - c)))
	switch {
	case (0.0 <= hprim ) && (hprim <= 1.0):
 		return cint + m, xint + m, m, 0xffff
	case (1.0 <= hprim ) && (hprim <= 2.0):
		return xint + m, cint + m, m, 0xffff
	case (2.0 <= hprim ) && (hprim <= 3.0):
		return m, cint + m, xint + m, 0xffff
	case (3.0 <= hprim ) && (hprim <= 4.0):
		return m, xint + m, cint + m, 0xffff
	case (4.0 <= hprim ) && (hprim <= 5.0):
		return xint + m, m, cint + m, 0xffff
	case (5.0 <= hprim ) && (hprim <= 6.0):
		return cint + m, m, xint + m, 0xffff
	default:
		return m, m, m, 0xffff		
	}
	return 0, 0, 0, 0xffff
}
