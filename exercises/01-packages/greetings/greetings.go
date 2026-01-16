package greetings

import "strings"

func capitalize(name string) string {
	if len(name) == 0 {
		return name

	}
	return strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
}

func Formal(name string) string {
	properName := capitalize(name)
	return "Good day, " + properName + ". Welcome to our establishment."
}

func Casual(name string) string {
	properName := capitalize(name)
	return "Hey " + properName + "! What's up?"
}
