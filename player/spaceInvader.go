package player

import (
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Invader struct {
	position Vector
	invader  *ebiten.Image
}

var invaderImage *ebiten.Image = utils.MustLoadImage("space_invader6.png")

func NewInvader() *Invader {

	pos := Vector{
		x: 0,
		y: 10,
	}

	return &Invader{
		position: pos,
		invader:  invaderImage,
	}
}

func (i *Invader) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(i.position.x, i.position.y)
	screen.DrawImage(i.invader, opts)
}

func (i *Invader) DrawInvaderMatrix(screen *ebiten.Image, rows int, columns int) {
	invaders := GenerateInvaders(rows, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			invaders[i][j].Draw(screen)
		}
	}
}

func (i *Invader) Update(offset int) {

	i.position.x += float64(offset)

}

func GenerateInvaders(rows, columns int) [][]*Invader {
	invaders := make([]*Invader, 0)
	screenXMid := screenWidth/2 - invaderImage.Bounds().Dx()/2
	invaderMatrix := make([][]*Invader, rows)

	for i := 0; i < rows; i++ {
		invaderMatrix[i] = make([]*Invader, columns)
		for j := 0; j < columns; j++ {
			invader := NewInvader()
			invader.position.x = float64(screenXMid + j*invaderImage.Bounds().Dx())
			invader.position.y = float64(i * invaderImage.Bounds().Dy())
			invaderMatrix[i][j] = invader
			invaders = append(invaders, invader)
		}
	}
	return invaderMatrix
}
