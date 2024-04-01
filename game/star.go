package game

import (
	"game/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Star struct {
	position      Vector
	rotation      float64
	movement      Vector
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewStar() *Star {
	pos := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	velocity := float64(6)

	movement := Vector{
		X: 0,
		Y: velocity,
	}

	sprite := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]

	m := &Star{
		position: pos,
		movement: movement,
		sprite:   sprite,
	}
	return m
}

func (m *Star) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (m *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, op)
}
