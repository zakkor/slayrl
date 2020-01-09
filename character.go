package main

import "math"

type Character struct {
	Actor
	Speed        float64
	moveLeftover float64
}

func (c *Character) Update(d float64, w *World) {
	c.MoveRandom(d, w)
}

func (c *Character) Move(d float64, w *World, dir Direction) {
	distance := c.Speed*d + c.moveLeftover
	c.moveLeftover = math.Remainder(distance, 1.0)
	move := int(math.Round(distance))

	dx, dy := dir.Coords()
	dx *= move
	dy *= move
	if dx == 0 && dy == 0 {
		return
	}
	newDir := DirectionFromCoords(dx, dy)
	c.Actor.Move(w, newDir, true)
	// w.CalculateVisibility(w.player.X, w.player.Y, w.player.VisibilityRange)
}

func (c *Character) MoveRandom(d float64, w *World) {
	c.Move(d, w, randomDirection())
}
