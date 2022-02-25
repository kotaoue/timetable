package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/kotaoue/timetable/pkg/processing"
	"github.com/kotaoue/timetable/pkg/processing/palettes"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	width    = flag.Int("w", 400, "image width")
	height   = flag.Int("h", 400, "image height")
	fontFile *truetype.Font
)

func init() {
	flag.Parse()

	if err := initFont(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initFont() error {
	ftBinary, err := ioutil.ReadFile("Koruri-Regular.ttf")
	if err != nil {
		return err
	}

	fontFile, err = truetype.Parse(ftBinary)
	if err != nil {
		return err
	}

	return nil
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

	opt := &truetype.Options{Size: 10}
	for k, s := range []string{"タイムテーブル", "Line1", "Line2", "Line3", "Line4"} {
		drawString(img, 15, 15+(15*k), s, opt)
	}

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
	size := int(math.Max(float64(*width), float64(*height)) * 0.75)
	for i := 0; i < hours; i++ {
		prc.Fill(&palette[i%4])
		prc.Pie((*width / 2), (*height / 2), size, float64(i*ang), float64((i+1)*ang))
	}
}

func drawString(img *image.RGBA, x, y int, s string, opt *truetype.Options) error {
	face := truetype.NewFace(fontFile, opt)

	col := color.RGBA{0, 0, 0, 255}
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}
	d.DrawString(s)

	return nil
}
