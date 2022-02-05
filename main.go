package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/kotaoue/timetable/pkg/processing"
)

func main() {
	width := 200
	height := 100

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	prc := processing.NewProcessing(processing.Config{Image: img})
	prc.Fill(&color.RGBA{255, 255, 255, 255})
	prc.Stroke(&color.RGBA{0, 0, 0, 255})
	prc.Rect(0, 0, width, height)
	prc.NoFill()
	prc.Rect(50, 40, 100, 20)

	prc.Circle(50, (height / 2), 50)
	prc.Fill(&color.RGBA{0, 0, 0, 255})
	prc.Circle(150, (height / 2), 50)

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
