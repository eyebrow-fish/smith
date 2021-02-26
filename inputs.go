package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

type InputState struct {
	debugMode bool
	raw       []ebiten.Key
	released  []ebiten.Key
}

func (i *InputState) rawIndex(key ebiten.Key) int {
	for i, v := range i.raw {
		if key == v {
			return i
		}
	}
	return -1
}

func (i *InputState) hasReleased(key ebiten.Key) bool {
	for _, v := range i.released {
		if key == v {
			return true
		}
	}
	return false
}

func (i *InputState) handleInputs() {
	i.released = []ebiten.Key{}
	f1Key := i.rawIndex(ebiten.KeyF1)
	if ebiten.IsKeyPressed(ebiten.KeyF1) {
		if f1Key < 0 {
			i.raw = append(i.raw, ebiten.KeyF1)
		}
	} else if f1Key > -1 {
		i.raw = append(i.raw[:f1Key], i.raw[f1Key+1:]...)
		i.released = append(i.released, ebiten.KeyF1)
	}
}

func (i *InputState) String() string {
	return fmt.Sprintf("raw: %v\nreleased: %v\n", i.raw, i.released)
}
