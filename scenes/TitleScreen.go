package scenes

import (
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	manager "github.com/tducasse/ebiten-manager"
)

var englishTextSource = utils.MustLoadFont()
var spaceInvaderTextSource = utils.MustLoadSpaceInvaderFont()

var TitleScreen *manager.Scene = &manager.Scene{
	Init: func(setReady func()) {
		// this has to be called to guarantee that we don't call anything before Init
		setReady()
	},
	Update: func(setReady func()) error {
		// this has to be called to guarantee that we don't call Update before Draw
		setReady()
		if ebiten.IsKeyPressed(ebiten.KeyControl) {
			Context.Manager.SwitchTo("lvl1")
		}
		return nil
	},
	Draw: func(screen *ebiten.Image) {
		op := &text.DrawOptions{}
		textWidth, textHeight := text.Measure("Press Ctrl to start the game", &text.GoTextFace{
			Source: englishTextSource,
			Size:   15,
		}, float64(2))
		op.GeoM.Translate(float64(320/2)-float64(textWidth/2), float64(240/2)-float64(textHeight))
		text.Draw(screen, "Press Ctrl to start the game", &text.GoTextFace{
			Source: englishTextSource,
			Size:   15,
		}, op)

		titleWidth, titleHeight := text.Measure("Space Invaders", &text.GoTextFace{
			Source: spaceInvaderTextSource,
			Size:   30,
		}, float64(2))

		op2 := &text.DrawOptions{}
		op2.GeoM.Translate(float64(320/2)-float64(titleWidth/2), float64(240/2)-float64(titleHeight)-float64(textHeight))
		text.Draw(screen, "Space Invaders", &text.GoTextFace{
			Source: spaceInvaderTextSource,
			Size:   30,
		}, op2)
	},
}
