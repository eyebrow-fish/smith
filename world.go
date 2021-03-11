package smith

import (
	"bytes"
	"encoding/json"
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

	var wd worldData
	if err := json.Unmarshal(WorldRawData, &wd); err != nil {
		return nil, err
	}

	var tiles []tile
	for _, t := range wd.Tiles {
		tt := wd.TileMap[t.Type]
		position := vertex2{t.Position[0], t.Position[1]}
		tiles = append(
			tiles,
			tile{tt[0], tt[1], wd.Scale, position},
		)
	}

	return &World{worldMap: spriteMap, tiles: tiles}, nil
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

type tile struct {
	column   int
	row      int
	scale    float64
	position vertex2
}

type worldData struct {
	Scale   float64             `json:"scale"`
	Tiles   []tileData          `json:"tiles"`
	TileMap map[string][2]int `json:"tile_map"`
}

type tileData struct {
	Type     string     `json:"type"`
	Position [2]float64 `json:"position"`
}
