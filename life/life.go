package life

import (
	"GameOfLife/universe"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"sync"
)

type Life struct {
	win           *pixelgl.Window
	imd           *imdraw.IMDraw
	aliveColor    pixel.RGBA
	cellSize      int
	u             *universe.Universe
	previousState []byte
}

func NewLife(win *pixelgl.Window, aColor pixel.RGBA, cellSize, width, height int) *Life {
	u := universe.NewUniverse(height, width)
	u.Seed()
	imd := imdraw.New(nil)
	return &Life{
		win:           win,
		imd:           imd,
		aliveColor:    aColor,
		cellSize:      cellSize,
		u:             u,
		previousState: u.State(),
	}
}

func (l *Life) Render() {
	l.win.Clear(colornames.Black)
	l.imd.Clear()
	//currentState := l.u.State()
	//if !bytes.Equal(currentState, l.previousState) {

	var wg sync.WaitGroup
	partHeight := 2
	partWidth := 2
	for i := 0; i < l.u.Height(); i += partHeight {
		for j := 0; j < l.u.Width(); j += partWidth {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				for x := i; x < i+partHeight && x < l.u.Height(); x++ {
					for y := j; y < j+partWidth && y < l.u.Width(); y++ {
						if l.u.IsAlive(x, y) {
							l.drawRect(x, y)
						}
					}
				}
			}(i, j)
		}
	}

	wg.Wait()

	//for i := 0; i < l.u.Height(); i++ {
	//	for j := 0; j < l.u.Width(); j++ {
	//		if l.u.IsAlive(i, j) {
	//			l.drawRect(i, j)
	//		}
	//	}
	//}

	//l.previousState = currentState

	l.imd.Draw(l.win)
	//}
	l.win.Update()
	l.u.Step()
}

func (l *Life) drawRect(i, j int) {
	l.imd.Color = l.aliveColor
	l.imd.Push(pixel.V(float64(j*l.cellSize), float64(i*l.cellSize)))
	l.imd.Push(pixel.V(float64((j+1)*l.cellSize), float64((i+1)*l.cellSize)))
	l.imd.Rectangle(0)
}
