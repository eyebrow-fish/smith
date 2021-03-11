package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

const (
	SpriteSize int = 10
)

type Game struct {
	InputState
	Options GameOptions
	Player  Player
	Hud     Hud
	World   World
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.collectInputs()
	if g.hasReleased(ebiten.KeyF1) {
		g.debugMode = !g.debugMode
	}
	g.Player.handle(g.InputState)

	if err := g.Player.physics(g.World); err != nil {
		return err
	}

	if err := screen.Fill(color.RGBA{A: 0xFF}); err != nil {
		return err
	}

	if err := g.World.draw(screen); err != nil {
		return err
	}
	if err := g.Player.draw(screen); err != nil {
		return err
	}
	if err := g.Hud.draw(screen, g.Player); err != nil {
		return err
	}

	if g.debugMode {
		debugText := fmt.Sprintf(
			"fps: %.f\n%v\n%v",
			ebiten.CurrentFPS(),
			g.InputState,
			g.Player,
		)
		err := ebitenutil.DebugPrint(screen, debugText)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / g.Options.Scale, outsideHeight / g.Options.Scale
}

type GameOptions struct {
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
