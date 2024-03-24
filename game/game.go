package game

import (
	"fmt"
	"game/assets"
	"game/ui"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var bestScore = 0

const (
	screenWidth  = 800
	screenHeight = 600

	meteorSpawnTime = 1 * time.Second
	starSpawnTime   = 2 * time.Second

	baseMeteorVelocity  = 0.25
	meteorSpeedUpAmount = 0.1
	meteorSpeedUpTime   = 5 * time.Second
)

type Game struct {
	player           *Player
	meteorSpawnTimer *Timer
	starSpawnTimer   *Timer
	menu             *ui.Menu
	meteors          []*Meteor
	stars            []*Star
	bullets          []*Bullet
	isStarted        bool
	score            int

	baseVelocity float64
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		starSpawnTimer:   NewTimer(starSpawnTime),
		baseVelocity:     baseMeteorVelocity}

	g.player = NewPlayer(g)
	g.menu = ui.NewMenu()

	return g
}

func (g *Game) Update() error {

	if !g.isStarted {

		g.menu.Update()

		if g.menu.IsReady() {
			g.isStarted = true
		}

	}

	g.player.Update()

	g.starSpawnTimer.Update()
	if g.starSpawnTimer.IsReady() {
		g.starSpawnTimer.Reset()

		s := NewStar(0.1)
		g.stars = append(g.stars, s)
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(g.baseVelocity)
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.stars {
		m.Update()
	}

	for _, b := range g.bullets {
		b.Update()
	}

	// Check for meteor/bullet collisions
	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++

			}
		}
	}

	// Check for meteor/player collisions
	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
			break
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//g.background.Draw(screen)

	for _, b := range g.stars {
		b.Draw(screen)
	}

	if !g.isStarted {
		g.menu.Draw(screen)
		return
	}

	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, b := range g.bullets {
		b.Draw(screen)
	}

	scoreSprite := &ebiten.DrawImageOptions{}

	//scoreSprite.GeoM.Scale(-1, -1)
	scoreSprite.GeoM.Translate(60, 450)

	text.Draw(screen, "________________________________", assets.FontUi, 0, 520, color.White)
	text.Draw(screen, fmt.Sprintf("Points: %d            High Score: %d", g.score, bestScore), assets.FontUi, 20, 570, color.White)

}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.bullets = nil
	g.meteorSpawnTimer.Reset()
	g.starSpawnTimer.Reset()

	g.baseVelocity = baseMeteorVelocity

	if g.score >= bestScore {
		bestScore = g.score

	}

	g.score = 0
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
