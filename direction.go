package main

type Direction int

const (
	DirectionUp   Direction = 0
	DirectionDown Direction = iota
	DirectionLeft
	DirectionRight
	DirectionUpLeft
	DirectionUpRight
	DirectionDownLeft
	DirectionDownRight
)

func DirectionFromCoords(dx, dy int) Direction {
	if dx == 0 && dy == -1 {
		return DirectionUp
	} else if dx == 0 && dy == 1 {
		return DirectionDown
	} else if dx == -1 && dy == 0 {
		return DirectionLeft
	} else if dx == 1 && dy == 0 {
		return DirectionRight
	} else if dx == -1 && dy == -1 {
		return DirectionUpLeft
	} else if dx == 1 && dy == -1 {
		return DirectionUpRight
	} else if dx == -1 && dy == 1 {
		return DirectionDownLeft
	} else if dx == 1 && dy == 1 {
		return DirectionDownLeft
	}
	panic("bad coords for direction")
}

func (d Direction) Coords() (dx, dy int) {
	switch d {
	case DirectionUp:
		dy--
	case DirectionDown:
		dy++
	case DirectionLeft:
		dx--
	case DirectionRight:
		dx++
	case DirectionUpLeft:
		dy--
		dx--
	case DirectionUpRight:
		dy--
		dx++
	case DirectionDownLeft:
		dy++
		dx--
	case DirectionDownRight:
		dy++
		dx++
	default:
		panic("bad direction encountered")
	}
	return
}

func randomDirection() Direction {
	return Direction(RandIntRange(0, 7))
}
