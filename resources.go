package main

import (
	"image"
	_ "image/png"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

var (
	Images          map[string]*ebiten.Image
	DefaultFontFace font.Face
)

func init() {
	// Images
	tileset, _, err := ebitenutil.NewImageFromFile("./tileset_zilk.png")
	if err != nil {
		panic(err)
	}
	Images = map[string]*ebiten.Image{
		"player":    getTileFromTileset(tileset, 0, 4),
		"looker":    getTileFromTileset(tileset, 4, 0),
		"character": getTileFromTileset(tileset, 2, 0),
		"ground":    getTileFromTileset(tileset, 14, 2),
		"wall":      getTileFromTileset(tileset, 0, 11),
	}

	// Fonts
	fontData, err := ioutil.ReadFile("./slkscr.ttf")
	if err != nil {
		panic(err)
	}
	defaultFont, err := freetype.ParseFont(fontData)
	if err != nil {
		panic(err)
	}
	DefaultFontFace = truetype.NewFace(defaultFont, &truetype.Options{
		Size: 24,
		// DPI:     72,
		Hinting: font.HintingFull,
	})
}

func getTileFromTileset(tileset *ebiten.Image, ix, iy int) *ebiten.Image {
	x := ix * TileSizeX
	y := iy * TileSizeY
	return tileset.SubImage(image.Rect(x, y, x+TileSizeX, y+TileSizeY)).(*ebiten.Image)
}
