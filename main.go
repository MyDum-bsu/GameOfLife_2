package main

import (
	"GameOfLife/graphics"
	"GameOfLife/life"
	"golang.org/x/image/colornames"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	width    = 1920
	height   = 1080
	cellSize = 3
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Conway's Game of Life",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetMonitor(pixelgl.PrimaryMonitor())

	imd := imdraw.New(nil)
	l := life.NewLife(height/cellSize, width/cellSize)
	l.Seed()
	for !win.Closed() {
		win.Clear(colornames.Black)
		graphics.Render(win, imd, l, cellSize)
		win.Update()
		time.Sleep(time.Second / 24)
	}
}

func main() {
	pixelgl.Run(run)
}
