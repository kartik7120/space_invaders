package utils

import (
	"bytes"
	_ "embed"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed OpenSans-Regular.ttf
var englishTTF []byte

var EnglighFileSource *text.GoTextFaceSource

//go:embed ElectronPulse.ttf
var spaceInvaderTTF []byte

//go:embed Invaders.ttf
var InvaderTTF []byte

func MustLoadImage(name string) *ebiten.Image {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func MustLoadFont() *text.GoTextFaceSource {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(englishTTF))

	if err != nil {
		panic(err)
	}

	return s
}

func MustLoadSpaceInvaderFont() *text.GoTextFaceSource {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(spaceInvaderTTF))

	if err != nil {
		panic(err)
	}

	return s
}

func MustLoadInvaderFont() *text.GoTextFaceSource {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(InvaderTTF))

	if err != nil {
		panic(err)
	}

	return s
}
