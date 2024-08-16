package player

import (
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	lazerSpeed = 350
)

var lazerImage *ebiten.Image = utils.MustLoadImage("lazer.png")

type Lazer struct {
	X, Y float64
}

func NewLazer(x, y float64) *Lazer {

	bounds := lazerImage.Bounds()
	halfWidth := float64(bounds.Dx()) / 2
	halfHeight := float64(bounds.Dy()) / 2

	x = x - halfWidth
	y = y - halfHeight

	return &Lazer{
		X: x,
		Y: y,
	}
}

func (l *Lazer) Update(player *Player) {

	speed := lazerSpeed / ebiten.TPS()

	l.Y -= float64(speed)
}

func (l *Lazer) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(l.X, l.Y)
	screen.DrawImage(lazerImage, op)
}

func (l *Lazer) Collider() utils.Rect {
	bounds := lazerImage.Bounds()

	return utils.NewRect(l.X, l.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
