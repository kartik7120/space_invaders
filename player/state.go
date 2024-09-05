package player

import (
	"game/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var gameMap map[string]*GameState = make(map[string]*GameState)
var currKey *string

type GameState struct {
	InvaderAnimation   *utils.Timer
	AnimationPositive  bool
	AnimationNegative  bool
	Player             *Player
	Invaders           [][]*Invader
	Bullets            []*Lazer
	LazerTimer         *utils.Timer
	Score              int
	Audioplayer        *audio.Player
	StateInstanceCount int
	Invaders3          []*Invader
	InvaderTimer       *utils.Timer
	InvaderBullets     []*Lazer
}

func NewGameState() *GameState {
	return &GameState{
		InvaderAnimation:   utils.NewTimer(2 * time.Second),
		AnimationPositive:  true,
		AnimationNegative:  false,
		Player:             NewPlayer(),
		Invaders:           GenerateInvaders(5, 5),
		Bullets:            []*Lazer{},
		LazerTimer:         utils.NewTimer(2 * time.Second),
		Score:              0,
		StateInstanceCount: 0,
		InvaderTimer:       utils.NewTimer(2 * time.Second),
	}
}

func SetGameState(key string, state *GameState) {
	gameMap[key] = state
	currKey = &key
}

func GetGameState(key string) *GameState {
	if state, ok := gameMap[key]; ok {
		return state
	} else {
		return nil
	}
}

func DeleteState(key string) {
	delete(gameMap, key)
	currKey = nil
}

func GetCurrKey() string {
	return *currKey
}
