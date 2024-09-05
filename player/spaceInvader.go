package player

import (
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Invader struct {
	position    Vector
	invader     *ebiten.Image
	InvaderType string
}

var invaderImage *ebiten.Image = utils.MustLoadImage("space_invader6.png")
var redInvaderImage *ebiten.Image = utils.MustLoadImage("red_invader7.png")
var InvaderTextSource = utils.MustLoadInvaderFont()

func NewInvader(invaderType string) *Invader {

	pos := Vector{
		x: 0,
		y: 10,
	}
	if invaderType == "red" {
		return &Invader{
			position:    pos,
			invader:     redInvaderImage,
			InvaderType: invaderType,
		}
	}
	if invaderType == "yellow" {
		return &Invader{
			position:    pos,
			invader:     nil,
			InvaderType: invaderType,
		}
	}
	return &Invader{
		position:    pos,
		invader:     invaderImage,
		InvaderType: invaderType,
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

func (i *Invader) GetPostion() (x, y int) {
	x = int(i.position.x)
	y = int(i.position.y)
	return
}

func GenerateInvaders(rows, columns int) [][]*Invader {
	screenXMid := screenWidth/2 - invaderImage.Bounds().Dx()/2
	invaderMatrix := make([][]*Invader, rows)

	for i := 0; i < rows; i++ {
		invaderMatrix[i] = make([]*Invader, columns)
		for j := 0; j < columns; j++ {
			invader := NewInvader("white")
			invader.position.x = float64(screenXMid + j*invaderImage.Bounds().Dx() - (columns-1)*invaderImage.Bounds().Dx()/2 + j*5)
			invader.position.y = float64(i*invaderImage.Bounds().Dy() + ((i + 1) * invaderImage.Bounds().Dy() / 2))
			invaderMatrix[i][j] = invader
		}
	}
	return invaderMatrix
}

func GenerateInvadersLvl2(rows, columns int) [][]*Invader {
	screenXMid := screenWidth/2 - invaderImage.Bounds().Dx()/2
	invaderMatrix := make([][]*Invader, rows)
	screenXRedMid := screenWidth/2 - redInvaderImage.Bounds().Dx()/2

	for i := 0; i < rows; i++ {
		invaderMatrix[i] = make([]*Invader, columns)
		for j := 0; j < columns; j++ {
			if i%2 != 0 {
				invader := NewInvader("red")
				invader.position.x = float64(screenXRedMid + j*redInvaderImage.Bounds().Dx() - (columns-1)*redInvaderImage.Bounds().Dx()/2 + j*5)
				invader.position.y = float64(i*redInvaderImage.Bounds().Dy() + ((i + 1) * redInvaderImage.Bounds().Dy() / 2))
				invaderMatrix[i][j] = invader
				continue
			}
			invader := NewInvader("white")
			invader.position.x = float64(screenXMid + j*invaderImage.Bounds().Dx() - (columns-1)*invaderImage.Bounds().Dx()/2 + j*5)
			invader.position.y = float64(i*invaderImage.Bounds().Dy() + ((i + 1) * invaderImage.Bounds().Dy() / 2))
			invaderMatrix[i][j] = invader
		}
	}
	return invaderMatrix
}

func GenerateInvadersLvl3(rows, columns int) []*Invader {
	yellowInvaderWidth, yellowInvaderHeight := text.Measure("%", &text.GoTextFace{
		Source: InvaderTextSource,
		Size:   15,
	}, float64(2))

	yellowXMid := screenWidth/2 - yellowInvaderWidth/2
	yellowYMid := screenHeight/2 - yellowInvaderHeight/2
	whiteXMid := screenWidth/2 - invaderImage.Bounds().Dx()/2
	whiteYMid := screenHeight/2 - invaderImage.Bounds().Dy()/2
	// redXMid := screenWidth/2 - redInvaderImage.Bounds().Dx()/2

	invaderMatrix := make([]*Invader, 0)
	// 1st side of cross
	for i := 0; i < columns; i++ {
		if i < columns-1 {
			invader := NewInvader("white")
			invader.position.x = float64(whiteXMid - i*invaderImage.Bounds().Dx() - invaderImage.Bounds().Dx())
			invader.position.y = float64(whiteYMid - i*invaderImage.Bounds().Dy() - invaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader)
		}
		if i == columns-1 {
			invader := NewInvader("yellow")
			invader.position.x = float64(yellowXMid)
			invader.position.y = float64(yellowYMid)
			invaderMatrix = append(invaderMatrix, invader)
			invader2 := NewInvader("red")
			invader2.position.x = float64(whiteXMid - (i-1)*invaderImage.Bounds().Dx() - invaderImage.Bounds().Dx() - redInvaderImage.Bounds().Dx())
			invader2.position.y = float64(whiteYMid - (i-1)*invaderImage.Bounds().Dy() - redInvaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader2)
			invader3 := NewInvader("yellow")
			invader3.position.x = float64(whiteXMid - (i-1)*invaderImage.Bounds().Dx() - invaderImage.Bounds().Dx() + redInvaderImage.Bounds().Dx())
			invader3.position.y = float64(whiteYMid - (i-1)*invaderImage.Bounds().Dy() - redInvaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader3)
		}
	}

	// 2nd side of cross

	for i := 0; i < columns; i++ {
		if i < columns-1 {
			invader := NewInvader("white")
			invader.position.x = float64(whiteXMid + i*invaderImage.Bounds().Dx() + invaderImage.Bounds().Dx())
			invader.position.y = float64(whiteYMid - i*invaderImage.Bounds().Dy() - invaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader)
		}
		if i == columns-1 {
			invader2 := NewInvader("red")
			invader2.position.x = float64(whiteXMid + (i-1)*invaderImage.Bounds().Dx() + invaderImage.Bounds().Dx() + redInvaderImage.Bounds().Dx())
			invader2.position.y = float64(whiteYMid - (i-1)*invaderImage.Bounds().Dy() - redInvaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader2)
			invader3 := NewInvader("yellow")
			invader3.position.x = float64(whiteXMid + (i-1)*invaderImage.Bounds().Dx() + invaderImage.Bounds().Dx() - redInvaderImage.Bounds().Dx())
			invader3.position.y = float64(whiteYMid - (i-1)*invaderImage.Bounds().Dy() - redInvaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader3)
		}
	}

	// 3rd side of cross

	for i := 0; i < columns-2; i++ {
		if i < columns-1-2 {
			invader := NewInvader("white")
			invader.position.x = float64(whiteXMid - i*invaderImage.Bounds().Dx() - invaderImage.Bounds().Dx())
			invader.position.y = float64(whiteYMid + i*invaderImage.Bounds().Dy() + invaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader)
		}
		if i == columns-1-2 {
			invader2 := NewInvader("red")
			invader2.position.x = float64(whiteXMid - (i-1)*invaderImage.Bounds().Dx() - invaderImage.Bounds().Dx() - redInvaderImage.Bounds().Dx())
			invader2.position.y = float64(whiteYMid + (i-1)*invaderImage.Bounds().Dy() + redInvaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader2)
		}
	}

	// 4th side of cross

	for i := 0; i < columns-2; i++ {
		if i < columns-1-2 {
			invader := NewInvader("white")
			invader.position.x = float64(whiteXMid + i*invaderImage.Bounds().Dx() + invaderImage.Bounds().Dx())
			invader.position.y = float64(whiteYMid + i*invaderImage.Bounds().Dy() + invaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader)
		}
		if i == columns-1-2 {
			invader2 := NewInvader("red")
			invader2.position.x = float64(whiteXMid + (i-1)*invaderImage.Bounds().Dx() + invaderImage.Bounds().Dx() + redInvaderImage.Bounds().Dx())
			invader2.position.y = float64(whiteYMid + (i-1)*invaderImage.Bounds().Dy() + redInvaderImage.Bounds().Dy())
			invaderMatrix = append(invaderMatrix, invader2)
		}
	}

	return invaderMatrix
}

func (i *Invader) Collider() utils.Rect {
	if i.InvaderType == "red" {
		bounds := redInvaderImage.Bounds()
		return utils.NewRect(i.position.x, i.position.y, float64(bounds.Dx()), float64(bounds.Dy()))
	}

	if i.InvaderType == "yellow" {
		yellowInvaderWidth, yellowInvaderHeight := text.Measure("%", &text.GoTextFace{
			Source: InvaderTextSource,
			Size:   15,
		}, float64(2))

		return utils.NewRect(i.position.x, i.position.y, float64(yellowInvaderWidth), float64(yellowInvaderHeight))
	}

	bounds := invaderImage.Bounds()

	return utils.NewRect(i.position.x, i.position.y, float64(bounds.Dx()), float64(bounds.Dy()))
}
