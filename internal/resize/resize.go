package resize

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
)

func Resize(ps PicSize, f string, img image.Image, w io.Writer) error {
	if f != "jpeg" && f != "png" {
		return fmt.Errorf("only JPEG and PNG are supported")
	}
	resized := imaging.Fit(img, ps.Width-ps.Border*2, ps.Height-ps.Border*2, imaging.CatmullRom)
	bordered := imaging.New(resized.Bounds().Dx()+ps.Border*2, resized.Bounds().Dy()+ps.Border*2, color.White)
	bordered = imaging.Paste(bordered, resized, image.Pt(ps.Border, ps.Border))
	return encodeImage(bordered, f, w)
}

func encodeImage(i image.Image, f string, w io.Writer) error {
	switch f {
	case "jpeg":
		return jpeg.Encode(w, i, nil)
	case "png":
		return png.Encode(w, i)
	default:
		return fmt.Errorf("only JPEG and PNG are supported")
	}
}
