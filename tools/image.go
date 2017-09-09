package tools

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func ImageScale(r io.Reader, w io.Writer, dx int, dy int) error {
	if img, img_type, err := image.Decode(r); err != nil {
		return err
	} else {
		var (
			dx_size = float64(img.Bounds().Dx()) / float64(dx)
			dy_size = float64(img.Bounds().Dy()) / float64(dy)
			err     error
		)
		img_new := image.NewNRGBA64(image.Rect(0, 0, dx, dy))
		for x := 0; x < dx; x++ {
			for y := 0; y < dy; y++ {
				img_new.Set(x, y, img.At(int(float64(x)*dx_size), int(float64(y)*dy_size)))
			}
		}
		switch img_type {
		case "png":
			err = png.Encode(w, img_new)
		case "jpeg":
			err = jpeg.Encode(w, img_new, &jpeg.Options{5})
		}
		if err != nil {
			return err
		} else {
			return nil
		}
	}

}
