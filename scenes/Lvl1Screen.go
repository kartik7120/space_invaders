package scenes

import (
	"fmt"
	"game/player"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	manager "github.com/tducasse/ebiten-manager"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Lvl1Screen *manager.Scene = &manager.Scene{
	Init: func(setReady func()) {
		// this has to be called to guarantee that we don't call anything before Init
		state := player.NewGameState()
		player.SetGameState("state", state)
		setReady()
	},
	Update: func(setReady func()) error {
		// this has to be called to guarantee that we don't call Update before Draw
		state := player.GetGameState("state")
		setReady()
		state.InvaderAnimation.Update()
		if state.InvaderAnimation.IsReady() {
			if state.AnimationPositive {
				for i := 0; i < len(state.Invaders); i++ {
					for j := 0; j < len(state.Invaders[i]); j++ {
						state.Invaders[i][j].Update(2)
					}
				}
				state.AnimationPositive = false
				state.AnimationNegative = true
			} else if state.AnimationNegative {
				for i := 0; i < len(state.Invaders); i++ {
					for j := 0; j < len(state.Invaders[i]); j++ {
						state.Invaders[i][j].Update(-2)
					}
				}
				state.AnimationPositive = true
				state.AnimationNegative = false
			}
			state.InvaderAnimation.Reset()
		}

		for i := 0; i < len(state.Bullets); i++ {
			state.Bullets[i].Update(state.Player)
		}

		// Collision detection

		for i := 0; i < len(state.Invaders); i++ {
			for j := 0; j < len(state.Invaders[i]); j++ {
				for k := 0; k < len(state.Bullets); k++ {
					if state.Bullets[k].Collider().Intersects(state.Invaders[i][j].Collider()) {
						if state.Invaders[i][j].InvaderType == "white" {
							state.Score += 10
						}
						if state.Invaders[i][j].InvaderType == "red" {
							state.Score += 20
						}
						state.Invaders[i] = append(state.Invaders[i][:j], state.Invaders[i][(j+1):]...)
						state.Bullets = append(state.Bullets[:k], state.Bullets[k+1:]...)
					}
				}
			}
		}

		state.Player.Update()

		return nil
	},
	Draw: func(screen *ebiten.Image) {
		state := player.GetGameState("state")
		for i := 0; i < len(state.Invaders); i++ {
			for j := 0; j < len(state.Invaders[i]); j++ {
				state.Invaders[i][j].Draw(screen)
			}
		}

		for i := 0; i < len(state.Bullets); i++ {
			state.Bullets[i].Draw(screen)
		}

		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", state.Score))

		state.Player.Draw(screen)

	},
}
