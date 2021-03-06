package processing

import (
	"image"
	"image/color"
	"math"
)

type Processing struct {
	img         *image.RGBA
	fillColor   *color.RGBA
	strokeColor *color.RGBA
}

type Config struct {
	Image       *image.RGBA
	FillColor   *color.RGBA
	StrokeColor *color.RGBA
}

func NewProcessing(cfg Config) *Processing {
	return &Processing{
		img:         cfg.Image,
		fillColor:   cfg.FillColor,
		strokeColor: cfg.StrokeColor,
	}
}

func (p *Processing) Fill(c *color.RGBA) {
	p.fillColor = c
}

func (p *Processing) NoFill() {
	p.fillColor = nil
}

func (p *Processing) Stroke(c *color.RGBA) {
	p.strokeColor = c
}

func (p *Processing) NoStroke() {
	p.strokeColor = nil
}

func (p *Processing) Line(startX, startY, endX, endY int) {
	sx := float64(startX)
	sy := float64(startY)
	ex := float64(endX)
	ey := float64(endY)

	dx := math.Abs(ex - sx)
	dy := math.Abs(ey - sy)

	swx := 1.0
	if sx >= ex {
		swx = -1
	}
	swy := 1.0
	if sy >= ey {
		swy = -1
	}

	e := dx - dy

	x := sx
	y := sy
	for {
		p.img.Set(int(x), int(y), p.strokeColor)
		if x == ex && y == ey {
			return
		}

		e2 := 2 * e
		if e2 > -dy {
			e = e - dy
			x = x + swx
		}
		if e2 < dx {
			e = e + dx
			y = y + swy
		}
	}
}

func (p *Processing) Rect(x, y, width, height int) {
	if p.fillColor != nil {
		p.fillRect(x, y, width, height)
	}
	if p.strokeColor != nil {
		p.noFillRect(x, y, width, height)
	}
}

func (p *Processing) fillRect(x, y, width, height int) {
	for i := x; i < (width + x); i++ {
		for j := y; j < (height + y); j++ {
			p.img.Set(i, j, p.fillColor)
		}
	}
}

func (p *Processing) noFillRect(x, y, width, height int) {
	for i := x; i < (x + width); i++ {
		p.img.Set(i, y, p.strokeColor)
		p.img.Set(i, (y + height), p.strokeColor)
	}
	for i := y; i < (y + height); i++ {
		p.img.Set(x, i, p.strokeColor)
		p.img.Set((x + width), i, p.strokeColor)
	}
}

func (p *Processing) Pie(x, y, extent int, startAngle, endAngle float64) {
	if p.fillColor != nil {
		p.fillPie(x, y, extent, startAngle, endAngle)
	}
	if p.strokeColor != nil {
		p.noFillPie(x, y, extent, startAngle, endAngle)
	}
}

func (p *Processing) Circle(x, y, extent int) {
	if p.fillColor != nil {
		p.fillPie(x, y, extent, 0, 360)
	}
	if p.strokeColor != nil {
		p.noFillPie(x, y, extent, 0, 360)
	}
}

func (p *Processing) fillPie(x, y, extent int, startAngle, endAngle float64) {
	for ang := startAngle; ang < endAngle; ang += 0.1 {
		for r := 0.0; r < float64(extent/2); r++ {
			i := int(float64(x) + r*math.Cos(ang*math.Pi/180))
			j := int(float64(y) + r*math.Sin(ang*math.Pi/180))
			p.img.Set(i, j, p.fillColor)
		}
	}
}

func (p *Processing) noFillPie(x, y, extent int, startAngle, endAngle float64) {
	r := float64(extent / 2)
	for ang := startAngle; ang < endAngle; ang += 0.1 {
		i := int(float64(x) + r*math.Cos(ang*math.Pi/180))
		j := int(float64(y) + r*math.Sin(ang*math.Pi/180))
		p.img.Set(i, j, p.strokeColor)
	}

	if startAngle != 0 && endAngle != 360 {
		{
			i := int(float64(x) + r*math.Cos(startAngle*math.Pi/180))
			j := int(float64(y) + r*math.Sin(startAngle*math.Pi/180))
			p.Line(x, y, i, j)
		}
		{
			i := int(float64(x) + r*math.Cos(endAngle*math.Pi/180))
			j := int(float64(y) + r*math.Sin(endAngle*math.Pi/180))
			p.Line(x, y, i, j)
		}
	}
}
