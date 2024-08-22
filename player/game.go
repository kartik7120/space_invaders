package player

import (
	"fmt"
	"game/utils"
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Timer struct {
	currentTicks int
	targetTicks  int
}

func newTimer(d time.Duration) *Timer {
	return &Timer{
		targetTicks:  int(d.Milliseconds()) * ebiten.TPS() / 1000,
		currentTicks: 0,
	}
}

func (t *Timer) update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) reset() {
	t.currentTicks = 0
}

func (t *Timer) isReady() bool {
	return t.currentTicks >= t.targetTicks
}

type Game struct {
	invaderAnimation  *Timer
	animationPositive bool
	animationNegative bool
	player            *Player
	invaders          [][]*Invader
	bullets           []*Lazer
	lazerTimer        *Timer
	score             int
	Audioplayer       *audio.Player
}

func (g *Game) Update() error {
	g.invaderAnimation.update()
	if g.invaderAnimation.isReady() {
		if g.animationPositive {
			for i := 0; i < len(g.invaders); i++ {
				for j := 0; j < len(g.invaders[i]); j++ {
					g.invaders[i][j].Update(2)
				}
			}
			g.animationPositive = false
			g.animationNegative = true
		} else if g.animationNegative {
			for i := 0; i < len(g.invaders); i++ {
				for j := 0; j < len(g.invaders[i]); j++ {
					g.invaders[i][j].Update(-2)
				}
			}
			g.animationPositive = true
			g.animationNegative = false
		}
		g.invaderAnimation.reset()
	}

	for i := 0; i < len(g.bullets); i++ {
		g.bullets[i].Update(g.player)
	}

	// Collision detection

	for i := 0; i < len(g.invaders); i++ {
		for j := 0; j < len(g.invaders[i]); j++ {
			for k := 0; k < len(g.bullets); k++ {
				if g.bullets[k].Collider().Intersects(g.invaders[i][j].Collider()) {
					if g.invaders[i][j].invaderType == "white" {
						g.score += 10
					}
					if g.invaders[i][j].invaderType == "red" {
						g.score += 20
					}
					g.invaders[i] = append(g.invaders[i][:j], g.invaders[i][(j+1):]...)
					g.bullets = append(g.bullets[:k], g.bullets[k+1:]...)
				}
			}
		}
	}

	g.player.Update()
	return nil
}

func (g *Game) AddBullet(lazer *Lazer) {
	g.bullets = append(g.bullets, lazer)
}

func (g *Game) Draw(screen *ebiten.Image) {

	for i := 0; i < len(g.invaders); i++ {
		for j := 0; j < len(g.invaders[i]); j++ {
			g.invaders[i][j].Draw(screen)
		}
	}

	for i := 0; i < len(g.bullets); i++ {
		g.bullets[i].Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))

	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func NewGame() *Game {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Space Invaders")

	g := &Game{
		invaderAnimation:  newTimer(2 * time.Second),
		bullets:           []*Lazer{},
		animationPositive: true,
		animationNegative: false,
		invaders:          GenerateInvaders(5, 5),
		lazerTimer:        newTimer(2 * time.Second),
		score:             0,
		Audioplayer:       utils.Audioplayer("lazerSound.mp3"),
	}

	g.player = NewPlayer(g)

	return g
}
