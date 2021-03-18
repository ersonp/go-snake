package pkg

import (
	"fmt"
	"strings"
)

type screen struct {
	cells [][]string
}

const (
	snakeSymbol    = "◻"
	horizontalLine = "█"
	verticalLine   = "|"
	emptySymbol    = " "
	fieldTop       = 1
	fieldLeft      = 1
	move           = "Move:"
	moveUsage      = "W,A,S,D"
	reset          = "Reset:"
	resetUsage     = "R"
	close          = "Close:"
	closeUsage     = "ESC"
	foodSymbol     = "⍟"
	score          = "Score: "
	round          = "Round: "
	snakeLen       = "Snake Length: "
	snakeLoc       = "Snake Location: "
	gameOver       = "Game over!"
)

func (g *GameState) render() error {
	ascii := ""
	m := g.generateScreen()
	for _, row := range m.cells {
		ascii += strings.Join(row, "") + "\n"
	}
	fmt.Print(ascii)
	return nil
}

func (g *GameState) generateScreen() *screen {
	m := new(screen)
	renderHelp()
	renderInfo(g.score, g.round, g.board.snake.length, g.board.snake.head())
	m.renderArena(g.board, g)
	m.renderSnake(g.board.snake)
	m.renderFood(g.board.food.x, g.board.food.y, g.board.height, g.board.width)
	return m
}

func (m *screen) renderArena(a *board, g *GameState) {
	// Add horizontal line on top
	m.cells = append(m.cells, strings.Split(verticalLine+strings.Repeat(horizontalLine, a.width)+verticalLine, ""))

	// Render battlefield
	for i := 0; i < a.height; i++ {
		if i == 1 && g.isOver {
			row := []string{verticalLine, emptySymbol}
			for _, r := range gameOver {
				row = append(row, string(r))
			}
			for j := len(gameOver) + 1; j < a.width; j++ {
				row = append(row, emptySymbol)
			}
			row = append(row, verticalLine)
			m.cells = append(m.cells, row)
		} else {
			m.cells = append(m.cells, strings.Split(verticalLine+strings.Repeat(emptySymbol, a.width)+verticalLine, ""))
		}
	}

	// Add horizontal line on bottom
	m.cells = append(m.cells, strings.Split(verticalLine+strings.Repeat(horizontalLine, a.width)+verticalLine, ""))
}

func (m *screen) renderSnake(s *snake) {
	for _, b := range s.body {
		m.cells[b.x+fieldTop][b.y+fieldLeft] = snakeSymbol
	}
}

func (m *screen) renderFood(x, y, h, w int) {
	m.cells[x+fieldTop][y+fieldLeft] = foodSymbol
}

func renderHelp() {
	fmt.Printf("%v %v    %v %v    %v %v\n", move, moveUsage, reset, resetUsage, close, closeUsage)
}

func renderInfo(scoreVal, roundVal, snakeLenVal int, snakeCood coord) {
	fmt.Printf("%s %d    %s %d\n", score, scoreVal, round, roundVal)
	fmt.Printf("%s %d    %s %d\n", snakeLen, snakeLenVal, snakeLoc, snakeCood)
}
