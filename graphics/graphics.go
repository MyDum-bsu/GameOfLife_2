package graphics

import (
	imd2 "GameOfLife/graphics/imd"
	"GameOfLife/life"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func Render(win *pixelgl.Window, imd *imdraw.IMDraw, l *life.Life, cellSize int) {
	imd.Clear()

	for i := 0; i < l.Height(); i++ {
		for j := 0; j < l.Width(); j++ {
			if l.IsAlive(i, j) {
				imd2.DrawRect(imd, pixel.RGB(1, 0, 1), i, j, cellSize)
			}
		}
	}
	imd.Draw(win)
	l.Step()
}
