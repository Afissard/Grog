package main

import (
	"Grog/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &game.Game{}

	ebiten.SetWindowSize(900, 600)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Grog")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
