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
	world := NewWorld(100, 100)
	world.Tiles().Point(20, 20).Line(10, 10, 10, 15).Do(func(e *Entity) {
		e.ObstructsView = true
		e.Walkable = false
		e.Image = Images["wall"]
	})

	player := NewPlayer()
	player.X = 15
	player.Y = 10

	world.CalculateVisibility(player.X, player.Y, player.VisibilityRange)

	log := NewLog()
	log.WriteLine("luamiai pula sa mio sugi ca pe vaca cand o mulgi")
	log.WriteLine("hey andreea")

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
	g.LogicalUpdate()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	g.Draw(screen)

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

func (g *Game) LogicalUpdate() {
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.player.Draw(screen)
	g.log.Draw(screen)
}
