package slider

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

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
	imd.Push(s.position.Add(pixel.V(s.currentValue, 0)))
	imd.Circle(s.radius, 0)
}

func (s *Slider) GetValue() float64 {
	return s.currentValue
}

func (s *Slider) UpdateValue(value float64) {
	if value >= s.minVal && value <= s.maxVal {
		s.currentValue = value
	}
}
