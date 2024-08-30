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
		return nil
	},
	Draw: func(screen *ebiten.Image) {
		// this has to be called to guarantee that we don't call Draw before Update

		textWidth, textHeight := text.Measure("Game Over", &text.GoTextFace{
			Source: SpaceInvaderTextSource,
			Size:   20,
		}, float64(2))

		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(320/2)-float64(textWidth/2), float64(240/2)-float64(textHeight))
		text.Draw(screen, "Game Over", &text.GoTextFace{Source: SpaceInvaderTextSource, Size: 20}, op)
	},
}
