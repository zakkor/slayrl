package main

import (
	"github.com/hajimehoshi/ebiten"
)

type World struct {
	SizeX, SizeY   int
	tiles          [][]Tile
	visibility     [][]float64
	zeroVisibility float64
}

type Tile struct {
	X, Y          int
	Image         *ebiten.Image
	Walkable      bool
	ObstructsView bool
}

func NewWorld(sizex, sizey int) World {
	tiles := make([][]Tile, sizex)
	for x := 0; x < sizex; x++ {
		tiles[x] = make([]Tile, sizey)
		for y := 0; y < sizey; y++ {
			tiles[x][y] = Tile{X: x, Y: y, Image: GroundImage, Walkable: true}
		}
	}

	visibility := make([][]float64, sizex)
	for x := 0; x < sizex; x++ {
		visibility[x] = make([]float64, sizey)
	}

	return World{
		SizeX:          sizex,
		SizeY:          sizey,
		tiles:          tiles,
		visibility:     visibility,
		zeroVisibility: 0.1,
	}
}

func (w *World) Tiles() *builder {
	return newBuilder(w)
}

func (w *World) ClearTile(t *Tile, x, y int) {
	*t = Tile{X: x, Y: y, Image: GroundImage, Walkable: true}
}

func (w *World) CalculateVisibility(x, y, visrange int) {
	for x := 0; x < w.SizeX; x++ {
		for y := 0; y < w.SizeY; y++ {
			w.visibility[x][y] = w.zeroVisibility
		}
	}

	circle := CircleThickPoints(x, y, visrange)
	for _, cp := range circle {
		line := LinePoints(x, y, cp.X, cp.Y)

		var (
			brightness         = 1.0
			falloff    float64 = (brightness - w.zeroVisibility) / float64(len(line))
		)
		for _, lp := range line {
			if lp.X < 0 || lp.Y < 0 {
				continue
			}
			w.visibility[lp.X][lp.Y] = brightness
			brightness -= falloff

			// We draw the thing that is obstructing the view, and next iteration stop drawing
			if w.tiles[lp.X][lp.Y].ObstructsView {
				break
			}
		}
	}
}

func (w *World) Draw(screen *ebiten.Image) {
	w.Tiles().All().Do(func(t *Tile, x, y int) {
		t.Draw(screen, w.visibility[x][y])
	})
}

func (t *Tile) Draw(screen *ebiten.Image, visibility float64) {
	op := ebiten.DrawImageOptions{}
	// Set position
	op.GeoM.Translate(float64(t.X*TileSizeX), float64(t.Y*TileSizeY))
	// Set visibility
	op.ColorM.Scale(1.0, 1.0, 1.0, visibility)
	screen.DrawImage(t.Image, &op)
}
