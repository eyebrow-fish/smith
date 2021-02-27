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
	gameScale    int = 2
)

//go:embed smith_proto_1.png
var playerSprite []byte

//go:embed heart.png
var heartSprite []byte

func main() {
	player, err := smith.NewPlayer(playerSprite)
	if err != nil {
		log.Fatalf("failed to load sprite: %v", err)
	}
	hud, err := smith.NewHud(heartSprite)
	if err != nil {
		log.Fatalf("failed to load sprite: %v", err)
	}
	options := smith.GameOptions{Scale: gameScale}
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("smith")
	if err := ebiten.RunGame(smith.NewGame(*player, *hud, options)); err != nil {
		log.Fatalf("failure running game: %v", err)
	}
}
