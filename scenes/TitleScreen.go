package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	manager "github.com/tducasse/ebiten-manager"
)

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
		ebitenutil.DebugPrint(screen, "Press Ctrl to start the game")
	},
}
