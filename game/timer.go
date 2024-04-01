package game

import (
	"fmt"
)

type Timer struct {
	currentTicks int
	targetTicks  int
}

func NewTimer(targetTicks int) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks:  targetTicks,
	}
}

func (t *Timer) GetTargetTicks() {
	fmt.Println(t.targetTicks)
}

func (t *Timer) Update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}
