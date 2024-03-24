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

func NewStar(baseVelocity float64) *Star {
	pos := Vector{
		X: rand.Float64() * screenWidth, // Random X position within the window width
		Y: -100,                         // Set Y position above the window
	}

	// Set the velocity of the meteors to move downwards.
	velocity := baseVelocity + rand.Float64()*2

	movement := Vector{
		X: 0,        // No horizontal movement
		Y: velocity, // Move downwards
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
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}
