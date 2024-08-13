package utils

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	currentTicks int
	targetTicks  int
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		targetTicks:  int(d.Milliseconds()) * ebiten.TPS() / 1000,
		currentTicks: 0,
	}
}

func (t *Timer) Update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= t.targetTicks
}
