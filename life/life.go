package life

import (
	"GameOfLife/universe"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type Life struct {
	win                   *pixelgl.Window
	imd                   *imdraw.IMDraw
	aliveColor, deadColor pixel.RGBA
	cellSize              int
	u                     *universe.Universe
}

func NewLife(win *pixelgl.Window, imd *imdraw.IMDraw, aColor, dColor pixel.RGBA, cellSize int) *Life {
	width := int(win.Bounds().W()) / cellSize
	height := int(win.Bounds().H()) / cellSize
	u := universe.NewUniverse(width, height)
	u.Seed()
	return &Life{
		win:        win,
		imd:        imd,
		aliveColor: aColor,
		deadColor:  dColor,
		cellSize:   cellSize,
		u:          u,
	}
}

//func (l *Life) Render() {
//	type chunk struct {
//		x, y int
//		imd  *imdraw.IMDraw
//	}
//
//	const chunkSize = 10
//
//	chunks := make(chan chunk)
//	done := make(chan bool)
//
//	for i := 0; i < l.u.Width(); i += chunkSize {
//		for j := 0; j < l.u.Height(); j += chunkSize {
//			go func(x, y int) {
//				imd := imdraw.New(nil)
//				for i := x; i < x+chunkSize && i < l.u.Width(); i++ {
//					for j := y; j < y+chunkSize && j < l.u.Height(); j++ {
//						if l.u.IsAlive(i, j) {
//							l.drawRect(i, j, l.aliveColor)
//						}
//					}
//				}
//				chunks <- chunk{x, y, imd}
//			}(i, j)
//		}
//	}
//
//	go func() {
//		for i := 0; i < (l.u.Width()*l.u.Height())/(chunkSize*chunkSize); i++ {
//			c := <-chunks
//			c.imd.Draw(l.win)
//		}
//		done <- true
//	}()
//
//	l.u.Step()
//
//	<-done
//}

func (l *Life) Render() {
	for i := 0; i < l.u.Width(); i++ {
		for j := 0; j < l.u.Height(); j++ {
			if l.u.IsAlive(i, j) {
				l.drawRect(i, j, l.aliveColor)
			}
		}
	}
	l.imd.Draw(l.win)
}

func (l *Life) Step() {
	l.u.Step()
}

func (l *Life) drawRect(i, j int, color pixel.RGBA) {
	l.imd.Color = color
	l.imd.Push(pixel.V(float64(i*l.cellSize+1), float64(j*l.cellSize+1)))
	l.imd.Push(pixel.V(float64((i+1)*l.cellSize-1), float64((j+1)*l.cellSize-1)))
	l.imd.Rectangle(0)
}

func (l *Life) Erase() {
	l.u.Erase()
	l.Render()
}

func (l *Life) HandleInput() {
	pos := l.win.MousePosition()
	x := int(math.Floor(pos.X / float64(l.cellSize)))
	y := int(math.Floor(pos.Y / float64(l.cellSize)))
	if l.validateMousePosition(x, y) {
		if l.win.Pressed(pixelgl.MouseButtonLeft) {
			l.u.Alive(x, y)
			l.drawRect(x, y, l.aliveColor)
		}
		if l.win.Pressed(pixelgl.MouseButtonRight) {
			l.u.Dead(x, y)
			l.drawRect(x, y, l.deadColor)
		}
	}
}

func (l *Life) validateMousePosition(x, y int) bool {
	return x >= 0 && y >= 0 && x < l.u.Width() && y < l.u.Height()
}
