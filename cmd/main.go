package main

import (
	"github.com/eyebrow-fish/smith"
	"github.com/hajimehoshi/ebiten"
	"log"
)

const (
	windowWidth  int = 640
	windowHeight int = 480
	gameScale    int = 2
)

func main() {
	player, err := smith.NewPlayer()
	if err != nil {
		log.Fatalf("failed to load sprite: %v", err)
	}

	hud, err := smith.NewHud()
	if err != nil {
		log.Fatalf("failed to load sprite: %v", err)
	}

	world, err := smith.NewWorld()
	if err != nil {
		log.Fatalf("failed to load world sprite map: %v", err)
	}

	options := smith.GameOptions{Scale: gameScale}
	game := &smith.Game{
		Player:  *player,
		Hud:     *hud,
		World:   *world,
		Options: options,
	}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("smith")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failure running game: %v", err)
	}
}
