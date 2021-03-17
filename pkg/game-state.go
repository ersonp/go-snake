package pkg

import (
	"math"
	"time"

	"github.com/ersonp/go-snake/cmd"
)

var (
	keyboardEventsChan = make(chan keyboardEvent)
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

func (g *GameState) retry(h, w int) {
	g.board = initialBoard(h, w)
	g.score = initialScore()
}

func (g *GameState) Start() {
	go listenToKeyboard(keyboardEventsChan)
	for {
		select {
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				g.board.snake.changeDirection(d)
			case RETRY:
				g.retry(g.board.height, g.board.width)
			case END:
				break
			}
		default:
			if err := g.board.moveSnake(); err != nil {
				g.IsOver = true
			}
			time.Sleep(g.moveInterval())

		}
		cmd.CallClear()
		if err := g.render(); err != nil {
			panic(err)
		}
	}
}
