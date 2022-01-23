package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TileSizeX   = 16
	TileSizeY   = 16
	ScreenSizeX = 1280
	ScreenSizeY = 720
	NumTilesX   = ScreenSizeX / TileSizeX
	NumTilesY   = ScreenSizeY / TileSizeY
)

func main() {
	ebiten.SetWindowSize(ScreenSizeX, ScreenSizeY)
	ebiten.SetWindowTitle("SlayRL")

	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
