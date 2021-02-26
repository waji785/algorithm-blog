package game_of_snake

import (
	"errors"
)

const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

type direction int
type snake struct {
	body      []coord
	direction direction
	length    int
}

func newSnake(d direction, b []coord) *snake {
	return &snake{
		length:    len(b),
		body:      b,
		direction: d,
	}
}
func (s *snake) changeDirecion(d direction) {
	opposities := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}
	if o := opposities[d]; o != 0 && o != s.direction {
		s.direction = d
	}
}
func (s *snake) die() error {
	return errors.New("die")
}
func (s *snake) head() coord {
	return s.body[len(s.body)-1]
}
func (s *snake) move() error {
	h := s.head()
	c := coord{x: h.x, y: h.y}

	switch s.direction {
	case RIGHT:
		c.x++
	case LEFT:
		c.x--
	case UP:
		c.y++
	case DOWN:
		c.y--
	}
	if s.isOnPosition(c) {
		return s.die()
	}
	if s.length > len(s.body) {
		s.body = append(s.body, c)
	} else {
		s.body = append(s.body[1:], c)
	}
	return nil
}
func (s *snake) isOnPosition(c coord) bool {
	for _, b := range s.body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}