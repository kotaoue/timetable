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

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x%2 == 0:
				img.Set(x, y, color.Black)
			case x%2 != 0:
				img.Set(x, y, color.White)
			}
		}
	}

	circle(img, image.Point{100, 50}, 25, color.RGBA{255, 0, 0, 0})

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func circle(img *image.RGBA, center image.Point, radius int, color color.Color) {
	for rad := 0.0; rad < 2.0*float64(radius); rad += 0.1 {
		x := int(float64(center.X) + float64(radius)*math.Cos(rad))
		y := int(float64(center.Y) + float64(radius)*math.Sin(rad))
		img.Set(x, y, color)
	}
}
