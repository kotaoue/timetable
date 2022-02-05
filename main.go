package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	width := 200
	height := 100

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	fillRect(img, 0, 0, width, height, color.White)
	rect(img, 50, 40, 100, 20, color.Black)

	circle(img, 50, (height / 2), 50, color.Black)
	fillCircle(img, 150, (height / 2), 50, color.Black)

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func rect(img *image.RGBA, x, y, width, height int, color color.Color) {
	for i := x; i < (x + width); i++ {
		img.Set(i, y, color)
		img.Set(i, (y + height), color)
	}
	for i := y; i < (y + height); i++ {
		img.Set(x, i, color)
		img.Set((x + width), i, color)
	}
}

func fillRect(img *image.RGBA, x, y, width, height int, color color.Color) {
	for i := x; i < (width + x); i++ {
		for j := y; j < (height + y); j++ {
			img.Set(i, j, color)
		}
	}
}

func circle(img *image.RGBA, x, y, extent int, color color.Color) {
	r := float64(extent / 2)
	for ang := 0.0; ang < 360.0; ang += 0.1 {
		i := int(float64(x) + r*math.Cos(ang))
		j := int(float64(y) + r*math.Sin(ang))
		img.Set(i, j, color)
	}
}

func fillCircle(img *image.RGBA, x, y, extent int, color color.Color) {
	for ang := 0.0; ang < 360.0; ang += 0.1 {
		for r := 0.0; r < float64(extent/2); r++ {
			i := int(float64(x) + r*math.Cos(ang))
			j := int(float64(y) + r*math.Sin(ang))
			img.Set(i, j, color)
		}
	}
}
