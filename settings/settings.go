package settings

import (
	"GameOfLife/life"
	"GameOfLife/settings/slider"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"math"
)

var (
	atlas      = text.NewAtlas(basicfont.Face7x13, text.ASCII)
	resolution *text.Text
	sizeSlider *slider.Slider
)

type Settings struct {
	win    *pixelgl.Window
	imd    *imdraw.IMDraw
	center pixel.Vec
}

func NewSettings(win *pixelgl.Window, imd *imdraw.IMDraw, center pixel.Vec) *Settings {
	sizeSlider = slider.NewSlider(
		center.Sub(pixel.V(140, -150)),
		8, 130, 5, 3, 300,
		pixel.RGB(0, 0, 0),
		pixel.RGB(1, 0, 0))
	return &Settings{
		win:    win,
		imd:    imd,
		center: center,
	}
}

func (s *Settings) OpenSettings() {

	s.roundRect(400, 400, 60, 60)
	sizeSlider.Draw(s.imd)
	s.imd.Draw(s.win)
	s.createText()
}

func (s *Settings) createText() {
	resolution = text.New(pixel.V(s.center.X, s.center.Y+145), atlas)
	resolution.Color = colornames.Black

	_, err := fmt.Fprint(resolution, "resolution")
	if err != nil {
		return
	}
	resolution.Draw(s.win, pixel.IM.Scaled(resolution.Orig, 2))
}

func (s *Settings) roundRect(width, height, rx, ry float64) {
	s.imd.Color = pixel.RGB(1, 1, 1)

	vec := pixel.V(width/2, height/2)
	s.imd.Push(s.center.Sub(vec))
	s.imd.Push(s.center.Add(vec))
	s.imd.Rectangle(0)

	topLeft := pixel.V(s.center.X-width/2+rx, s.center.Y+height/2)
	topRight := pixel.V(s.center.X+width/2-rx, s.center.Y+height/2)
	bottomLeft := pixel.V(s.center.X-width/2+rx, s.center.Y-height/2)
	bottomRight := pixel.V(s.center.X+width/2-rx, s.center.Y-height/2)

	// Define radius for each corner
	elVec := pixel.V(rx, ry)

	// Draw 4 arcs for each corner, starting from Pi, Pi/2, 0, -Pi/2 respectively
	s.imd.Push(topLeft)
	s.imd.EllipseArc(elVec, math.Pi, math.Pi/2, 0)

	s.imd.Push(topRight)
	s.imd.EllipseArc(elVec, math.Pi/2, 0, 0)

	s.imd.Push(bottomRight)
	s.imd.EllipseArc(elVec, 0, -math.Pi/2, 0)

	s.imd.Push(bottomLeft)
	s.imd.EllipseArc(elVec, -math.Pi/2, -math.Pi, 0)

	// Draw second rectangle
	vec2 := vec.Sub(pixel.V(rx, -ry))
	s.imd.Push(s.center.Add(vec2))
	s.imd.Push(s.center.Sub(vec2))
	s.imd.Rectangle(0)
}

func (s *Settings) Listen(l *life.Life) {

	if s.win.Pressed(pixelgl.MouseButtonLeft) {
		mousePos := s.win.MousePosition()
		if math.Abs(mousePos.Y-sizeSlider.Position().Y) < 20 && ((mousePos.X-sizeSlider.Position().X < sizeSlider.Length()) || (sizeSlider.Position().X-mousePos.X > 0)) {
			var sizeValue = (s.win.MousePosition().X - sizeSlider.Position().X) / sizeSlider.Length() * (sizeSlider.MaxValue() - sizeSlider.MinValue())
			if sizeValue >= sizeSlider.MinValue() && sizeValue <= sizeSlider.MaxValue() {
				sizeSlider.UpdateValue(sizeValue)
				l.SetCellSize(int(sizeValue))
				l.Render()
				s.OpenSettings()
			}
		}
	}
}
