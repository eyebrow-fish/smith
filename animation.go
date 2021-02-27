package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type Animation struct {
	spriteMap *ebiten.Image
	frame     int
	maxFrame  int
	debounce  int
	buffer    int
}

func (a *Animation) update() (*ebiten.Image, error) {
	a.buffer++
	if a.buffer >= a.debounce {
		a.buffer = 0
		if a.frame < a.maxFrame {
			a.frame++
		} else {
			a.frame = 0
		}
	}
	spriteY := a.frame * 10
	spriteTile := a.spriteMap.SubImage(image.Rect(10, spriteY, 0, spriteY+10)).(*ebiten.Image)
	return spriteTile, nil
}

func (a Animation) String() string {
	return fmt.Sprintf(
		"  frame: %d, maxFrame: %d, buffer: %d\n  debounce: %d",
		a.frame,
		a.maxFrame,
		a.buffer,
		a.debounce,
	)
}
