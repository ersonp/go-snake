package pkg

import (
	"math/rand"
)

type board struct {
	food       *food
	snake      *snake
	hasFood    func(*board, coord) bool
	height     int
	width      int
	pointsChan chan (int)
}

func newBoard(s *snake, p chan (int), h, w int) *board {
	// rand.Seed(time.Now().UnixNano())

	a := &board{
		snake:      s,
		height:     h,
		width:      w,
		pointsChan: p,
		hasFood:    hasFood,
	}
	a.placeFood()
	return a
}

func (a *board) placeFood() {
	var x, y int

	for {
		x = rand.Intn(a.height)
		y = rand.Intn(a.width)

		if !a.isOccupied(coord{x: x, y: y}) {
			break
		}
	}

	a.food = newFood(x, y)
}

func (a *board) moveSnake() error {
	if err := a.snake.move(); err != nil {
		return err
	}

	if a.snakeLeftArena() {
		return a.snake.die()
	}

	if a.hasFood(a, a.snake.head()) {
		go a.addPoints(a.food.points)
		a.snake.length++
		a.placeFood()
	}

	return nil
}

func (a *board) snakeLeftArena() bool {
	h := a.snake.head()
	return h.x > a.height-1 || h.y > a.width-1 || h.x < 0 || h.y < 0
}

func (a *board) addPoints(p int) {
	a.pointsChan <- p
}

func hasFood(a *board, c coord) bool {
	return c.x == a.food.x && c.y == a.food.y
}

func (a *board) isOccupied(c coord) bool {
	return a.snake.hits(c)
}
