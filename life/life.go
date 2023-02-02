package life

import "math/rand"

type Life struct {
	board [][]bool
	w, h  int
}

func NewLife(height, width int) *Life {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}
	return &Life{board: board, h: height, w: width}
}

func (l *Life) Seed() {
	for i := 0; i < (l.h * l.w / 4); i++ {
		l.board[rand.Intn(l.h)][rand.Intn(l.w)] = true
	}
}

func (l *Life) Step() {
	newBoard := make([][]bool, l.h)
	for i := range newBoard {
		newBoard[i] = make([]bool, l.w)
	}

	for i := 0; i < l.h; i++ {
		for j := 0; j < l.w; j++ {
			onCount := l.countOnNeighbors(i, j)
			if l.board[i][j] && (onCount == 2 || onCount == 3) {
				newBoard[i][j] = true
			}
			if !l.board[i][j] && onCount == 3 {
				newBoard[i][j] = true
			}
		}
	}
	l.board = newBoard
}

func (l *Life) countOnNeighbors(x, y int) int {
	onCount := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			xx, yy := (x+dx+l.h)%l.h, (y+dy+l.w)%l.w
			if l.board[xx][yy] {
				onCount++
			}
		}
	}
	return onCount
}

func (l *Life) IsAlive(x, y int) bool {
	return l.board[x][y]
}

func (l *Life) Width() int {
	return l.w
}

func (l *Life) Height() int {
	return l.h
}
