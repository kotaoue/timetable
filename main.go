package main

import (
	"image"
	"image/color"
	"image/png"
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

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
