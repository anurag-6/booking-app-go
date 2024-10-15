package main

import (
	"fmt"
	"strings"
    "booking-app/helper"
)

const MaxTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50

func main(){

    greetUser()

    var firstName string
    var lastName string
    var email string
    var userTickets uint
    var userLocation string
    var bookings = []string{}

    for {

        firstName, lastName, email, userLocation, userTickets  = getUserInputs()


        isValidName, isValidEmail,isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

        if !isValidName{
            fmt.Println("The First name or Last name entered is invalid")
            continue
        } else if !isValidEmail{
            fmt.Println("The Email entered is invalid")
            continue
        } else if !isValidTicketNumber{
            fmt.Println("The Ticket Number entered is invalid")
            continue
        }

        // DO something with the the userLocation
        switch userLocation {
            case "India":
                fmt.Println("User is from India")
                // Custom processes
            case "Fiji, Australia, New Zealand":
                fmt.Println("User is from Oceania")
            default:
                fmt.Println("User is from other country")
        }

        // Validation when user enter more than available tickets
        if userTickets > remainingTickets{
            fmt.Printf("Sorry! We dont have %v tickets available, please select a count under %v\n", userTickets, remainingTickets)
            
            // skip this booking and continue for another input
            continue
        }

        remainingTickets = remainingTickets - userTickets
        
        // add the booking to slice object
        bookings = append(bookings, firstName + " " + lastName)

        fmt.Printf("Thankyou %v %v for booking %v tickets\nYou will recieve a confirmation email at %v shortly\n", firstName, lastName, userTickets, email)

        fmt.Printf("Remaining tickets for %v conferece = %v\n", conferenceName, remainingTickets)

        // For data privacy only show first name of the booking holders
        firstNames := [] string {}
        for _, holder := range(bookings){
            firstName := strings.Fields(holder)[0]
            firstNames = append(firstNames, firstName)
        }
        fmt.Printf("List of all booking so far.. %v\n", firstNames)


        // CHeck tickets full
        if remainingTickets == 0 {
            fmt.Println("Sorry tickets for this year have been sold out.. come back next year")
        }else{
            fmt.Println("Welcome to next booking....")
        }

    }



}



func greetUser(){
    fmt.Printf("Welcome to the %v booking app...\n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v are available\n", MaxTickets, remainingTickets)
    fmt.Print("Get your tickets here to attend\n")
}


func getUserInputs()(
    firstName string,
    lastName string,
    email string,
    userLocation string,
    userTickets uint){

    // uses pointer conecpt to access the memory location of vaiable
    fmt.Println("Enter the First Name: ")
    fmt.Scan(&firstName)

    fmt.Println("Enter the Last Name:")
    fmt.Scan(&lastName)

    fmt.Println("Enter the Email address:")
    fmt.Scan(&email)

    fmt.Println("Please enter your location")
    fmt.Scan(&userLocation)

    fmt.Println("Enter the number of tickets needed:")
    fmt.Scan(&userTickets)

    return firstName, lastName, email, userLocation, userTickets
}

