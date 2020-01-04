package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	TileSizeX   = 16
	TileSizeY   = 16
	ScreenScale = 1
	ScreenSizeX = 1280 / ScreenScale
	ScreenSizeY = 720 / ScreenScale
)

var (
	NumTilesX = ScreenSizeX / TileSizeX
	NumTilesY = ScreenSizeY / TileSizeY

	PlayerImage *ebiten.Image
	GroundImage *ebiten.Image
	WallImage   *ebiten.Image
)

func main() {
	g := NewGame()
	ebiten.SetWindowSize(ScreenSizeX, ScreenSizeY)
	ebiten.SetWindowTitle("SlayRL")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func loadTileset(path string) (*ebiten.Image, error) {
	f, err := ebitenutil.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func getTileFromTileset(tileset *ebiten.Image, ix, iy int) *ebiten.Image {
	x := ix * TileSizeX
	y := iy * TileSizeY
	return tileset.SubImage(image.Rect(x, y, x+TileSizeX, y+TileSizeY)).(*ebiten.Image)
}
