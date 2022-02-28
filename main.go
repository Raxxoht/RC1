package main

import (
	"main.go/state"
)

func main() {
	state.Put("kylling")
	state.Put("mann")
	state.Cross()
	state.Takeout("kylling")
	state.ViewState()
}
