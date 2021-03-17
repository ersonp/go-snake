package pkg

import (
	"strings"

	"github.com/eiannone/keyboard"
)

type keyboardEventType int

// Keyboard events
const (
	MOVE keyboardEventType = 1 + iota
	RETRY
	END
)

// KeyboardEvent type
type keyboardEvent struct {
	key       string
	eventType keyboardEventType
}

func keyToDirection(k string) direction {
	switch k {
	case "a":
		return LEFT
	case "s":
		return DOWN
	case "d":
		return RIGHT
	case "w":
		return UP
	default:
		return 0
	}
}

func listenToKeyboard(evChan chan keyboardEvent) {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		k := strings.ToLower(string(event.Rune))
		if k == "r" {
			evChan <- keyboardEvent{eventType: RETRY, key: k}
		} else if k == "w" || k == "a" || k == "s" || k == "d" {
			evChan <- keyboardEvent{eventType: MOVE, key: k}
		} else if event.Key == keyboard.KeyEsc {
			evChan <- keyboardEvent{eventType: END, key: k}
		}
	}
}
