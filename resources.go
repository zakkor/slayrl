package main

import (
	"image"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

var (
	Images          map[string]*ebiten.Image
	DefaultFontFace font.Face
)

func init() {
	// Images
	tileset, err := loadTileset("./resources/tileset_zilk.png")
	if err != nil {
		panic(err)
	}
	Images = map[string]*ebiten.Image{
		"player": getTileFromTileset(tileset, 0, 4),
		"ground": getTileFromTileset(tileset, 14, 2),
		"wall":   getTileFromTileset(tileset, 0, 11),
	}

	// Fonts
	fontData, err := ioutil.ReadFile("./resources/slkscr.ttf")
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
