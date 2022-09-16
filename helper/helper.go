package helper

import "Strings"

func UserInputValidation (firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isVaidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isVaidName, isValidEmail, isValidTicketNumber	
}