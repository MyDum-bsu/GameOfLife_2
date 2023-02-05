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
	atlas            = text.NewAtlas(basicfont.Face7x13, text.ASCII)
	title            *text.Text
	txt              *text.Text
	sizeSlider       *slider.Slider
	densitySlider    *slider.Slider
	redColorSlider   *slider.Slider
	greenColorSlider *slider.Slider
	blueColorSlider  *slider.Slider
)

type Settings struct {
	win    *pixelgl.Window
	imd    *imdraw.IMDraw
	center pixel.Vec
}

func NewSettings(win *pixelgl.Window, imd *imdraw.IMDraw, center pixel.Vec) *Settings {
	baseVec := center.Sub(pixel.V(160, -150))
	deltaVec := pixel.V(0, 40)
	var r float64 = 8
	var l float64 = 130
	lineC := pixel.RGB(0, 0, 0)
	circleC := pixel.RGB(1, 0, 0)
	sizeSlider = slider.NewSlider(baseVec, r, l, 5, 3, 300, lineC, circleC)
	densitySlider = slider.NewSlider(baseVec.Sub(deltaVec), r, l, 0.25, 0, 1, lineC, circleC)
	redColorSlider = slider.NewSlider(densitySlider.Position().Sub(deltaVec), r, l, 0, 0, 1, lineC, circleC)
	greenColorSlider = slider.NewSlider(redColorSlider.Position().Sub(deltaVec), r, l, 1, 0, 1, lineC, circleC)
	blueColorSlider = slider.NewSlider(greenColorSlider.Position().Sub(deltaVec), r, l, 0, 0, 1, lineC, circleC)
	return &Settings{
		win:    win,
		imd:    imd,
		center: center,
	}
}

func (s *Settings) OpenSettings() {
	s.drawRoundRect(400, 400, 60, 60)
	s.drawSliders()
	s.imd.Draw(s.win)
	s.drawText()
}

func (s *Settings) drawText() {
	title = text.New(pixel.V(s.center.X-160, s.center.Y+190), atlas)
	title.Color = colornames.Darkslategray
	fmt.Fprintln(title, "Settings")
	txt = text.New(pixel.V(s.center.X+30, s.center.Y+145), atlas)

	txt.Color = colornames.Black
	txt.LineHeight = 20
	fmt.Fprintln(txt, "resolution")
	fmt.Fprintln(txt, "density")
	fmt.Fprintln(txt, "red")
	fmt.Fprintln(txt, "green")
	fmt.Fprintln(txt, "blue")
	txt.Draw(s.win, pixel.IM.Scaled(txt.Orig, 2))
	title.Draw(s.win, pixel.IM.Scaled(title.Orig, 6))
}

func (s *Settings) drawRoundRect(width, height, rx, ry float64) {
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
		if s.moveSlider(sizeSlider) {
			l.SetCellSize(int(sizeSlider.Value()), densitySlider.Value())
		}
		if s.moveSlider(densitySlider) {
			l.SetDensity(densitySlider.Value())
			l.Render()
		} else {
			if s.moveSlider(redColorSlider) {
				l.SetRed(redColorSlider.Value())
			}
			if s.moveSlider(greenColorSlider) {
				l.SetGreen(greenColorSlider.Value())
			}
			if s.moveSlider(blueColorSlider) {
				l.SetBlue(blueColorSlider.Value())
			}
			l.Render()
		}
	}
}

func (s *Settings) moveSlider(sl *slider.Slider) bool {
	mousePos := s.win.MousePosition()
	if math.Abs(mousePos.Y-sl.Position().Y) < 20 &&
		((mousePos.X-sl.Position().X < sl.Length()) || (sl.Position().X-mousePos.X > 0)) {
		var value = (s.win.MousePosition().X - sl.Position().X) / sl.Length() * (sl.MaxValue() - sl.MinValue())
		return sl.UpdateValue(value)
	}
	return false
}

func (s *Settings) drawSliders() {
	sizeSlider.Draw(s.imd)
	densitySlider.Draw(s.imd)
	redColorSlider.Draw(s.imd)
	greenColorSlider.Draw(s.imd)
	blueColorSlider.Draw(s.imd)
}
