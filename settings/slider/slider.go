package slider

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type MouseActionListener interface {
	MouseClicked()
}

type Slider struct {
	position        pixel.Vec
	backgroundColor pixel.RGBA
	circleColor     pixel.RGBA
	radius          float64
	length          float64
	currentValue    float64
	baseVal         float64
	minVal          float64
	maxVal          float64
}

func NewSlider(position pixel.Vec, radius, length, baseV, minV, maxV float64, backgroundColor, circleColor pixel.RGBA) *Slider {
	return &Slider{
		position:        position,
		backgroundColor: backgroundColor,
		circleColor:     circleColor,
		radius:          radius,
		length:          length,
		currentValue:    baseV,
		baseVal:         baseV,
		minVal:          minV,
		maxVal:          maxV,
	}
}

func (s *Slider) Draw(imd *imdraw.IMDraw) {
	imd.Color = s.backgroundColor
	imd.Push(s.position)
	imd.Push(pixel.V(s.position.X+s.length, s.position.Y))
	imd.Line(2)

	imd.Color = s.circleColor
	x := s.currentValue / (s.maxVal - s.minVal) * s.length
	imd.Push(s.position.Add(pixel.V(x, 0)))
	imd.Circle(s.radius, 0)
}

func (s *Slider) Value() float64 {
	return s.currentValue
}

func (s *Slider) UpdateValue(value float64) {
	if value >= s.minVal && value <= s.maxVal {
		s.currentValue = value
	}
}

func (s *Slider) Length() float64 {
	return s.length
}

func (s *Slider) MinValue() float64 {
	return s.minVal
}

func (s *Slider) MaxValue() float64 {
	return s.maxVal
}

func (s *Slider) Position() pixel.Vec {
	return s.position
}

func (s *Slider) AddActionListener(listener MouseActionListener) {
	listener.MouseClicked()
}
