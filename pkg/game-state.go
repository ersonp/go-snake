package pkg

var (
	pointsChan = make(chan int)
)

type GameState struct {
	board *board
	score int
}

func initialScore() int {
	return 0
}

func initialBoard(h, w int) *board {
	return newBoard(pointsChan, h, w)
}

func NewGame(h, w int) *GameState {
	return &GameState{
		board: initialBoard(h, w),
		score: initialScore(),
	}
}
func (g *GameState) Start() {
	if err := g.render(); err != nil {
		panic(err)
	}

}
