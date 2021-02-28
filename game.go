package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
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
	if err := screen.Fill(color.RGBA{A: 0xFF}); err != nil {
		return err
	}
	if err := g.World.draw(screen); err != nil {
		return err
	}
	g.Player.handle(g.InputState)
	if err := g.Player.draw(screen); err != nil {
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
	if err := g.Hud.draw(screen, g.Player); err != nil {
		return err
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
