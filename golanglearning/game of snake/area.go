package game_of_snake

import (
	"math/rand"
	"time"
)

type area struct {
	food       *food
	snake      *snake
	hasFood    func(*area, coord) bool
	height     int
	width      int
	pointsChan chan (int)
}

func newArea(s *snake, p chan (int), h, w int) *area {
	rand.Seed(time.Now().UnixNano())
	a := &area{
		snake:      s,
		height:     h,
		width:      w,
		pointsChan: p,
		hasFood:    hasFood,
	}

	a.placeFood()

	return a
}

func (a *area) moveSnake() error {
	if err := a.snake.move(); err != nil {
		return err
	}

	if a.snakeLeftArena() {
		return a.snake.die()
	}

	if a.hasFood(a, a.snake.head()) {
		go a.addPoints(a.food.points)
		a.snake.length++
		a.placeFood()
	}

	return nil
}

func (a *area) snakeLeftArena() bool {
	h := a.snake.head()
	return h.x > a.width || h.y > a.height || h.x < 0 || h.y < 0
}

func (a *area) addPoints(p int) {
	a.pointsChan <- p
}

func (a *area) placeFood() {
	var x, y int

	for {
		x = rand.Intn(a.width)
		y = rand.Intn(a.height)

		if !a.isOccupied(coord{x: x, y: y}) {
			break
		}
	}

	a.food = newFood(x, y)
}

func hasFood(a *area, c coord) bool {
	return c.x == a.food.x && c.y == a.food.y
}

func (a *area) isOccupied(c coord) bool {
	return a.snake.isOnPosition(c)
}
