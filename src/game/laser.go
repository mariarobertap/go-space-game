package game

import (
	"game/src/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	game          *Game
	position      Vector
	speed         float64
	rotation      float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewLaser(g *Game, pos Vector) *Laser {
	sprite := assets.LaserSprite
	speed := 7.0

	if g.superPowerActive {
		sprite = assets.SuperPowerSprite
		speed = 12.0
	}

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos.X -= halfW
	pos.Y -= halfH

	b := &Laser{
		game:          g,
		position:      pos,
		speed:         speed,
		rotationSpeed: rotationSpeedMax * 2,
		sprite:        sprite,
	}

	return b
}

func (l *Laser) Update() {
	l.position.Y += -l.speed
	l.rotation += l.rotationSpeed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if l.game.superPowerActive && l.sprite == assets.SuperPowerSprite {
		bounds := assets.SuperPowerSprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		op.GeoM.Translate(-halfW, -halfH)
		op.GeoM.Rotate(l.rotation)
		op.GeoM.Translate(halfW, halfH)
	}

	op.GeoM.Translate(l.position.X, l.position.Y)

	screen.DrawImage(l.sprite, op)
}

func (l *Laser) Collider() Rect {
	bounds := l.sprite.Bounds()

	return NewRect(
		l.position.X,
		l.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
