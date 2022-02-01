package main

import (
	"main.go/state"
)

func main() {
	state.ViewState()
	state.Put("kylling")
	state.Cross()
	state.Takeout("kylling")
	state.ViewState()
}
