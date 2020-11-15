package cmd

import (
	"io"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Various fractal generators.",
}

var (
	side    int
	outname string
	out     io.Writer
	minZ    complex128
	maxZ    complex128
	maxiter int
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVar(&side, "side", 256, "Side (in pixels) of the image.")
	rootCmd.PersistentFlags().StringVar(&outname, "out", "-", "Name of output file.")
	rootCmd.PersistentFlags().Complex128Var(&minZ, "minz", -5.0-5.0i, "Lower-left corner's coordinate.")
	rootCmd.PersistentFlags().Complex128Var(&maxZ, "maxz", 5.0+5.0i, "Upper right corner's coordinate.")
	rootCmd.PersistentFlags().IntVar(&maxiter, "maxiter", 500, "Maximum iteration value.")	
}
