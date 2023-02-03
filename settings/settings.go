package settings

import (
	"GameOfLife/settings/slider"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math"
)

type Settings struct {
	sizeSlider *slider.Slider
	center     pixel.Vec
}

func NewSettings(center pixel.Vec) *Settings {
	sizeSlider := slider.NewSlider(center.Sub(pixel.V(140, 0)), 8, 100, 5, 2, 100, pixel.RGB(1, 1, 1), pixel.RGB(0, 0, 1))
	return &Settings{
		sizeSlider: sizeSlider,
		center:     center,
	}
}

func (s *Settings) OpenSettings(imd *imdraw.IMDraw) {
	s.roundRect(imd, 300, 400, 60, 60)
	s.sizeSlider.Draw(imd)
}

func (s *Settings) roundRect(imd *imdraw.IMDraw, width, height, rx, ry float64) {
	imd.Color = pixel.RGB(0.2, 0.2, 0.2)

	vec := pixel.V(width/2, height/2)
	imd.Push(s.center.Sub(vec))
	imd.Push(s.center.Add(vec))
	imd.Rectangle(0)

	topLeft := pixel.V(s.center.X-width/2+rx, s.center.Y+height/2-ry/8)
	topRight := pixel.V(s.center.X+width/2-rx, s.center.Y+height/2-ry/8)
	bottomLeft := pixel.V(s.center.X-width/2+rx, s.center.Y-height/2+ry/8)
	bottomRight := pixel.V(s.center.X+width/2-rx, s.center.Y-height/2+ry/8)

	// Define radius for each corner
	elVec := pixel.V(rx, ry)

	// Draw 4 arcs for each corner, starting from Pi, Pi/2, 0, -Pi/2 respectively
	imd.Push(topLeft)
	imd.EllipseArc(elVec, math.Pi, math.Pi/2, 0)

	imd.Push(topRight)
	imd.EllipseArc(elVec, math.Pi/2, 0, 0)

	imd.Push(bottomRight)
	imd.EllipseArc(elVec, 0, -math.Pi/2, 0)

	imd.Push(bottomLeft)
	imd.EllipseArc(elVec, -math.Pi/2, -math.Pi, 0)

	// Draw second rectangle
	vec2 := vec.Sub(pixel.V(rx, -ry*(1-1/7.5)))
	imd.Push(s.center.Add(vec2))
	imd.Push(s.center.Sub(vec2))
	imd.Rectangle(0)
}
