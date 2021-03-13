package main

import (
	"github.com/eyebrow-fish/smith"
	"github.com/hajimehoshi/ebiten"
	"log"
)

const (
	windowWidth  int = 640
	windowHeight int = 480
	windowScale  int = 2
)

func main() {
	options := smith.WindowOptions{Scale: windowScale}
	game := smith.NewGame(options)

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("smith")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failure running game: %v", err)
	}
}
