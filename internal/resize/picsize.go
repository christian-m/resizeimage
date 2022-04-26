package resize

import "image"

type PicSize struct {
	Width  int
	Height int
	Border int
}

func (ps *PicSize) AddBorder(border int) {
	if ps.Width > 0 && ps.Height > 0 {
		ps.Border = min(border, min(ps.Width, ps.Height)/3)
	} else {
		ps.Border = border
	}
}

func (ps *PicSize) EnsureImageBounds(img image.Image) {
	if ps.Width < 1 || ps.Width > img.Bounds().Dx() {
		ps.Width = img.Bounds().Dx()
	}
	if ps.Height < 1 || ps.Height > img.Bounds().Dy() {
		ps.Height = img.Bounds().Dy()
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
