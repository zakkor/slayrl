package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
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
	tileset, err := loadTileset("./resources/tileset_zilk.png")
	if err != nil {
		panic(err)
	}
	PlayerImage = getTileFromTileset(tileset, 0, 4)
	GroundImage = getTileFromTileset(tileset, 14, 2)
	WallImage = getTileFromTileset(tileset, 0, 11)

	world := NewWorld(100, 100)

	player := NewPlayer(PlayerImage)
	player.X = 10
	player.Y = 10

	world.CalculateVisibility(player.X, player.Y, player.VisibilityRange)

	log := NewLog(450, 250)
	log.WriteLine("cf andreea")

	update := func(screen *ebiten.Image) error {
		if repeatingKeyPressed(ebiten.KeyUp) {
			player.Move(world, 0, -1)
		} else if repeatingKeyPressed(ebiten.KeyDown) {
			player.Move(world, 0, 1)
		} else if repeatingKeyPressed(ebiten.KeyLeft) {
			player.Move(world, -1, 0)
		} else if repeatingKeyPressed(ebiten.KeyRight) {
			player.Move(world, 1, 0)
		}

		if ebiten.IsDrawingSkipped() {
			return nil
		}

		world.Draw(screen)
		player.Draw(screen)

		log.Draw(screen)

		return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	}

	if err := ebiten.Run(update, ScreenSizeX, ScreenSizeY, ScreenScale, "SlayRL"); err != nil {
		panic(err)
	}
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
