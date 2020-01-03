package main

import "github.com/hajimehoshi/ebiten"

type Player struct {
	X, Y            int
	Image           *ebiten.Image
	VisibilityRange int
}

func NewPlayer(image *ebiten.Image) *Player {
	return &Player{
		X:               0,
		Y:               0,
		Image:           image,
		VisibilityRange: 32,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X*TileSizeX), float64(p.Y*TileSizeY))
	op.CompositeMode = ebiten.CompositeModeCopy
	screen.DrawImage(p.Image, &op)
}

func (p *Player) Move(w *World, offx, offy int) {
	newx, newy := p.X+offx, p.Y+offy
	if w.tiles[newx][newy].Walkable == false {
		return
	}

	p.X = newx
	p.Y = newy

	w.CalculateVisibility(p.X, p.Y, p.VisibilityRange)
}
