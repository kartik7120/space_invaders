package scenes

import (
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	manager "github.com/tducasse/ebiten-manager"
)

var SpaceInvaderTextSource = utils.MustLoadSpaceInvaderFont()

var GameOverScreen *manager.Scene = &manager.Scene{
	Init: func(setReady func()) {
		// this has to be called to guarantee that we don't call anything before Init
		setReady()
	},
	Update: func(setReady func()) error {
		// this has to be called to guarantee that we don't call Update before Draw
		setReady()
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			Context.Manager.SwitchTo("title")
		}
		return nil
	},
	Draw: func(screen *ebiten.Image) {
		// this has to be called to guarantee that we don't call Draw before Update

		textWidth, textHeight := text.Measure("Game Over", &text.GoTextFace{
			Source: SpaceInvaderTextSource,
			Size:   20,
		}, float64(2))

		textWidth2, textHeight2 := text.Measure("Press Enter to go back to the title screen", &text.GoTextFace{
			Source: SpaceInvaderTextSource,
			Size:   10,
		}, float64(2))

		op := &text.DrawOptions{}
		op2 := &text.DrawOptions{}

		op.GeoM.Translate(float64(320/2)-float64(textWidth/2), float64(240/2)-float64(textHeight))
		op2.GeoM.Translate(float64(320/2)-float64(textWidth2/2), float64(240/2)-float64(textHeight2)+float64(20))
		text.Draw(screen, "Game Over", &text.GoTextFace{Source: SpaceInvaderTextSource, Size: 20}, op)
		text.Draw(screen, "Press Enter to go back to the title screen", &text.GoTextFace{Source: SpaceInvaderTextSource, Size: 10}, op2)
	},
}
