package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	manager "github.com/tducasse/ebiten-manager"
)

type ContextType struct {
	World   *ebiten.Image
	Manager *manager.Manager
}

var Context *ContextType
