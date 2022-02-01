package state

// Jeg importerer fmt for å printe ut forskjellige variabler og setninger
// Jeg importerer errors for å gi feilmeldinger
import (
	"errors"
	"fmt"
)

//ervest må linkes opp med øst og vest, gjenstander fra øst kan ikke i båt som ligger i vest
var ervest bool = true
var vest = []string{"mann", "rev", "kylling", "korn"}
var øst []string
var båt []string
var index int

func ViewState() {
	if ervest == true {
		fmt.Printf("[%q---v/__%q__/_____________ø---%q]", vest, båt, øst)
	} else if ervest == false {
		fmt.Printf("[%q---v__________/__%q__/ø---%q]", vest, båt, øst)
	}
}

func Put(item string) {
	if len(båt) > 2 {
		errors.New("Du kan ikke ha flere enn to ting i båten")
	} else if Contains(vest, item) && ervest == true {
		vest = append(vest[:index], vest[index+1:]...)
		båt = append(båt, item)
	} else if Contains(øst, item) && ervest == false {
		øst = append(øst[:index], øst[index+1:]...)
		båt = append(båt, item)
	}
}

func Contains(l []string, x string) bool {
	index = 0
	for _, i := range l {
		if i == x {
			return true
		}
		index += 1
	}
	return false
}

func Takeout(item string) {
	if Contains(båt, item) {
		båt = append(båt[:index], båt[index+1:]...)
	}
	switch ervest {
	case true:
		vest = append(vest, item)
	case false:
		øst = append(øst, item)
	}
}

func Cross() {
	if Contains(båt, "mann") == false {
		errors.New("Du kan ikke krysse uten at menneske er i båten")
	} else if ervest == true {
		ervest = false
	} else if ervest == false {
		ervest = true
	}
}
