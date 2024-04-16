package game

import (
	"game/src/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type PowerUp struct {
	position Vector
	movement Vector
	sprite   *ebiten.Image
}

func NewPowerUp() *PowerUp {
	pos := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	velocity := float64(6)

	movement := Vector{
		X: 0,
		Y: velocity,
	}

	sprite := assets.PowerUpSprite

	return &PowerUp{
		position: pos,
		movement: movement,
		sprite:   sprite,
	}
}

func (p *PowerUp) Update() {
	p.position.X += p.movement.X
	p.position.Y += p.movement.Y
}

func (p *PowerUp) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func (p *PowerUp) Collider() Rect {
	bounds := p.sprite.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
