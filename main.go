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
	rect(img, 0, 0, width, height, color.White)

	circle(img, (width / 2), (height / 2), 50, color.RGBA{255, 0, 0, 0})

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func rect(img *image.RGBA, x, y, width, height int, color color.Color) {
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
