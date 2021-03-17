package pkg

import (
	"math"
	"time"

	"github.com/ersonp/go-snake/cmd"
)

type GameState struct {
	board  *board
	score  int
	IsOver bool
}

func initialScore() int {
	return 0
}

func initialSnake() *snake {
	return newSnake(RIGHT, []coord{
		coord{
			x: 1,
			y: 1,
		},
		coord{
			x: 1,
			y: 2,
		},
		coord{
			x: 1,
			y: 3,
		},
		coord{
			x: 1,
			y: 4,
		},
	})
}

func initialBoard(h, w int) *board {
	return newBoard(initialSnake(), h, w)
}

func NewGame(h, w int) *GameState {
	return &GameState{
		board: initialBoard(h, w),
		score: initialScore(),
	}
}

func (g *GameState) moveInterval() time.Duration {
	ms := 400 - math.Max(float64(g.score), 100)
	return time.Duration(ms) * time.Millisecond
}

func (g *GameState) Start() {
	for {
		if err := g.board.moveSnake(); err != nil {
			g.IsOver = true
		}
		cmd.CallClear()
		if err := g.render(); err != nil {
			panic(err)
		}

		time.Sleep(g.moveInterval())
	}
}
