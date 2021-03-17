package pkg

import (
	"math/rand"
	"time"
)

type coord struct {
	x, y int
}

type board struct {
	snake  *snake
	height int
	width  int
}

func newBoard(s *snake, h, w int) *board {
	rand.Seed(time.Now().UnixNano())

	a := &board{
		snake:  s,
		height: h,
		width:  w,
	}
	return a
}
func (a *board) moveSnake() error {
	if err := a.snake.move(); err != nil {
		return err
	}
	return nil
}
