package main

type Player struct {
	Entity
	VisibilityRange int
}

func NewPlayer() Player {
	return Player{
		Entity: Entity{
			X:          0,
			Y:          0,
			Image:      Images["player"],
			Walkable:   false,
			Visibility: 1.0,
		},
		VisibilityRange: 16,
	}
}

func (p *Player) Move(w *World, offx, offy int) {
	newx, newy := p.X+offx, p.Y+offy
	if w.At(newx, newy).Walkable == false {
		return
	}

	p.X = newx
	p.Y = newy
	w.CalculateVisibility(p.X, p.Y, p.VisibilityRange)
}
