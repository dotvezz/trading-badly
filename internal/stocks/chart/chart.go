package chart

import (
	"errors"
	"image"
	"image/color"
	"math"

	gochart "github.com/wcharczuk/go-chart"
	"github.com/dotvezz/trading-badly/internal/stocks"
)

func FromTicks(ticks []stocks.Tick, dimensions ...int) (image.Image, error) {
	if len(dimensions) < 2 {
		dimensions = []int{400, 200}
	}
	if len(ticks) == 0 {
		return nil, errors.New("nothing to chart")
	}

	min, max := ticks[0].Price, ticks[0].Price
	for i := range ticks {
		if ticks[i].Price == 0 {
			continue
		}
		if ticks[i].Price < min {
			min = ticks[i].Price
		}
		if ticks[i].Price > max {
			max = ticks[i].Price
		}
	}

	return &chart{
		min:    min,
		max:    max,
		ticks:  ticks,
		width:  dimensions[0],
		height: dimensions[1],
	}, nil
}

type chart struct {
	ticks    []stocks.Tick
	width    int
	height   int
	min, max float64
}

func (c chart) ColorModel() color.Model {
	return color.RGBAModel
}

func (c chart) Bounds() image.Rectangle {
	return image.Rectangle{
		Max: image.Point{
			X: c.width,
			Y: c.height,
		},
	}
}

func (c chart) At(x, y int) color.Color {
	tickSelector := int((float64(x)/float64(c.width))*float64(len(c.ticks)))
	tick := c.ticks[tickSelector]

	track := (tick.Price-c.min) * float64(c.height) / (c.max - c.min)
	if y == int(math.Round(track)) {
		return color.RGBA{G: 255, A: 255}
	} else {
		return color.White
	}
}
