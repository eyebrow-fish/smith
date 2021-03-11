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

func NewWorld(worldSprite []byte) (*World, error) {
	spriteMapImage, _, err := image.Decode(bytes.NewReader(worldSprite))
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
				vertex2{2, 2},
				vertex2{float64(10 * x), float64(10 * y)},
			)
			groundTiles = append(groundTiles, groundTile)
		}
	}

	return &World{worldMap: spriteMap, tiles: groundTiles}, nil
}

func (w *World) draw(screen *ebiten.Image) error {
	for _, tile := range w.tiles {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(tile.position.x, tile.position.y)
		options.GeoM.Scale(tile.scale.x, tile.scale.y)

		spriteX := tile.row * 10
		spriteY := tile.column * 10
		tile := w.worldMap.SubImage(image.Rect(spriteX, spriteY, spriteX+10, spriteY+10)).(*ebiten.Image)

		if err := screen.DrawImage(tile, options); err != nil {
			return err
		}
	}

	return nil
}

func newGroundTile(scale vertex2, position vertex2) tile {
	return tile{scale: scale, position: position}
}

type tile struct {
	row      int
	column   int
	scale    vertex2
	position vertex2
}