package main

import (
	_ "embed"
	"github.com/eyebrow-fish/smith"
	"github.com/hajimehoshi/ebiten"
	"log"
)

const (
	windowWidth  int = 640
	windowHeight int = 480
)

//go:embed smith_proto_1.png
var playerSprite []byte

func main() {
	player, err := smith.NewPlayer(playerSprite)
	if err != nil {
		log.Fatalf("failed to load sprite: %v", err)
	}
	options := smith.GameOptions{Scale: 2}
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("smith")
	if err := ebiten.RunGame(smith.NewGame(*player, options)); err != nil {
		log.Fatalf("failure running game: %v", err)
	}
}
