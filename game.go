package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{
	InputState
	player Player
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.collectInputs()
	if g.hasReleased(ebiten.KeyF1) {
		g.debugMode = !g.debugMode
	}
	if g.debugMode {
		debugText := fmt.Sprintf("debug\n%v\n%v", &g.InputState, &g.player)
		err := ebitenutil.DebugPrint(screen, debugText)
		if err != nil {
			return err
		}
	}
	if g.rawIndex(ebiten.KeyW) > -1 {
		g.player.position.y += 0.1
	}
	if g.rawIndex(ebiten.KeyS) > -1 {
		g.player.position.y -= 0.1
	}
	if g.rawIndex(ebiten.KeyA) > -1 {
		g.player.position.x -= 0.1
	}
	if g.rawIndex(ebiten.KeyD) > -1 {
		g.player.position.x += 0.1
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type Vertex2 struct {
	x float64
	y float64
}
