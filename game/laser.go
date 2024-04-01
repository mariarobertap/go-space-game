package game

import (
	"game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	position Vector
	sprite   *ebiten.Image
}

func NewLaser(pos Vector) *Laser {
	sprite := assets.LaserSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos.X -= halfW
	pos.Y -= halfH

	l := &Laser{
		position: pos,
		sprite:   sprite,
	}

	return l
}

func (b *Laser) Update() {
	speed := 7.0

	b.position.Y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
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
