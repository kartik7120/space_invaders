package utils

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

func PlayLazerSound() *audio.Player {
	file, err := os.Open("lazer_sound3.mp3")
	if err != nil {
		log.Fatal(err)
	}
	// defer file.Close()

	s, err := mp3.DecodeWithoutResampling(file)

	if err != nil {
		log.Fatal(err)
	}

	// Create audio player

	p, err := (*audio.NewContext(44100)).NewPlayer(s)
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func PlayInvaderKillSound() *audio.Player {
	// Load MP3 file
	file, err := os.Open("hit_by_a_wood.mp3")
	if err != nil {
		log.Fatal(err)
	}
	// defer file.Close()

	s, err := mp3.DecodeWithoutResampling(file)

	if err != nil {
		log.Fatal(err)
	}

	// Create audio player

	p, err := (*audio.NewContext(44100)).NewPlayer(s)
	if err != nil {
		log.Fatal(err)
	}

	return p
}
