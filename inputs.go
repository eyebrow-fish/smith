package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

type inputState struct {
	debugMode bool
	raw       []ebiten.Key
	released  []ebiten.Key
}

func (i *inputState) rawIndex(key ebiten.Key) int {
	for i, v := range i.raw {
		if key == v {
			return i
		}
	}
	return -1
}

func (i *inputState) hasReleased(key ebiten.Key) bool {
	for _, v := range i.released {
		if key == v {
			return true
		}
	}
	return false
}

func (i *inputState) collectInputs() {
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

func (i *inputState) collectKeys(keys ...ebiten.Key) {
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

func (i inputState) String() string {
	return fmt.Sprintf("raw: %v\nreleased: %v", i.raw, i.released)
}
