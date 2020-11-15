package utils

import (
	"image"
	"image/png"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// Return an io.Writer based on teh input name, interpret "-" as
// stdout, as per somewhat normal unix manners.
func WriterFromName(name string) (io.Writer, error) {
	if name == "-" {
		return os.Stdout, nil
	}

	f, err := os.Create(name)

	if err != nil {
		log.WithFields(log.Fields{"error": err, "name": name}).Error("error opening file")
		return nil, err
	}

	return f, nil
}

// Write an image to an io.writer
func DumpImage(img image.Image, w io.Writer) error {
	return png.Encode(w, img)
}

func DumpImageToName(img image.Image, name string) error {
	w, err := WriterFromName(name)
	if err != nil {
		return err
	}

	err = DumpImage(img, w)
	if err != nil {
		return err
	}

	if f, ok := w.(*os.File); ok {
		f.Close()
	}

	return nil
}
