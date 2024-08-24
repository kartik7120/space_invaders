package main

import (
	"game/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	manager "github.com/tducasse/ebiten-manager"
)

type Game struct {
}

var m *manager.Manager = manager.MakeManager(map[string]*manager.Scene{
	"title": scenes.TitleScreen,
	"lvl1":  scenes.Lvl1Screen,
},
	// the first scene
	"title")

func (game *Game) Update() error {
	return m.Update()
}

func (game *Game) Draw(screen *ebiten.Image) {
	m.Draw(screen)
}

func (game *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Space Invaders")

	scenes.Context = &scenes.ContextType{
		Manager: m,
		World:   ebiten.NewImage(320, 240),
	}
	g := &Game{}
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
