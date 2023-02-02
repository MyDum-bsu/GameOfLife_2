package main

import (
	"GameOfLife/life"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

const (
	width    = 1920
	height   = 1080
	cellSize = 5
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
	win.SetCursorVisible(false)
	win.SetMonitor(pixelgl.PrimaryMonitor())
	aliveColor := pixel.RGB(1, 0, 0)
	deadColor := pixel.RGB(0, 0, 0)
	l := life.NewLife(win, aliveColor, deadColor, cellSize, width/cellSize, height/cellSize)

	paused := false
	for !win.Closed() {
		if win.JustPressed(pixelgl.KeySpace) {
			paused = !paused
		}
		if !paused {
			l.Render()
		}
		win.Update()
		time.Sleep(time.Second / 120)
	}
}

func main() {
	pixelgl.Run(run)
}
