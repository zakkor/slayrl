package main

type Player struct {
	Actor
	Speed           float64
	VisibilityRange int
}

func NewPlayer() *Player {
	return &Player{
		Actor: Actor{
			Entity: Entity{
				X:          0,
				Y:          0,
				Image:      Images["player"],
				Walkable:   false,
				Visibility: 1.0,
			},
		},
		Speed:           1.0,
		VisibilityRange: 16,
	}
}

func (p *Player) Move(w *World, dir Direction) {
	p.Actor.Move(w, dir, true)
	// w.CalculateVisibility(p.X, p.Y, p.VisibilityRange)
}
