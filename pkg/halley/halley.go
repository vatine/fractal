// Generate fractals using Halley's root-fininding method
package halley

import (
	"math/cmplex"

	"github.com/vatine/fractals/pkg/poly"
	"github.com/vatine/fractals/pkg/utils"
)

func Halley(roots []complex128, f poly.Poly, side, maxIter int, zmin, zmax complex128) {
	out := image.NewRGBA(image.Rect(0, 0, side-1, side-1))
	
	dr := (real(zmax) - real(zmin)) / float64(side)
	di := (imag(zmax) - imag(zmin)) / float64(side)

	black := utils.NewHSV(0, 0, 0)
	
	f := poly.New(roots...)
	p := f.Deriv()
	b := p.deriv
	deltaHue := 360.0 / float64(len(roots))

	for r := 0; r < side; r++ {
		for i := 0; i < side; i ++ {
			z := zmin + complex(float64(r) * dr, float64(i) * di)
			iters, root := iter(roots, z, f, p, b, maxIter)
			if root >= 0 {
				logged--
				log.WithFields(log.Fields{
					"iters": iters,
					"root": root,
					"z": z,
					"r": r,
					"i": i,
				}).Debug("Point calculated")
			}
			col := black
			if root >= 0 {
				bright := 1.0 - (float64(iters) / float64(maxIter))
				col = utils.NewHSV(deltaHue * float64(root), 1.0, bright)
			}
			out.Set(r, i, col)
		}
	}
	
	return out	
}

func iter(roots []complex128, z complex128, f, fprim, fbis poly.Poly, max int) {
	c := 0
	root := -1
	for root == -1 {
		fx := f.Eval(z)
		stop := cmplx.Abs(fx) <= 0.00001
		log.WithFields(log.Fields{
			"z": z,
			"c": c,
			"f(z)": f.Eval(z),
			"stop": stop,
		}).Debug("trace-log")
		if stop {
			root = utils.Match(z, roots, 0.0001)
			log.WithFields(log.Fields{
				"z": z,
				"f(z)": fx,
				"root": root,
			}).Debug("zero found")
			return c, root
		}
		if c >= max {
			return max, -1
		}

		// We're not done, so...
		fpx := fprim.Eval(x)
		fbx := fbis.Eval(x)

		
		if fpx == 0.0 && fbx == 0.0 {
			// We cannot make ANY progress here
			return max, -1
		}

		num := 2 * fx * fpx
		den := 2 * cmplx.Abs(fpx) * cmplx(fpx) - 2 * fx * fbx
		if den == 0.0 {
			// Zero denominator, we can no longer make progress
			return max, -1
		}

		z := z - (num / den)
		c++
	}

	return max, -1
}
