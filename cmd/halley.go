package cmd

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/vatine/fractals/pkg/halley"
	"github.com/vatine/fractals/pkg/utils"
)

func halleyfunc(cmd *cobra.Command, args []string) {
	var roots []complex128
	
	if len(args) == 0 {
		roots = []complex128{1, -1, complex(0, 1), complex(0, -1)}
	} else {
		for _, r := range args {
			root, err := strconv.ParseComplex(r, 128)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
					"arg": r,
				}).Error("parsing failed")
				continue
			} 
			roots = append(roots, root)
		}
	}

	img := halley.Halley(roots, side, maxiter, minZ, maxZ)
	utils.DumpImageToName(img, outname)
}

var halleyCmd = &cobra.Command{
	Args: cobra.MinimumNArgs(0),
	Use: "halley",
	Short: "Halley fractal",
	Long: "Halley fractal, specify the roots of the polynomial you want to render as a fractal.",
	Run: halleyfunc,
}


func init() {
	rootCmd.AddCommand(halleyCmd)
}
