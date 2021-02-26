package smith

import "fmt"

type Player struct {
	position Vertex2
}

func (p *Player) String() string {
	return fmt.Sprintf("position: [%.2f, %.2f]\n", p.position.x, p.position.y)
}
