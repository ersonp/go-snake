package pkg

import (
	"fmt"
	"strings"
)

type screen struct {
	cells [][]string
}

const (
	snakeSymbol    = "*"
	horizontalLine = "-"
	verticalLine   = "|"
	emptySymbol    = " "
	fieldTop       = 7
	fieldLeft      = 1
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
	m.renderArena(g.board, g)
	m.renderSnake(g.board.snake)
	return m
}

func (m *screen) renderArena(a *board, g *GameState) {
	// Add horizontal line on top
	m.cells = append(m.cells, strings.Split(verticalLine+strings.Repeat(horizontalLine, a.width)+verticalLine, ""))

	// Render battlefield
	for i := 0; i < a.height; i++ {
		m.cells = append(m.cells, strings.Split(verticalLine+strings.Repeat(emptySymbol, a.width)+verticalLine, ""))
	}

	// Add horizontal line on bottom
	m.cells = append(m.cells, strings.Split(verticalLine+strings.Repeat(horizontalLine, a.width)+verticalLine, ""))
}

func (m *screen) renderSnake(s *snake) {
	for _, b := range s.body {
		m.cells[b.x+fieldTop][b.y+fieldLeft] = snakeSymbol
	}
}
