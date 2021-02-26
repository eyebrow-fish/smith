package smith

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	InputState
	Player Player
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.collectInputs()
	if g.hasReleased(ebiten.KeyF1) {
		g.debugMode = !g.debugMode
	}

	if g.rawIndex(ebiten.KeyW) > -1 {
		g.Player.position.y -= g.Player.speed
	}
	if g.rawIndex(ebiten.KeyS) > -1 {
		g.Player.position.y += g.Player.speed
	}
	if g.rawIndex(ebiten.KeyA) > -1 {
		g.Player.position.x -= g.Player.speed
	}
	if g.rawIndex(ebiten.KeyD) > -1 {
		g.Player.position.x += g.Player.speed
	}
	g.Player.animation.tick()
	if err := g.Player.draw(screen); err != nil {
		return err
	}
	if g.debugMode {
		debugText := fmt.Sprintf("debug\n%v\n%v", &g.InputState, &g.Player)
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

type Vertex2 struct {
	x float64
	y float64
}
