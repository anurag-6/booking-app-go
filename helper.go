package main


import (
	"strings"
)


func validateUserInputs(
    firstName string,
    lastName string,
    email string,
    userTickets uint)(bool,  bool, bool){


    // Basic validation for inputs
    isValidName := len(firstName) > 2 && len(lastName) > 0
    isValidEmail := strings.Contains(email, "@")
    isValidTicketNumber := userTickets > 0 && userTickets <= 50

    return isValidName, isValidEmail,isValidTicketNumber
}