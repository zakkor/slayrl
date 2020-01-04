package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	world  World
	player Player
	log    Log
}

func NewGame() *Game {
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
	log.WriteLine("Hello World")

	return &Game{
		world:  world,
		player: player,
		log:    log,
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenSizeX, ScreenSizeY
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.ProcessInput()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.world.Draw(screen)
	g.player.Draw(screen)
	g.log.Draw(screen)

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) ProcessInput() {
	if repeatingKeyPressed(ebiten.KeyUp) {
		g.player.Move(&g.world, 0, -1)
	} else if repeatingKeyPressed(ebiten.KeyDown) {
		g.player.Move(&g.world, 0, 1)
	} else if repeatingKeyPressed(ebiten.KeyLeft) {
		g.player.Move(&g.world, -1, 0)
	} else if repeatingKeyPressed(ebiten.KeyRight) {
		g.player.Move(&g.world, 1, 0)
	}
}
