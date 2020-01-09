package main

type Actor struct {
	Entity
}

func (a *Actor) Move(w *World, dir Direction, collision bool) {
	dx, dy := dir.Coords()
	newx, newy := a.X+dx, a.Y+dy
	if newx < 0 || newy < 0 {
		return
	}

	if collision && w.At(newx, newy).Walkable == false {
		return
	}

	a.X = newx
	a.Y = newy
}
