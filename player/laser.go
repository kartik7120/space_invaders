package player

import (
	"game/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var lazerImage *ebiten.Image = utils.MustLoadImage("lazer.png")

type Lazer struct {
	X, Y float64
}

func NewLazer(x, y float64) *Lazer {
	return &Lazer{
		X: x,
		Y: y,
	}
}

func (l *Lazer) Update(player *Player) {

	if l.Y < 0 {
		lazerImage.Deallocate()
	}
	// l.Coliision(player.game.invaders)
	l.X = player.position.x
	l.Y -= 5
}

func (l *Lazer) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(lazerImage, float32(l.X), float32(l.Y), 2, 5, color.White, true)
}

func (l *Lazer) Coliision(invaders [][]*Invader) {
	for i := 0; i < len(invaders); i++ {
		for j := 0; j < len(invaders[i]); j++ {
			if invaders[i][j].position.x < l.X && l.X < invaders[i][j].position.x+float64(invaders[i][j].invader.Bounds().Dx()) && invaders[i][j].position.y < l.Y && l.Y < invaders[i][j].position.y+float64(invaders[i][j].invader.Bounds().Dy()) {
				invaders[i][j] = nil
				l.Y = -1
			}
		}
	}
}
