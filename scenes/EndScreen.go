package scenes

import (
	"game/player"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	manager "github.com/tducasse/ebiten-manager"
)

var EndScreen *manager.Scene = &manager.Scene{
	Init: func(setReady func()) {
		// this has to be called to guarantee that we don't call anything before Init
		setReady()
	},
	Update: func(setReady func()) error {
		// this has to be called to guarantee that we don't call Update before Draw
		setReady()
		if ebiten.IsKeyPressed(ebiten.KeyControl) {
			Context.Manager.SwitchTo("title")
		}
		return nil
	},
	Draw: func(screen *ebiten.Image) {

		op := &text.DrawOptions{}
		textWidth, textHeight := text.Measure("Press Ctrl to restart the game", &text.GoTextFace{
			Source: englishTextSource,
			Size:   15,
		}, float64(2))
		op.GeoM.Translate(float64(320/2)-float64(textWidth/2), float64(240/2)-float64(textHeight))
		text.Draw(screen, "Press Ctrl to restart the game", &text.GoTextFace{
			Source: englishTextSource,
			Size:   15,
		}, op)

		titleWidth, titleHeight := text.Measure("Thank you for playing!", &text.GoTextFace{
			Source: spaceInvaderTextSource,
			Size:   20,
		}, float64(2))

		op2 := &text.DrawOptions{}
		op2.GeoM.Translate(float64(320/2)-float64(titleWidth/2), float64(240/2)-float64(titleHeight)-float64(textHeight))
		text.Draw(screen, "Thank you for playing!", &text.GoTextFace{
			Source: spaceInvaderTextSource,
			Size:   20,
		}, op2)

		scoreWidth, scoreHeight := text.Measure("Your score was: "+strconv.Itoa(player.GetGameState("lvl3").Score), &text.GoTextFace{
			Source: englishTextSource,
			Size:   15,
		}, float64(2))

		op3 := &text.DrawOptions{}
		op3.GeoM.Translate(float64(320/2)-float64(scoreWidth/2), float64(240/2)-float64(scoreHeight)-float64(titleHeight)-float64(textHeight))
		text.Draw(screen, "Your score was: "+strconv.Itoa(player.GetGameState("lvl3").Score), &text.GoTextFace{
			Source: englishTextSource,
			Size:   15,
		}, op3)

	},
}
