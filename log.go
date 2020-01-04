package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Log struct {
	Image         *ebiten.Image
	Width, Height int
	Lines         []string
}

func NewLog(width, height int) Log {
	image, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	return Log{
		Image:  image,
		Width:  width,
		Height: height,
	}
}

func (l *Log) WriteLine(message string) {
	l.Lines = append([]string{message}, l.Lines...)
}

func (l *Log) Draw(screen *ebiten.Image) {
	l.Image.Fill(color.Black)

	// for _, line := range l.Lines {
	// 	ebiten.Util
	// 	l.Image.draw
	// }

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(ScreenSizeY-l.Height))
	// op.GeoM.Translate(0, 50)
	screen.DrawImage(l.Image, &op)
}
