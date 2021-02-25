package main

import (
	"math/rand"
	"time"
)

type area struct {
	food       *food
	snake      *snake
	hasfood    func(*area, coord) bool
	height     int
	width      int
	pointsChan chan (int)
}

func newArea(s *snake, p chan (int), h, w int) *area {
	rand.Seed(time.Now().UnixNano())
}
