package smith

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"math"
)

type hud struct {
	heartSprite *ebiten.Image
}

func newHud() (*hud, error) {
	spriteMapImage, _, err := image.Decode(bytes.NewReader(HeartSprite))
	if err != nil {
		return nil, err
	}

	spriteMap, err := ebiten.NewImageFromImage(spriteMapImage, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	return &hud{heartSprite: spriteMap}, nil
}

func (h *hud) draw(screen *ebiten.Image, p player) error {
	if p.health <= 0 {
		x, y := screen.Size()
		ebitenutil.DebugPrintAt(screen, "x(", x-100, y-80)
		ebitenutil.DebugPrintAt(screen, "space to retry", x-100, y-60)

		return nil
	}

	roundedHealth := math.Round(float64(p.health))
	for i := 0; i < int(roundedHealth); i += 2 {
		_, height := screen.Size()

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(i/2*10+4), float64(height-14))

		switch p.healthStatus {
		case healthy:
			options.ColorM.Scale(0xff, 0x0, 0x0, 0xff)
		case poisoned:
			options.ColorM.Scale(0x0, 0xff, 0x0, 0xff)
		case freezing:
			options.ColorM.Scale(0x0, 0x0, 0xff, 0xff)
		}

		sprite := h.heartSprite
		if i == int(roundedHealth-1) && int(roundedHealth)%2 == 1 {
			sprite = h.heartSprite.SubImage(image.Rect(0, 0, 4, 8)).(*ebiten.Image)
		}

		if err := screen.DrawImage(sprite, options); err != nil {
			return err
		}
	}

	return nil
}
