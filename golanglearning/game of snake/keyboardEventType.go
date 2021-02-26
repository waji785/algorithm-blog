package game_of_snake

import "github.com/nsf/termbox-go"

type keyboardEventType int

const (
	MOVE keyboardEventType = 1 + iota
	RETRY
	END
)

type keyboardEvent struct {
	eventType keyboardEventType
	key       termbox.Key
}

func keyToDirection(k termbox.Key) direction {
	switch k {
	case termbox.KeyArrowLeft:
		return LEFT
	case termbox.KeyArrowDown:
		return DOWN
	case termbox.KeyArrowRight:
		return RIGHT
	case termbox.KeyArrowUp:
		return UP
	}
}
func listenToKeyboard(evChan chan keyboardEvent) {
	termbox.SetInputMode(termbox.InputAlt)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowDown:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowRight:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowUp:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyEsc:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			default:
				if ev.Ch == 'r' {
					evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
