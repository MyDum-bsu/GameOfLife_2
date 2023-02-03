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
	cellSize = 10
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
	aliveColor := pixel.RGB(1, 0, 0)
	deadColor := pixel.RGB(0, 0, 0)
	imd := imdraw.New(nil)
	l := life.NewLife(win, imd, aliveColor, deadColor, cellSize)
	s := settings2.NewSettings(win)

	paused := false
	settings := false
	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyEscape) {
			settings = !settings
			paused = !paused
			s.OpenSettings(imd)
		}
		l.HandleInput()
		if win.JustPressed(pixelgl.KeySpace) {
			paused = !paused
		}
		if win.JustPressed(pixelgl.KeyE) {
			l.Erase()
		}
		if !paused {
			win.Clear(deadColor)
			imd.Clear()
			l.Render()
		}
		imd.Draw(win)
		win.Update()
		time.Sleep(time.Second / 120)
	}
}

func main() {
	pixelgl.Run(run)
}
