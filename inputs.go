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

func (i *InputState) collectInputs() {
	i.released = []ebiten.Key{}
	i.collectKeys(
		ebiten.KeyF1,
		ebiten.KeyW,
		ebiten.KeyA,
		ebiten.KeyS,
		ebiten.KeyD,
		ebiten.KeySpace,
	)
}

func (i *InputState) collectKeys(keys ...ebiten.Key) {
	for _, key := range keys {
		keyIndex := i.rawIndex(key)
		if ebiten.IsKeyPressed(key) {
			if keyIndex < 0 {
				i.raw = append(i.raw, key)
			}
		} else if keyIndex > -1 {
			i.raw = append(i.raw[:keyIndex], i.raw[keyIndex+1:]...)
			i.released = append(i.released, key)
		}
	}
}

func (i InputState) String() string {
	return fmt.Sprintf("raw: %v\nreleased: %v", i.raw, i.released)
}
