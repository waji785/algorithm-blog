package game_of_snake

import (
	"time"

	"github.com/nsf/termbox-go"
)

var (
	pointsChan         = make(chan int)
	keyboardEventsChan = make(chan keyboardEvent)
)

type Game struct {
	area   *area
	score  int
	isOver bool
}

func initialSnake() *snake {
	return newSnake(RIGHT, []coord{
		coord{x: 1, y: 1},
		coord{x: 1, y: 2},
		coord{x: 1, y: 3},
		coord{x: 1, y: 4},
	})
}

func initialScore() int {
	return 0
}

func initialArena() *area {
	return newArea(initialSnake(), pointsChan, 20, 50)
}

func (g *Game) end() {
	g.isOver = true
}

func (g *Game) moveInterval() time.Duration {
	ms := 100 - (g.score / 10)
	return time.Duration(ms) * time.Millisecond
}

func (g *Game) retry() {
	g.area = initialArena()
	g.score = initialScore()
	g.isOver = false
}

func (g *Game) addPoints(p int) {
	g.score += p
}

func NewGame() *Game {
	return &Game{area: initialArena(), score: initialScore()}
}

func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyboard(keyboardEventsChan)

	if err := g.render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case p := <-pointsChan:
			g.addPoints(p)
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				g.area.snake.changeDirecion(d)
			case RETRY:
				g.retry()
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if err := g.area.moveSnake(); err != nil {
					g.end()
				}
			}

			if err := g.render(); err != nil {
				panic(err)
			}

			time.Sleep(g.moveInterval())
		}
	}
}
