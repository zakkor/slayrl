package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
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

	const lineHeight = 25
	for i, line := range l.Lines {
		text.Draw(l.Image, line, DefaultFontFace, 0, (i+1)*lineHeight, color.White)
	}

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(ScreenSizeY-l.Height))
	screen.DrawImage(l.Image, &op)
}
