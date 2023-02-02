package life

import (
	"GameOfLife/universe"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Life struct {
	win                   *pixelgl.Window
	imd                   *imdraw.IMDraw
	aliveColor, deadColor pixel.RGBA
	cellSize              int
	u                     *universe.Universe
	previousState         []byte
}

func NewLife(win *pixelgl.Window, aColor, dColor pixel.RGBA, cellSize, width, height int) *Life {
	u := universe.NewUniverse(height, width)
	u.Seed()
	imd := imdraw.New(nil)
	return &Life{
		win:           win,
		imd:           imd,
		aliveColor:    aColor,
		deadColor:     dColor,
		cellSize:      cellSize,
		u:             u,
		previousState: u.State(),
	}
}

func (l *Life) Render() {
	l.win.Clear(l.deadColor)
	l.imd.Clear()
	for i := 0; i < l.u.Height(); i++ {
		for j := 0; j < l.u.Width(); j++ {
			if l.u.IsAlive(i, j) {
				l.drawRect(i, j)
			}
		}
	}
	l.imd.Draw(l.win)
	l.u.Step()
}

func (l *Life) drawRect(i, j int) {
	l.imd.Color = l.aliveColor
	l.imd.Push(pixel.V(float64(j*l.cellSize+1), float64(i*l.cellSize+1)))
	l.imd.Push(pixel.V(float64((j+1)*l.cellSize-1), float64((i+1)*l.cellSize-1)))
	l.imd.Rectangle(0)
}
