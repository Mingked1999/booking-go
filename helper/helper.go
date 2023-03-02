package helper

import (
	"strings" // pkg.go.dev/fmt
)

func ValidInputs(firstname string, lastname string, email string, userTickets uint, remianingTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumber := userTickets > 0 && userTickets <= remianingTickets
	return isValidName, isValidEmail, isValidNumber
}
