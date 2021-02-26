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
	ebiten.SetWindowSize(windowWidth, windowHeight)
	if err := ebiten.RunGame(&smith.Game{Player: *player}); err != nil {
		log.Fatalf("failure running game: %v", err)
	}
}
