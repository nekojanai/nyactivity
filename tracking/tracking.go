package tracking

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type Tracker struct {
	KeyChannel <-chan keyboard.KeyEvent
}

func Init() *Tracker {
	keys, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting keyboard activity tracking...")
	return &Tracker{
		keys,
	}
}

func Close() {
	err := keyboard.Close()
	if err != nil {
		panic(err)
	}
}
