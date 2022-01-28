package main

import (
	"fmt"
)

func main() {
	fmt.Println(ViewState())
}

func ViewState() string {
	// Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
	return "[kylling rev korn hs ---\\ \\__/ _________________/---]"
}

func Put(item string) string {
	return "[kylling rev korn ---\\ \\__/ _________________/---]"
}
