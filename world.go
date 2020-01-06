package main

import (
	"github.com/hajimehoshi/ebiten"
)

type World struct {
	SizeX, SizeY int
	entities     []*Entity
}

func NewWorld(sizex, sizey int) World {
	var entities []*Entity
	for x := 0; x < sizex; x++ {
		for y := 0; y < sizey; y++ {
			entities = append(entities, &Entity{
				X:          x,
				Y:          y,
				Image:      Images["ground"],
				Walkable:   true,
				Visibility: EntityZeroVisibility,
			})
		}
	}

	return World{
		SizeX:    sizex,
		SizeY:    sizey,
		entities: entities,
	}
}

func (w *World) CalculateVisibility(x, y, visrange int) {
	// Reset visibility to default value
	for _, e := range w.entities {
		e.Visibility = EntityZeroVisibility
	}

	circle := CircleThickPoints(x, y, visrange)
	for _, cp := range circle {
		line := LinePoints(x, y, cp.X, cp.Y)

		var (
			brightness         = 1.0
			falloff    float64 = (brightness - EntityZeroVisibility) / float64(len(line))
		)
		for _, lp := range line {
			if lp.X < 0 || lp.Y < 0 {
				continue
			}
			ent := w.At(lp.X, lp.Y)
			ent.Visibility = brightness
			brightness -= falloff

			// We draw the thing that is obstructing the view, and next iteration stop drawing
			if ent.ObstructsView {
				break
			}
		}
	}
}

func (w *World) At(x, y int) *Entity {
	for _, ent := range w.entities {
		if ent.X == x && ent.Y == y {
			return ent
		}
	}
	panic("entity not found at position")
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, ent := range w.entities {
		ent.Draw(screen)
	}
}

func (w *World) Tiles() *builder {
	return newBuilder(w)
}
