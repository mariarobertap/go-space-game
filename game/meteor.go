package game

import (
	"game/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationSpeedMin = -0.02
	rotationSpeedMax = 0.02
)

type Meteor struct {
	position      Vector
	speed         float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewMeteor() *Meteor {
	pos := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	speed := (rand.Float64() * 13)

	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	m := &Meteor{
		position:      pos,
		speed:         speed,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin),
		sprite:        sprite,
	}
	return m
}

func (m *Meteor) Update() {

	m.position.Y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, op)
}

func (m *Meteor) Collider() Rect {
	bounds := m.sprite.Bounds()

	return NewRect(
		m.position.X,
		m.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
