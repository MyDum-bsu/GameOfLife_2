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
	cellSize = 3
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

	l := life.NewLife(win, pixel.RGB(1, 0, 1), 3, width/cellSize, height/cellSize)

	fps := time.Tick(time.Second / 30)
	for !win.Closed() {
		select {
		case <-fps:
			l.Render()
		}
	}
}

func main() {
	pixelgl.Run(run)
}
