package main

import (
	"github.com/ersonp/go-snake/cmd"
	"github.com/ersonp/go-snake/pkg"
)

func main() {
	h, w := cmd.GetBoardSize()
	pkg.NewGame(h, w).Start()
}
