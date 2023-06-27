package pets

import (
	"fmt"
	"strings"
)

type Dog struct {
	Name  string
	Color string
	Breed string
}

func (d Dog) Feed(food string) string {
	return fmt.Sprintf("%s is eating %s", d.Name, food)
}

func (d Dog) GiveAttention(activity string) string {
	response := ""
	switch strings.ToUpper(activity) {
	case "PET":
		response = "wags tail"
	case "Playing Fetch":
		response = "return the ball and jump waiting for you to throw it again"
	default:
		response = "barks"
	}
	return fmt.Sprintf("%s loves attention, %s will cause him to %s", d.Name, activity, response)
}
