package main

import (
	"github.com/eyebrow-fish/smith"
	"github.com/hajimehoshi/ebiten"
	"log"
)

const (
	windowWidth  int = 640
	windowHeight int = 480
)

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	if err := ebiten.RunGame(&smith.Game{InputState: &smith.InputState{}}); err != nil {
		log.Fatalf("failure running game: %v", err)
	}
}
