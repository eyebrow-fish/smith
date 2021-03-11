package smith

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type Player struct {
	position     vertex2
	animation    Animation
	speed        float64
	moving       bool
	health       float32
	maxHealth    float32
	healthStatus healthStatus
	falling      bool
}

func NewPlayer(sprite []byte) (*Player, error) {
	spriteMapImage, _, err := image.Decode(bytes.NewReader(sprite))
	if err != nil {
		return nil, err
	}

	spriteMap, err := ebiten.NewImageFromImage(spriteMapImage, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	animation := Animation{spriteMap: spriteMap, maxFrame: 2, debounce: 5}

	return &Player{animation: animation, speed: 1, health: 10, maxHealth: 10}, nil
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
			p.position.y -= p.speed
			verticalVelocity--
		}
		if game.rawIndex(ebiten.KeyS) > -1 {
			p.animation.direction = down
			p.position.y += p.speed
			verticalVelocity++
		}
		if game.rawIndex(ebiten.KeyA) > -1 {
			p.animation.direction = down
			p.position.x -= p.speed
			horizontalVelocity--
		}
		if game.rawIndex(ebiten.KeyD) > -1 {
			p.animation.direction = up
			p.position.x += p.speed
			horizontalVelocity++
		}

		p.moving = verticalVelocity != 0 || horizontalVelocity != 0
	} else {
		p.moving = false
	}
}

func (p *Player) draw(screen *ebiten.Image) error {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.position.x, p.position.y)
	options.GeoM.Scale(2, 2)

	if p.moving {
		p.animation.paused = false
	} else {
		p.animation.paused = true
	}

	spriteTile, err := p.animation.update()
	if err != nil {
		return err
	}

	return screen.DrawImage(spriteTile, options)
}

func (p *Player) physics(world World) error {
	var inTile bool
	for _, t := range world.tiles {
		playerPosition := p.position.scaleBy(2)
		playerSize := vertex2{10, 10}.scaleBy(2)

		tilePosition := t.position.scale(t.scale)
		tileSize := t.scale.scaleBy(10)

		boundUpperX := tilePosition.x <= playerPosition.x &&
			tilePosition.x + tileSize.x >= playerPosition.x
		boundUpperY := tilePosition.y <= playerPosition.y &&
			tilePosition.y + tileSize.y >= playerPosition.y

		boundLowerX := tilePosition.x <= playerPosition.x + playerSize.x &&
			tilePosition.x + tileSize.x >= playerPosition.x + playerSize.x
		boundLowerY := tilePosition.y <= playerPosition.y + playerSize.y &&
			tilePosition.y + tileSize.y >= playerPosition.y + playerSize.y

		if (boundUpperX && boundUpperY) || (boundLowerX && boundLowerY) {
			inTile = true
			break
		}
	}

	p.falling = !inTile

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
