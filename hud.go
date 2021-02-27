package smith

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type Hud struct {
	heartSprite *ebiten.Image
}

func NewHud(heartSprite []byte) (*Hud, error) {
	spriteMapImage, _, err := image.Decode(bytes.NewReader(heartSprite))
	if err != nil {
		return nil, err
	}
	spriteMap, err := ebiten.NewImageFromImage(spriteMapImage, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}
	return &Hud{heartSprite: spriteMap}, nil
}

func (h *Hud) draw(screen *ebiten.Image, player Player) error {
	for i := 0; i < int(player.health); i++ {
		_, height := screen.Size()
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(i*18+4), float64(height-22))
		if err := screen.DrawImage(h.heartSprite, options); err != nil {
			return err
		}
	}
	return nil
}
