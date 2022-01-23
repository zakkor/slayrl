package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	Image         *ebiten.Image
	X, Y          int
	Walkable      bool
	ObstructsView bool
	Visibility    float64
	Description   string
}

func (e *Entity) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	// Set position
	op.GeoM.Translate(float64(e.X*TileSizeX), float64(e.Y*TileSizeY))
	op.CompositeMode = ebiten.CompositeModeCopy
	// Set visibility
	// if e.Visibility != EntityZeroVisibility {
	// 	subtract := rand.Intn(2)
	// 	if subtract == 1 {
	// 		vis -= 0.02 + rand.Float64()*(0.05-0.02)
	// 	} else {
	// 		vis += 0.1 + rand.Float64()*(0.3-0.1)
	// 	}
	// }
	op.ColorM.Scale(0.961, 0.69, 0.016, e.Visibility)
	screen.DrawImage(e.Image, &op)
}
