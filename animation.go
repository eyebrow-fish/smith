package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type animation struct {
	spriteMap *ebiten.Image
	frame     int
	maxFrame  int
	debounce  int
	buffer    int
	direction animationDirection
	paused    bool
}

func (a *animation) update() (*ebiten.Image, error) {
	if a.paused {
		a.frame = 0
	} else {
		a.buffer++
		if a.buffer >= a.debounce {
			a.buffer = 0
			if a.direction == down {
				if a.frame < a.maxFrame {
					a.frame++
				} else {
					a.frame = 0
				}
			} else if a.direction == up {
				if a.frame > 0 {
					a.frame--
				} else {
					a.frame = a.maxFrame
				}
			}
		}
	}

	spriteY := a.frame * SpriteSize
	spriteTile := a.spriteMap.SubImage(image.Rect(0, spriteY, SpriteSize, spriteY+SpriteSize)).(*ebiten.Image)

	return spriteTile, nil
}

func (a animation) String() string {
	return fmt.Sprintf(
		"  frame: %d, maxFrame: %d, buffer: %d\n debounce: %d",
		a.frame,
		a.maxFrame,
		a.buffer,
		a.debounce,
	)
}

type animationDirection uint8

const (
	down animationDirection = iota
	up
)
