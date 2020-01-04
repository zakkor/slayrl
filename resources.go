package main

import "github.com/hajimehoshi/ebiten"

// type ResourceManager struct {
// 	Images ImageManager
// }

type ImageManager struct {
	Ground *ebiten.Image
	Player *ebiten.Image
	Wall   *ebiten.Image
}
