package game

import (
	"fmt"
	"game/assets"
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
	starSpawnTimer   *Timer

	player  *Player
	meteors []*Meteor
	stars   []*Star
	lasers  []*Laser

	isStarted bool
	score     int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
		starSpawnTimer:   NewTimer(10),
	}

	g.player = NewPlayer(g)

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

	for _, b := range g.lasers {
		b.Update()
	}

	for i, m := range g.meteors {
		for j, b := range g.lasers {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score++

			}
		}
	}

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

	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, b := range g.lasers {
		b.Draw(screen)
	}

	scoreSprite := &ebiten.DrawImageOptions{}

	scoreSprite.GeoM.Translate(60, 450)

	text.Draw(screen, "________________________________", assets.FontUi, 0, 520, color.White)
	text.Draw(screen, fmt.Sprintf("Points: %d            High Score: %d", g.score, bestScore), assets.FontUi, 20, 570, color.White)

}

func (g *Game) AddLaser(l *Laser) {
	g.lasers = append(g.lasers, l)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
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
