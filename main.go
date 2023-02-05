package main

import (
	"GameOfLife/life"
	settings2 "GameOfLife/settings"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

const (
	width    = 1920
	height   = 1080
	cellSize = 4
)

var (
	aliveColor = pixel.RGB(0, 1, 0)
	deadColor  = pixel.RGB(0, 0, 0)
	paused     = false
	settings   = false
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Conway's Game of Universe",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetMonitor(pixelgl.PrimaryMonitor())

	imd := imdraw.New(nil)
	l := life.NewLife(win, imd, aliveColor, deadColor, cellSize)
	s := settings2.NewSettings(win, imd, pixel.V(width/2, height/2))

	for !win.Closed() {
		if settings {
			s.Listen(l)
			s.OpenSettings()
		} else {
			l.HandleInput()
			if !paused {
				l.Render()
				l.Step()
			} else {
				imd.Draw(win)
			}
		}
		checkPressedButtons(win, l, s)
		win.Update()
		time.Sleep(time.Second / 120)
	}
}

func checkPressedButtons(win *pixelgl.Window, l *life.Life, s *settings2.Settings) {
	if win.JustPressed(pixelgl.KeyEscape) {
		settings = !settings
		paused = !paused
		//s.OpenSettings()
	}

	if win.JustPressed(pixelgl.KeySpace) {
		if !settings {
			paused = !paused
		}
	}
	if win.JustPressed(pixelgl.KeyE) {
		l.Erase()
	}
}

func main() {
	pixelgl.Run(run)
}
