package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	LogSizeX = ScreenSizeX
	LogSizeY = 200
)

type Log struct {
	Image         *ebiten.Image
	Width, Height int
	Lines         []string
}

func NewLog() Log {
	image, _ := ebiten.NewImage(LogSizeX, LogSizeY, ebiten.FilterDefault)
	return Log{
		Image:  image,
		Width:  LogSizeX,
		Height: LogSizeY,
	}
}

func (l *Log) WriteLine(message string) {
	l.Lines = append([]string{message}, l.Lines...)
}

func (l *Log) Draw(screen *ebiten.Image) {
	const (
		paddingTop  = 18
		paddingLeft = 8
		lineHeight  = 28
	)

	l.Image.Fill(color.Black)
	for i, line := range l.Lines {
		text.Draw(l.Image, line, DefaultFontFace,
			paddingLeft, paddingTop+i*lineHeight,
			// color.RGBA{R: 255, G: 255, B: 255, A: uint8(255 - i*50)})
			color.White)
	}

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(ScreenSizeY-l.Height))
	screen.DrawImage(l.Image, &op)
}
