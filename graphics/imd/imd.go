package imd

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func DrawRect(imd *imdraw.IMDraw, color pixel.RGBA, i, j, cellSize int) {
	imd.Color = color
	imd.Push(pixel.V(float64(j*cellSize), float64(i*cellSize)))
	imd.Push(pixel.V(float64((j+1)*cellSize), float64((i+1)*cellSize)))
	imd.Rectangle(0)
}
