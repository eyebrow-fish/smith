package smith

import (
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)

const (
	SpriteSize int = 10
)

var (
	//go:embed assets/smith.png
	PlayerSprite []byte

	//go:embed assets/heart.png
	HeartSprite []byte

	//go:embed assets/world.png
	WorldSprite []byte

	//go:embed assets/world.json
	WorldRawData []byte
)

type Game struct {
	inputState
	options WindowOptions
	player  player
	hud     hud
	world   world
}

func NewGame(options WindowOptions) *Game {
	player, err := newPlayer()
	if err != nil {
		log.Fatalf("failed to load sprite: %v", err)
	}

	hud, err := newHud()
	if err != nil {
		log.Fatalf("failed to load sprite: %v", err)
	}

	world, err := newWorld()
	if err != nil {
		log.Fatalf("failed to load world sprite map: %v", err)
	}

	return &Game{
		inputState{},
		options,
		*player,
		*hud,
		*world,
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.collectInputs()
	if g.hasReleased(ebiten.KeyF1) {
		g.debugMode = !g.debugMode
	}
	g.player.handle(g.inputState)

	if err := g.player.physics(g.world); err != nil {
		return err
	}

	if err := screen.Fill(color.RGBA{A: 0xFF}); err != nil {
		return err
	}

	if err := g.world.draw(screen); err != nil {
		return err
	}
	if err := g.player.draw(screen); err != nil {
		return err
	}
	if err := g.hud.draw(screen, g.player); err != nil {
		return err
	}

	return g.handle(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / g.options.Scale, outsideHeight / g.options.Scale
}

func (g *Game) handle(screen *ebiten.Image) error {
	if g.debugMode {
		debugText := fmt.Sprintf(
			"fps: %.f\n%v\n%v",
			ebiten.CurrentFPS(),
			g.inputState,
			g.player,
		)
		err := ebitenutil.DebugPrint(screen, debugText)
		if err != nil {
			return err
		}
	}

	if g.player.health <= 0 && g.hasReleased(ebiten.KeySpace) {
		if player, err := newPlayer(); err != nil {
			return err
		} else {
			g.player = *player
		}
	}

	return nil
}

type WindowOptions struct {
	Scale int
}

type vertex2 struct {
	x float64
	y float64
}

func (v vertex2) scale(scale vertex2) vertex2 {
	return vertex2{v.x * scale.x, v.y * scale.y}
}

func (v vertex2) scaleBy(scale float64) vertex2 {
	return vertex2{v.x * scale, v.y * scale}
}
