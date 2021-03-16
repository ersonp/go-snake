package pkg

import (
	"math/rand"
	"time"
)

type board struct {
	height     int
	width      int
	pointsChan chan (int)
}

func newBoard(p chan (int), h, w int) *board {
	rand.Seed(time.Now().UnixNano())

	a := &board{
		height:     h,
		width:      w,
		pointsChan: p,
	}
	return a
}
