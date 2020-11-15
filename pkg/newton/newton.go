package newton

// Code for Newton-Raphson fractals

import (
	"image"
	"math/cmplx"

	log "github.com/sirupsen/logrus"
	
	"github.com/vatine/fractals/pkg/poly"
	"github.com/vatine/fractals/pkg/utils"
)

var logged int = 0

// Compute a Newton-Raphson fractal with the given roots, return a
// square image with the specified side and the bottom-left corner at
// zmin, and eth top-right corner at zmax. Iterate for at most maxIter
// iterations before giving up.
func NewtonRaphson(roots []complex128, side, maxIter int, zmin, zmax complex128) image.Image {
	out := image.NewRGBA(image.Rect(0, 0, side-1, side-1))
	
	dr := (real(zmax) - real(zmin)) / float64(side)
	di := (imag(zmax) - imag(zmin)) / float64(side)

	black := utils.NewHSV(0, 0, 0)
	
	f := poly.New(roots...)
	d := f.Deriv()
	deltaHue := 360.0 / float64(len(roots))

	for r := 0; r < side; r++ {
		for i := 0; i < side; i ++ {
			z := zmin + complex(float64(r) * dr, float64(i) * di)
			iters, root := iter(roots, z, f, d, maxIter)
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

func iter(roots []complex128, z complex128, f, fprim poly.Poly, max int) (int, int) {
	c := 0
	root := -1
	for root == -1  {
		stop := cmplx.Abs(f.Eval(z)) <= 0.00001
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
				"f(z)": f.Eval(z),
				"root": root,
			}).Debug("zero found")
			return c, root
		}
		if c >= max {
			return max, -1
		}
		// Avoid divide-by-zero
		d := fprim.Eval(z)
		if d == 0.0 {
			log.WithFields(log.Fields{
				"f": f,
				"f'": fprim,
				"z": z,
			}).Debug("Zero slope, returning early.")
			return max, -1
		}

		z = z - (f.Eval(z) / d)
		c++
	}

			
	return c, root
}
