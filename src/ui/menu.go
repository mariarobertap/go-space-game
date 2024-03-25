package ui

import (
	"game/src/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Menu struct {
	readyToPlay bool
}

const (
	screenWidth  = 800
	screenHeight = 600
)

func NewMenu() *Menu {

	return &Menu{
		readyToPlay: false,
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {

	text.Draw(screen, "GO: Space War", assets.ScoreFont, 200, 300, color.White)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(315, 150)
	screen.DrawImage(assets.GopherPlayer, op)

	text.Draw(screen, "Press ENTER to start..", assets.FontUi, 100, 400, color.White)

}

func (m *Menu) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		m.readyToPlay = true
	}

}

func (m *Menu) IsReady() bool {
	return m.readyToPlay
}
func (m *Menu) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
