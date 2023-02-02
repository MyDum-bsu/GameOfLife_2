package universe

import "math/rand"

type Universe struct {
	board [][]bool
	w, h  int
}

func NewUniverse(height, width int) *Universe {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}
	return &Universe{board: board, h: height, w: width}
}

func (u *Universe) Seed() {
	for i := 0; i < (u.h * u.w / 4); i++ {
		u.board[rand.Intn(u.h)][rand.Intn(u.w)] = true
	}
}

func (u *Universe) Step() {
	newBoard := make([][]bool, u.h)
	for i := range newBoard {
		newBoard[i] = make([]bool, u.w)
	}

	for i := 0; i < u.h; i++ {
		for j := 0; j < u.w; j++ {
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
			xx, yy := (x+dx+u.h)%u.h, (y+dy+u.w)%u.w
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

func (u *Universe) Width() int {
	return u.w
}

func (u *Universe) Height() int {
	return u.h
}

//func (u *Universe) State() [][]bool {
//	return u.board
//}

func (u *Universe) State() []byte {
	state := make([]byte, u.Height()*u.Width())
	index := 0
	for i := 0; i < u.Height(); i++ {
		for j := 0; j < u.Width(); j++ {
			if u.IsAlive(i, j) {
				state[index] = 1
			} else {
				state[index] = 0
			}
			index++
		}
	}
	return state
}
