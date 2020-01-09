package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

type Game struct {
	states []State
	world  World
	// player Player
	looker Actor
	// chars  []*Character
	log Log
}

type State int

const (
	StateDefault State = 0
	StateLooking State = iota
)

func NewGame() *Game {
	world := NewWorld(100, 60)
	world.Tiles().Point(20, 20).Line(10, 10, 10, 15).Do(func(e *Entity) {
		e.ObstructsView = true
		e.Walkable = false
		e.Image = Images["wall"]
		e.Description = "wall"
	})

	world.CalculateVisibility()

	log := NewLog()
	log.WriteLine("Welcome to SlayRL")

	looker := Actor{
		Entity: Entity{
			X:          0,
			Y:          0,
			Image:      Images["looker"],
			Visibility: 1.0,
		},
	}

	return &Game{
		states: []State{StateDefault},
		world:  world,
		looker: looker,
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
	g.Draw(screen)

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) ProcessInput() {
	switch g.TopState() {
	case StateDefault:
		// Movement
		if repeatingKeyPressed(ebiten.KeyUp) {
			g.world.player.Move(&g.world, DirectionUp)
			g.LogicalUpdate(1 / g.world.player.Speed)
		} else if repeatingKeyPressed(ebiten.KeyDown) {
			g.world.player.Move(&g.world, DirectionDown)
			g.LogicalUpdate(1 / g.world.player.Speed)
		} else if repeatingKeyPressed(ebiten.KeyLeft) {
			g.world.player.Move(&g.world, DirectionLeft)
			g.LogicalUpdate(1 / g.world.player.Speed)
		} else if repeatingKeyPressed(ebiten.KeyRight) {
			g.world.player.Move(&g.world, DirectionRight)
			g.LogicalUpdate(1 / g.world.player.Speed)
		}
		// Wait
		if repeatingKeyPressed(ebiten.KeyPeriod) {
			g.LogicalUpdate(1)
		}
		// Look
		if repeatingKeyPressed(ebiten.KeyK) {
			g.PushState(StateLooking)
			g.looker.X = g.world.player.X
			g.looker.Y = g.world.player.Y
		}
	case StateLooking:
		// Move the looker on the map, ignoring collision.
		if repeatingKeyPressed(ebiten.KeyUp) {
			g.looker.Move(&g.world, DirectionUp, false)
		} else if repeatingKeyPressed(ebiten.KeyDown) {
			g.looker.Move(&g.world, DirectionDown, false)
		} else if repeatingKeyPressed(ebiten.KeyLeft) {
			g.looker.Move(&g.world, DirectionLeft, false)
		} else if repeatingKeyPressed(ebiten.KeyRight) {
			g.looker.Move(&g.world, DirectionRight, false)
		}
		cx, cy := ebiten.CursorPosition()
		x, y := cx/TileSizeX, cy/TileSizeY
		if x >= 0 && y >= 0 {
			g.looker.X = x
			g.looker.Y = y
		}
	}

	// Escape
	if g.TopState() != StateDefault && repeatingKeyPressed(ebiten.KeyEscape) {
		g.PopState()
	}
}

func (g *Game) LogicalUpdate(d float64) {
	for _, c := range g.world.chars {
		c.Update(d, &g.world)
	}
	g.world.CalculateVisibility()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	for _, c := range g.world.chars {
		c.Draw(screen)
	}
	g.world.player.Draw(screen)

	switch g.TopState() {
	case StateDefault:
		g.log.Draw(screen)
	case StateLooking:
		g.looker.Draw(screen)
		desc := g.world.At(g.looker.X, g.looker.Y).Description
		if desc == "" {
			desc = "unknown"
		}
		text.Draw(screen, desc, DefaultFontFace,
			g.looker.X*TileSizeX, g.looker.Y*TileSizeY-8,
			color.White)
	}
}

func (g *Game) TopState() State {
	return g.states[len(g.states)-1]
}

func (g *Game) PushState(s State) {
	g.states = append(g.states, s)
}

func (g *Game) PopState() {
	g.states = g.states[:len(g.states)-1]
}
