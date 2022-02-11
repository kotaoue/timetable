package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/kotaoue/timetable/pkg/processing"
	"github.com/kotaoue/timetable/pkg/processing/palettes"
)

var (
	width  = flag.Int("w", 400, "image width")
	height = flag.Int("h", 400, "image height")
)

func init() {
	flag.Parse()
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{*width, *height}})
	prc := processing.NewProcessing(processing.Config{Image: img})

	palette := palettes.ShadeGreen
	drawBG(prc, palette[4])
	drawTimeTable(prc, palette[:])

	f, err := os.Create("image.png")
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}

func drawBG(prc *processing.Processing, c color.RGBA) {
	bColor := palettes.MixWhite(c, 4)
	prc.Fill(&bColor)
	prc.Rect(0, 0, *width, *height)
}

func drawTimeTable(prc *processing.Processing, palette []color.RGBA) {
	hours := 24
	ang := 360 / hours
	size := int(math.Max(float64(*width), float64(*height)) / 2)
	for i := 0; i < hours; i++ {
		prc.Fill(&palette[i%4])
		prc.Pie((*width / 2), (*height / 2), size, float64(i*ang), float64((i+1)*ang))
	}
}
