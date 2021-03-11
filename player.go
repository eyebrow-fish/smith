package smith

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type Player struct {
	scale         float64
	position      vertex2
	animation     Animation
	speed         float64
	speedModifier float64
	moving        bool
	health        float32
	maxHealth     float32
	healthStatus  healthStatus
	falling       bool
}

func NewPlayer() (*Player, error) {
	spriteMapImage, _, err := image.Decode(bytes.NewReader(PlayerSprite))
	if err != nil {
		return nil, err
	}

	spriteMap, err := ebiten.NewImageFromImage(spriteMapImage, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	animation := Animation{spriteMap: spriteMap, maxFrame: 2, debounce: 5}

	return &Player{
		scale:         2,
		animation:     animation,
		speed:         1,
		speedModifier: 1,
		health:        10,
		maxHealth:     10,
	}, nil
}

func (p *Player) movementSpeed() float64 {
	return p.speed * p.speedModifier
}

func (p *Player) handle(game InputState) {
	movementKeysPressed := game.rawIndex(ebiten.KeyW) &
		game.rawIndex(ebiten.KeyS) &
		game.rawIndex(ebiten.KeyA) &
		game.rawIndex(ebiten.KeyD)

	if movementKeysPressed > -1 {
		var (
			verticalVelocity   int
			horizontalVelocity int
		)

		if game.rawIndex(ebiten.KeyW) > -1 {
			p.animation.direction = up
			p.position.y -= p.movementSpeed()
			verticalVelocity--
		}
		if game.rawIndex(ebiten.KeyS) > -1 {
			p.animation.direction = down
			p.position.y += p.movementSpeed()
			verticalVelocity++
		}
		if game.rawIndex(ebiten.KeyA) > -1 {
			p.animation.direction = down
			p.position.x -= p.movementSpeed()
			horizontalVelocity--
		}
		if game.rawIndex(ebiten.KeyD) > -1 {
			p.animation.direction = up
			p.position.x += p.movementSpeed()
			horizontalVelocity++
		}

		p.moving = (verticalVelocity != 0 || horizontalVelocity != 0) && p.speedModifier > 0
	} else {
		p.moving = false
	}

	if p.falling {
		p.speedModifier = 0
		if p.health > 0 {
			p.health--
		}
		if p.scale > 0 {
			p.scale -= 0.1
		}
	}
}

func (p *Player) draw(screen *ebiten.Image) error {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(p.scale, p.scale)
	options.GeoM.Translate(p.position.x, p.position.y)

	p.animation.paused = !p.moving

	spriteTile, err := p.animation.update()
	if err != nil {
		return err
	}

	return screen.DrawImage(spriteTile, options)
}

func (p *Player) physics(world World) error {
	var inTile bool
	for _, t := range world.tiles {
		playerPosition := p.position
		playerSize := vertex2{float64(SpriteSize), float64(SpriteSize)}.scaleBy(p.scale)
		invertedFallBuffer := 2.5

		tilePosition := t.position.scaleBy(t.scale)
		tileSize := t.scale * float64(SpriteSize)

		boundUpperX := tilePosition.x <= playerPosition.x+invertedFallBuffer &&
			tilePosition.x+tileSize >= playerPosition.x+invertedFallBuffer
		boundUpperY := tilePosition.y <= playerPosition.y+invertedFallBuffer &&
			tilePosition.y+tileSize >= playerPosition.y+invertedFallBuffer

		boundLowerX := tilePosition.x <= playerPosition.x+playerSize.x-invertedFallBuffer &&
			tilePosition.x+tileSize >= playerPosition.x+playerSize.x-invertedFallBuffer
		boundLowerY := tilePosition.y <= playerPosition.y+playerSize.y-invertedFallBuffer &&
			tilePosition.y+tileSize >= playerPosition.y+playerSize.y-invertedFallBuffer

		if (boundUpperX && boundUpperY) || (boundLowerX && boundLowerY) {
			inTile = true
			break
		}
	}

	p.falling = p.falling || !inTile

	return nil
}

func (p Player) String() string {
	return fmt.Sprintf(
		"position: [%.2f, %.2f]\nmoving: %v\nhealth: %.2f/%.2f %s\nanimation:\n%v\nfalling: %v",
		p.position.x,
		p.position.y,
		p.moving,
		p.health,
		p.maxHealth,
		p.healthStatus,
		p.animation,
		p.falling,
	)
}

type healthStatus uint8

const (
	healthy healthStatus = iota
	poisoned
	freezing
)

func (h healthStatus) String() string {
	return [...]string{"Healthy", "Poisoned", "Freezing"}[h]
}
