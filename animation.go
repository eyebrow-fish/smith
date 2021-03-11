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
	direction animationDirection
	paused    bool
}

func (a *Animation) update() (*ebiten.Image, error) {
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

	spriteY := a.frame * 10
	spriteTile := a.spriteMap.SubImage(image.Rect(0, spriteY, 10, spriteY+10)).(*ebiten.Image)

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

type animationDirection uint8

const (
	down animationDirection = iota
	up
)
