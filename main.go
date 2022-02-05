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

	circle(img, image.Point{100, 50}, 25, color.RGBA{255, 0, 0, 0})

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

func circle(img *image.RGBA, center image.Point, radius int, color color.Color) {
	for r := 0.0; r < 2.0*float64(radius); r += 0.1 {
		x := int(float64(center.X) + float64(radius)*math.Cos(r))
		y := int(float64(center.Y) + float64(radius)*math.Sin(r))
		img.Set(x, y, color)
	}
}
