package main

import (
	"neko.bar/nyactivity/tracking"
)

func main() {
	keys, err := tracking.Init()
	if err != nil {
		panic(err)
	}

}
