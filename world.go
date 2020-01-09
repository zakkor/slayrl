package main

import (
	"github.com/hajimehoshi/ebiten"
)

type World struct {
	SizeX, SizeY int
	ambientLight float64
	player       *Player
	entities     []*Entity
	chars        []*Character
}

func NewWorld(sizex, sizey int) World {
	w := World{
		SizeX:        sizex,
		SizeY:        sizey,
		ambientLight: 0.1,
	}

	var entities []*Entity
	for x := 0; x < sizex; x++ {
		for y := 0; y < sizey; y++ {
			entities = append(entities, &Entity{
				X:           x,
				Y:           y,
				Image:       Images["ground"],
				Description: "ground",
				Walkable:    true,
				Visibility:  w.ambientLight,
			})
		}
	}

	w.entities = entities

	player := NewPlayer()
	player.X = 25
	player.Y = 20

	w.player = player

	w.chars = []*Character{
		&Character{
			Actor: Actor{
				Entity: Entity{
					X:             30,
					Y:             20,
					Image:         Images["character"],
					Visibility:    1.0,
					ObstructsView: true,
					Description:   "character",
				},
			},
			Speed: 1.0,
		},
		&Character{
			Actor: Actor{
				Entity: Entity{
					X:             25,
					Y:             25,
					Image:         Images["character"],
					Visibility:    1.0,
					ObstructsView: true,
					Description:   "character",
				},
			},
			Speed: 1.0,
		},
	}

	return w
}

func (w *World) CalculateVisibility() {
	x, y := w.player.X, w.player.Y
	visrange := w.player.VisibilityRange

	// Reset visibility to default value
	for _, e := range w.entities {
		e.Visibility = w.ambientLight
	}
	for _, c := range w.chars {
		c.Visibility = w.ambientLight
	}

	circle := CircleThickPoints(x, y, visrange)
	for _, cp := range circle {
		line := LinePoints(x, y, cp.X, cp.Y)

		var (
			brightness         = 1.0
			falloff    float64 = (brightness - w.ambientLight) / float64(len(line))
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
	for _, c := range w.chars {
		if c.X == x && c.Y == y {
			return &c.Entity
		}
	}
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
