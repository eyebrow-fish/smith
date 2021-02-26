package smith

import (
	"image"
)

type Animation struct {
	spriteMap image.Image
	frame     int
	maxFrame  int
	debounce  int
	buffer    int
}

func (a *Animation) tick() {
	a.buffer++
	if a.buffer >= a.debounce {
		a.buffer = 0
		if a.frame < a.maxFrame {
			a.frame++
		} else {
			a.frame = 0
		}
	}
}
