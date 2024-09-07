package scenes

import (
	"fmt"
	"game/player"
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	manager "github.com/tducasse/ebiten-manager"
)

var InvaderTextSource = utils.MustLoadInvaderFont()

var Lvl3Screen *manager.Scene = &manager.Scene{
	Init: func(setReady func()) {
		// this has to be called to guarantee that we don't call anything before Init
		player.SetGameState("lvl3", player.NewGameState())
		state := player.GetGameState("lvl3")
		state.Score += player.GetGameState("lvl2").Score
		state.Invaders3 = player.GenerateInvadersLvl3(6, 6)

		setReady()
	},
	Update: func(setReady func()) error {
		// this has to be called to guarantee that we don't call Update before Draw
		state := player.GetGameState("lvl3")
		setReady()
		state.InvaderAnimation.Update()
		if state.InvaderAnimation.IsReady() {
			if state.AnimationPositive {
				for i := 0; i < len(state.Invaders3); i++ {
					state.Invaders3[i].Update(2)
				}
				state.AnimationPositive = false
				state.AnimationNegative = true
			} else if state.AnimationNegative {
				for i := 0; i < len(state.Invaders3); i++ {
					state.Invaders3[i].Update(-2)
				}
				state.AnimationPositive = true
				state.AnimationNegative = false
			}
			state.InvaderAnimation.Reset()
		}

		for i := 0; i < len(state.Bullets); i++ {
			state.Bullets[i].Update(state.Player)
		}

		state.InvaderTimer.Update()

		if state.InvaderTimer.IsReady() {
			for i := 0; i < len(state.Invaders3); i++ {
				if state.Invaders3[i].InvaderType == "yellow" {
					yellowInvaderWidth, yellowInvaderHeight := text.Measure("%", &text.GoTextFace{
						Source: InvaderTextSource,
						Size:   15,
					}, float64(2))
					var X, Y = state.Invaders3[i].GetPostion()
					halfWidth := float64(yellowInvaderWidth) / 2
					halfHeight := float64(yellowInvaderHeight) / 2
					lazerXPos := float64(X) + halfWidth
					lazerYPos := float64(Y) + halfHeight
					state.InvaderBullets = append(state.InvaderBullets, player.NewLazer(lazerXPos, lazerYPos))
					state.InvaderTimer.Reset()
				}
			}
		}

		for i := 0; i < len(state.InvaderBullets); i++ {
			state.InvaderBullets[i].InvaderUpdate()
		}

		// Collision detection

		for i := 0; i < len(state.Invaders3); i++ {
			for j := 0; j < len(state.Bullets); j++ {
				if state.Bullets[j].Collider().Intersects(state.Invaders3[i].Collider()) {
					if state.Invaders3[i].InvaderType == "yellow" {
						state.Score += 50
					} else if state.Invaders3[i].InvaderType == "red" {
						state.Score += 20
					} else {
						state.Score += 10
					}
					state.Invaders3 = append(state.Invaders3[:i], state.Invaders3[i+1:]...)
					state.Bullets = append(state.Bullets[:j], state.Bullets[j+1:]...)
				}
			}
		}

		for i := 0; i < len(state.InvaderBullets); i++ {
			if state.InvaderBullets[i].Collider().Intersects(state.Player.Collider()) {
				state.InvaderBullets = append(state.InvaderBullets[:i], state.InvaderBullets[i+1:]...)
				Context.Manager.SwitchTo("gameOver")
			}
		}

		state.Player.Update()

		if len(state.Invaders3) == 0 {
			Context.Manager.SwitchTo("endScreen")
		}

		return nil
	},
	Draw: func(screen *ebiten.Image) {
		// this has to be called to guarantee that we don't call Draw before Update
		state := player.GetGameState("lvl3")
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", state.Score))
		for i := 0; i < len(state.Invaders3); i++ {
			if state.Invaders3[i].InvaderType == "yellow" {
				op := &text.DrawOptions{}
				x, y := state.Invaders3[i].GetPostion()
				op.ColorScale.Scale(255, 255, 0, 255)
				op.GeoM.Translate(float64(x), float64(y))
				text.Draw(screen, "%", &text.GoTextFace{Source: InvaderTextSource, Size: 15}, op)
			} else {
				state.Invaders3[i].Draw(screen)
			}
		}

		for i := 0; i < len(state.Bullets); i++ {
			state.Bullets[i].Draw(screen)
		}

		for i := 0; i < len(state.InvaderBullets); i++ {
			state.InvaderBullets[i].Draw(screen)
		}

		state.Player.Draw(screen)
	},
}
