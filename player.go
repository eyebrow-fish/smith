package smith

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type Player struct {
	position  Vertex2
	animation Animation
	speed     float64
}

func NewPlayer(sprite []byte) (*Player, error) {
	spriteImage, _, err := image.Decode(bytes.NewReader(sprite))
	if err != nil {
		return nil, err
	}
	animation := Animation{spriteMap: spriteImage, maxFrame: 1, debounce: 10}
	return &Player{animation: animation, speed: 4}, nil
}

func (p *Player) draw(screen *ebiten.Image) error {
	spriteImage, err := ebiten.NewImageFromImage(p.animation.spriteMap, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.position.x, p.position.y)
	spriteTile := spriteImage.SubImage(image.Rect(32, 32, 0, 0)).(*ebiten.Image)
	return screen.DrawImage(spriteTile, options)
}

func (p *Player) String() string {
	return fmt.Sprintf("position: [%.2f, %.2f]\n", p.position.x, p.position.y)
}
