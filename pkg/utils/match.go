package utils

import (
	"math/cmplx"
)

// Match a found root to closest of eth roots. Return -1 if nothing is
// within range.
func Match(z complex128, roots []complex128, delta float64) int {
	for ix, r := range roots {
		if cmplx.Abs(r - z) < delta {
			return ix
		}
	}

	return -1
}
