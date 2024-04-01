package main

import (
	"game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
