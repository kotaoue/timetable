package palettes

import (
	"image/color"
	"math"
)

var ShadeGreen [5]color.RGBA = [5]color.RGBA{
	{46, 184, 37, 255},
	{26, 102, 20, 255},
	{58, 230, 46, 255},
	{61, 242, 48, 255},
	{52, 204, 41, 255},
}

func mix(v uint8, ratio float64) float64 {
	return math.Round(float64(v) * ratio)
}

func add(v uint8, ratio float64) uint8 {
	r := mix(v, ratio)
	if r >= 255 {
		return 255
	}
	return uint8(r)
}

func MixWhite(c color.RGBA, ratio float64) color.RGBA {
	return color.RGBA{
		R: add(c.R, ratio),
		G: add(c.G, ratio),
		B: add(c.B, ratio),
		A: c.A,
	}
}
