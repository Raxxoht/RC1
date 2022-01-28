package state

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var west string = "HS rev korn kylling"
var boat string = ""
var east string = ""

func ViewState() string {
	return fmt.Sprintf("[%s---V \\____%s_____/ _____________Ø---%s]", west, boat, east)
}

func PutInBoat(item string) error {

	itemArray := strings.Fields(item)
	itemArrayLength := len(itemArray)

	for c, item := range itemArray {
		bad1 := []string{"rev", "kylling"}
		bad2 := []string{"kylling", "rev"}
		bad3 := []string{"kylling", "korn"}
		bad4 := []string{"korn", "kylling"}

		if strings.Contains(west, item) {
			tmpWest := strings.ReplaceAll(west, item, "")
			tmpWest = strings.TrimSpace(tmpWest)

			tmpWestArray := strings.Fields(tmpWest)

			if c == itemArrayLength-1 {
				if reflect.DeepEqual(tmpWestArray, bad1) || reflect.DeepEqual(tmpWestArray, bad2) {
					return errors.New("Du kan ikke la rev og kylling være alene")
				}

				if reflect.DeepEqual(tmpWestArray, bad3) || reflect.DeepEqual(tmpWestArray, bad4) {
					return errors.New("Du kan ikke la kylling og korn være alene")
				}

				if len(strings.Fields(boat)) > 2 {
					return errors.New("Du kan ikke ta med deg flere enn to ting på båten")
				}
			}

			west = tmpWest
			boat = boat + " " + item
		}

		if strings.Contains(east, item) {
			tmpEast := strings.ReplaceAll(east, item, "")
			tmpEast = strings.TrimSpace(tmpEast)

			tmpEastArray := strings.Fields(tmpEast)

			if c == itemArrayLength-1 {
				if reflect.DeepEqual(tmpEastArray, bad1) || reflect.DeepEqual(tmpEastArray, bad2) {
					return errors.New("Du kan ikke la rev og kylling være alene")
				}

				if reflect.DeepEqual(tmpEastArray, bad3) || reflect.DeepEqual(tmpEastArray, bad4) {
					return errors.New("Du kan ikke la kylling og korn være alene")
				}

				if len(strings.Fields(boat)) > 2 {
					return errors.New("Du kan ikke ta med deg flere enn to ting på båten")
				}
			}

			east = tmpEast
			boat = boat + " " + item
		}
	}
	return nil
}

func Cross(place string) error {
	if place == "west" {
		west = west + boat
		boat = ""
	} else if place == "east" {
		east = east + boat
		boat = ""
	} else {
		return errors.New(fmt.Sprintf("'%s' is an invalid place", place))
	}

	return nil
}
