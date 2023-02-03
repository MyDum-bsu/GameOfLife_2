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

	aliveColor := pixel.RGB(1, 0, 0)
	deadColor := pixel.RGB(0, 0, 0)
	imd := imdraw.New(nil)
	l := life.NewLife(win, imd, aliveColor, deadColor, cellSize)
	s := settings2.NewSettings(pixel.V(width/2, height/2))

	//basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	//basicTxt := text.New(pixel.V(500, 700), basicAtlas)
	//basicTxt.Color = pixel.RGBA{R: 1, G: 1, B: 1, A: 1}
	//fmt.Fprintln(basicTxt, "Hello, World!")

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
		//basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 4))
		imd.Draw(win)
		win.Update()
		time.Sleep(time.Second / 120)
	}
}

func main() {
	pixelgl.Run(run)
}
