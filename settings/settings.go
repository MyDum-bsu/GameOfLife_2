package settings

import (
	"GameOfLife/settings/slider"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Settings struct {
	win        *pixelgl.Window
	sizeSlider *slider.Slider
}

func NewSettings(win *pixelgl.Window) *Settings {
	sizeSlider := slider.NewSlider(win.Bounds().W()/2, win.Bounds().H()/2, 8, 100, 5, 2, 100, pixel.RGB(1, 1, 1), pixel.RGB(0, 0, 1))
	return &Settings{
		win:        win,
		sizeSlider: sizeSlider,
	}
}

func (s *Settings) OpenSettings(imd *imdraw.IMDraw) {
	s.sizeSlider.Draw(imd)

}
