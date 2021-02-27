package smith

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image"
	"math"
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
	roundedHealth := math.Round(float64(player.health))
	for i := 0; i < int(roundedHealth); i += 2 {
		_, height := screen.Size()
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(i/2*10+4), float64(height-14))
		sprite := h.heartSprite
		if i == int(roundedHealth-1) && int(roundedHealth)%2 == 1 {
			sprite = h.heartSprite.SubImage(image.Rect(4, 8, 0, 0)).(*ebiten.Image)
		}
		if err := screen.DrawImage(sprite, options); err != nil {
			return err
		}
	}
	return nil
}
