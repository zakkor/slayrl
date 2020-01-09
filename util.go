package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Point struct {
	X, Y int
}

func LinePoints(x0, y0, x1, y1 int) []Point {
	var points []Point

	// implemented straight from WP pseudocode
	dx := x1 - x0
	if dx < 0 {
		dx = -dx
	}
	dy := y1 - y0
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {
		points = append(points, Point{X: x0, Y: y0})
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
	return points
}

func CirclePoints(x, y, r int) []Point {
	var points []Point

	if r < 0 {
		return nil
	}
	// Bresenham algorithm
	x1, y1, err := -r, 0, 2-2*r
	for {
		points = append(points,
			Point{X: x - x1, Y: y + y1},
			Point{X: x - y1, Y: y - x1},
			Point{X: x + x1, Y: y - y1},
			Point{X: x + y1, Y: y + x1},
		)
		r = err
		if r > x1 {
			x1++
			err += x1*2 + 1
		}
		if r <= y1 {
			y1++
			err += y1*2 + 1
		}
		if x1 >= 0 {
			break
		}
	}
	return points
}

func CircleThickPoints(x, y, r int) (points []Point) {
	circle := CirclePoints(x, y, r)
	for _, cp := range circle {
		points = append(points, cp)
		points = append(points, CirclePoints(cp.X, cp.Y, 1)...)
	}
	return
}

func RectPoints(x1, y1, x2, y2 int) (points []Point) {
	for x := x1; x <= x2; x++ {
		points = append(points, Point{X: x, Y: y1})
		points = append(points, Point{X: x, Y: y2})
	}
	for y := y1; y <= y2; y++ {
		points = append(points, Point{X: x1, Y: y})
		points = append(points, Point{X: x2, Y: y})
	}
	return
}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

// RandIntRange returns a random int in the range [min, max].
func RandIntRange(min, max int) int {
	return rand.Intn(max+1-min) + min
}
