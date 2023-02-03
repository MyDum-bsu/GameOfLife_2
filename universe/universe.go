package universe

import "math/rand"

type Universe struct {
	board [][]bool
	w, h  int
}

func NewUniverse(width, height int) *Universe {
	board := make([][]bool, width)
	for i := range board {
		board[i] = make([]bool, height)
	}
	return &Universe{board: board, w: width, h: height}
}

func (u *Universe) Seed() {
	for i := 0; i < (u.h * u.w / 4); i++ {
		u.board[rand.Intn(u.w)][rand.Intn(u.h)] = true
	}
}

func (u *Universe) Step() {
	newBoard := make([][]bool, u.w)
	for i := range newBoard {
		newBoard[i] = make([]bool, u.h)
	}

	for i := 0; i < u.w; i++ {
		for j := 0; j < u.h; j++ {
			onCount := u.countOnNeighbors(i, j)
			if u.board[i][j] && (onCount == 2 || onCount == 3) {
				newBoard[i][j] = true
			}
			if !u.board[i][j] && onCount == 3 {
				newBoard[i][j] = true
			}
		}
	}
	u.board = newBoard
}

func (u *Universe) countOnNeighbors(x, y int) int {
	onCount := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			xx, yy := (x+dx+u.w)%u.w, (y+dy+u.h)%u.h
			if u.board[xx][yy] {
				onCount++
			}
		}
	}
	return onCount
}

func (u *Universe) IsAlive(x, y int) bool {
	return u.board[x][y]
}

func (u *Universe) Alive(x, y int) {
	u.board[x][y] = true
}

func (u *Universe) Width() int {
	return u.w
}

func (u *Universe) Height() int {
	return u.h
}

func (u *Universe) Erase() {
	for i := range u.board {
		for j := range u.board[i] {
			u.board[i][j] = false
		}
	}
}
