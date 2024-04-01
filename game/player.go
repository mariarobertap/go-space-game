package game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"game/assets"
)

const (
	shootCooldown     = time.Millisecond * 200
	bulletSpawnOffset = 50.0
)

type Player struct {
	game          *Game
	position      Vector
	rotation      float64
	sprite        *ebiten.Image
	shootCooldown *Timer
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2

	pos := Vector{
		X: (screenWidth / 2) - halfW,
		Y: (screenHeight) - 170,
	}

	fmt.Println(shootCooldown)

	return &Player{
		game:          game,
		position:      pos,
		sprite:        sprite,
		shootCooldown: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed

	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed

	}

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.shootCooldown.Reset()

		bounds := p.sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfW,
			p.position.Y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLaser(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Collider() Rect {
	bounds := p.sprite.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
