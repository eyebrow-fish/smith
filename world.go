package smith

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type World struct {
	worldMap *ebiten.Image
	tiles    []tile
}

func NewWorld() (*World, error) {
	spriteMapImage, _, err := image.Decode(bytes.NewReader(WorldSprite))
	if err != nil {
		return nil, err
	}

	spriteMap, err := ebiten.NewImageFromImage(spriteMapImage, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	// Very temporary way of loading the world.
	var groundTiles []tile
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			groundTile := newGroundTile(
				2,
				vertex2{float64(SpriteSize * x), float64(SpriteSize * y)},
			)
			groundTiles = append(groundTiles, groundTile)
		}
	}

	return &World{worldMap: spriteMap, tiles: groundTiles}, nil
}

func (w *World) draw(screen *ebiten.Image) error {
	for _, t := range w.tiles {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(t.position.x, t.position.y)
		options.GeoM.Scale(t.scale, t.scale)

		spriteX := t.row * SpriteSize
		spriteY := t.column * SpriteSize
		tileSprite := w.worldMap.SubImage(image.Rect(
			spriteX,
			spriteY,
			spriteX+SpriteSize,
			spriteY+SpriteSize,
		)).(*ebiten.Image)

		if err := screen.DrawImage(tileSprite, options); err != nil {
			return err
		}
	}

	return nil
}

func newGroundTile(scale float64, position vertex2) tile {
	return tile{scale: scale, position: position}
}

type tile struct {
	row      int
	column   int
	scale    float64
	position vertex2
}
