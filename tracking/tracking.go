package tracking

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func Init() (<-chan keyboard.KeyEvent, error) {
	defer keyboard.Close()
	key, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting keyboard activity tracking...")
	for {
		event := <-keys
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}
