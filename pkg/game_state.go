package pkg

import (
	"math"
	"time"

	"github.com/ersonp/go-snake/cmd"
)

var (
	pointsChan         = make(chan int)
	keyboardEventsChan = make(chan keyboardEvent)
)

type GameState struct {
	board  *board
	score  int
	round  int
	isOver bool
}

func initialZero() int {
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
	return newBoard(initialSnake(), pointsChan, h, w)
}

func NewGame(h, w int) *GameState {
	return &GameState{
		board: initialBoard(h, w),
		score: initialZero(),
		round: initialZero(),
	}
}

func (g *GameState) moveInterval() time.Duration {
	ms := 400 - math.Max(float64(g.score), 100)
	return time.Duration(ms) * time.Millisecond
}

func (g *GameState) retry(h, w int) {
	g.board = initialBoard(h, w)
	g.score = initialZero()
	g.round = initialZero()
	g.isOver = false
}

func (g *GameState) Start() {
	go listenToKeyboard(keyboardEventsChan)
mainloop:
	for {
		select {
		case p := <-pointsChan:
			g.addPoints(p)
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				g.addRound()
				d := keyToDirection(e.key)
				g.board.snake.changeDirection(d)
			case RETRY:
				g.retry(g.board.height, g.board.width)
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if err := g.board.moveSnake(); err != nil {
					g.end()
				}
			}
			cmd.CallClear()

			if err := g.render(); err != nil {
				panic(err)
			}
			time.Sleep(g.moveInterval())
		}
	}
}

func (g *GameState) addPoints(p int) {
	g.score += p
}

func (g *GameState) addRound() {
	g.round++
}

func (g *GameState) end() {
	g.isOver = true
}
