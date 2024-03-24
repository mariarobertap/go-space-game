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
	screenWidth     = 800
	screenHeight    = 600
	meteorSpawnTime = 1 * time.Second
	starSpawnTime   = (1 * time.Second) / 2
	planetSpawnTime = 7 * time.Second
)

type Game struct {
	meteorSpawnTimer *Timer
	planetSpawnTime  *Timer
	starSpawnTimer   *Timer
	menu             *ui.Menu

	player  *Player
	meteors []*Meteor
	stars   []*Star
	planets []*Planet
	bullets []*Bullet

	isStarted bool
	score     int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		starSpawnTimer:   NewTimer(starSpawnTime),
		planetSpawnTime:  NewTimer(planetSpawnTime),
	}

	g.player = NewPlayer(g)
	g.menu = ui.NewMenu()

	return g
}

func (g *Game) Update() error {

	g.starSpawnTimer.Update()
	if g.starSpawnTimer.IsReady() {
		g.starSpawnTimer.Reset()

		s := NewStar()
		g.stars = append(g.stars, s)
	}

	for _, m := range g.stars {
		m.Update()
	}

	if !g.isStarted {

		g.menu.Update()

		if g.menu.IsReady() {
			g.planets = nil
			g.isStarted = true
		}

		g.planetSpawnTime.Update()
		if g.planetSpawnTime.IsReady() {
			g.planetSpawnTime.Reset()

			s := NewPlanet()
			g.planets = append(g.planets, s)
		}

		for _, m := range g.planets {
			m.Update()
		}

		return nil

	}

	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, b := range g.bullets {
		b.Update()
	}

	// Check for meteor/bullet collisions
	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				//remove meteor by index
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				//remove bullet by index
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

	for _, b := range g.stars {
		b.Draw(screen)
	}

	if !g.isStarted {
		for _, b := range g.planets {
			b.Draw(screen)
		}

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

	if g.score >= bestScore {
		bestScore = g.score

	}

	g.score = 0
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
