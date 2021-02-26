package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{
	*InputState
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.handleInputs()
	if g.hasReleased(ebiten.KeyF1) {
		g.debugMode = !g.debugMode
	}
	if g.debugMode {
		debugText := fmt.Sprintf("debug\n%v", g.InputState)
		err := ebitenutil.DebugPrint(screen, debugText)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
