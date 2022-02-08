package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/kotaoue/timetable/pkg/processing"
)

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	width := 400
	height := 400
	size := int(math.Max(float64(width), float64(height)) / 2)

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	prc := processing.NewProcessing(processing.Config{Image: img})

	prc.Fill(&color.RGBA{255, 255, 255, 255})
	prc.Rect(0, 0, width, height)

	prc.Stroke(&color.RGBA{0, 0, 0, 255})
	prc.Circle((width / 2), (height / 2), size)

	hours := 24
	ang := 360 / hours
	for i := 0; i < hours; i++ {
		prc.Fill(&color.RGBA{uint8(255 / (i + 1)), 255, 255, 255})
		fmt.Printf("i:%d start:%g end:%g\n", i, float64(i*ang), float64((i+1)*ang))
		prc.Pie((width / 2), (height / 2), size, float64(i*ang), float64((i+1)*ang))
	}

	f, err := os.Create("image.png")
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}
