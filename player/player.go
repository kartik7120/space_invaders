package player

import (
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Vector struct {
	x, y float64
}

type Player struct {
	position Vector
	player   *ebiten.Image
}

var spaceShip *ebiten.Image = utils.MustLoadImage("space_invader_ship.png")
var audioPlayer = utils.PlayLazerSound()

func NewPlayer() *Player {

	bounds := spaceShip.Bounds()
	halfWidth := float64(bounds.Dx()) / 2

	pos := Vector{
		x: float64(screenWidth/2) - halfWidth,
		y: float64(screenHeight) - float64(bounds.Dy()),
	}

	return &Player{
		position: pos,
		player:   spaceShip,
	}
}

func (p *Player) Update() {
	state := GetGameState(*currKey)
	state.LazerTimer.Update()
	screenXboundary := float64(screenWidth) - float64(p.player.Bounds().Dx())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.position.x >= 0 {
			p.position.x -= 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.position.x <= screenXboundary {
			p.position.x += 2
		}
	} else if state.LazerTimer.IsReady() && ebiten.IsKeyPressed(ebiten.KeyControl) {
		audioPlayer.Rewind()
		audioPlayer.Play()
		bounds := spaceShip.Bounds()
		halfWidth := float64(bounds.Dx()) / 2
		halfHeight := float64(bounds.Dy()) / 2
		lazerXPos := p.position.x + halfWidth
		lazerYPos := p.position.y + halfHeight
		t := audioPlayer.Position()
		if t.Seconds() > 1 {
			audioPlayer.Pause()
		}
		state.Bullets = append(state.Bullets, NewLazer(lazerXPos, lazerYPos))
		state.LazerTimer.Reset()
	}
	// audioPlayer.Close()
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.x, p.position.y)
	screen.DrawImage(p.player, op)
}

func (p *Player) Collider() utils.Rect {
	bounds := spaceShip.Bounds()

	return utils.NewRect(p.position.x, p.position.y, float64(bounds.Dx()), float64(bounds.Dy()))
}
