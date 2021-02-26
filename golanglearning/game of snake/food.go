package game_of_snake

import (
	"os"
	"strings"
)

type food struct {
	emoji        rune
	points, x, y int
}

func newFood(x, y int) *food {
	return &food{
		points: 10,
		emoji:  getFoodEmoji(),
		x:      x,
		y:      y,
	}
}
func getFoodEmoji() rune {
	if hasUnicodeSupport() {
		return FoodEmoji()
	}
	return '@'
}
func FoodEmoji() rune {
	f := []rune{
		'🍪',
	}
	return f[0]
}
func hasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}
