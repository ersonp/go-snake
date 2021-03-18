package pkg

import (
	"testing"
)

func BenchmarkRender(b *testing.B) {
	game := NewGame(20, 30)

	for n := 0; n < b.N; n++ {
		err := game.render()
		if err != nil {
			panic(err)
		}
	}
}
