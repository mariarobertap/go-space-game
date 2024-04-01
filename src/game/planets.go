package game

import (
	"game/src/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Planet struct {
	position      Vector
	rotation      float64
	movement      Vector
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewPlanet() *Planet {
	pos := Vector{
		X: rand.Float64() * screenWidth,
		Y: -500,
	}

	velocity := float64(2)

	movement := Vector{
		X: 0,
		Y: velocity,
	}

	sprite := assets.PlanetsSprites[rand.Intn(len(assets.PlanetsSprites))]

	m := &Planet{
		position: pos,
		movement: movement,
		sprite:   sprite,
	}
	return m
}

func (m *Planet) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (m *Planet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, op)
}
