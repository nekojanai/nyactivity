package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"neko.bar/nyactivity/tracking"
)

func main() {
	defer tracking.Close()
	tracker := tracking.Init()
	keys := tracker.KeyChannel
	for {
		event := <-keys
		if event.Err != nil {
			panic(event.Err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}
