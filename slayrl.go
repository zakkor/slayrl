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

	playerImage *ebiten.Image
	groundImage *ebiten.Image
	wallImage   *ebiten.Image
)

func main() {
	tileset, err := loadTileset("./resources/tileset_zilk.png")
	if err != nil {
		panic(err)
	}
	playerImage = getTileFromTileset(tileset, 0, 4)
	groundImage = getTileFromTileset(tileset, 14, 2)
	wallImage = getTileFromTileset(tileset, 0, 11)

	world := NewWorld(100, 100)

	// explosion := world.Tiles()
	// world.Tiles().Rect(0, 0, 40, 40).Do(func(t *Tile, x, y int) {
	// 	if x%5 == 0 && y%5 == 0 {
	// 		explosion.Line(20, 20, x, y)
	// 	}
	// })
	// for i := 0; i < 40; i += 3 {
	// 	explosion.Except().Line(0, i, 40, i)
	// }
	// for i := 0; i < 40; i += 3 {
	// 	explosion.Except().Line(i, 0, i, 40)
	// }

	// // Set explosion image
	// explosion.Do(func(t *Tile, x, y int) {
	// 	t.Image = wallImage
	// 	t.Walkable = false
	// })

	// // Build a rectangular wall, with a hole as an entrance.
	// world.Tiles().RectFill(5, 5, 20, 20).Do(world.ClearTile)
	// world.Tiles().Rect(5, 5, 20, 20).Except().Point(9, 5).Do(func(t *Tile, x, y int) {
	// 	t.Image = wallImage
	// 	t.Walkable = false
	// })

	player := NewPlayer(playerImage)
	player.X = 10
	player.Y = 4

	world.Tiles().Rect(10, 10, 20, 20).Rect(22, 10, 40, 20).Do(func(t *Tile, x, y int) {
		t.Image = wallImage
		t.Walkable = false
		t.ObstructsView = true
	})

	// world.Tiles().Circle(10, 10, 5).Line(15, 15, 20, 21).Do(func(t *Tile, x, y int) {
	// 	t.Image = wallImage
	// 	t.Walkable = false
	// 	t.ObstructsView = true
	// })

	world.CalculateVisibility(player.X, player.Y, player.VisibilityRange)

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
