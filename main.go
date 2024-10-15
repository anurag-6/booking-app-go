package main

import (
	"booking-app/helper"
	"fmt"
    "sync"
    "time"
)

const MaxTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50


// Define a struct for booking details
type Booking struct {
    firstName   string
    lastName    string
    email       string
    userTickets uint
}


var wg =  sync.WaitGroup{}

func main(){

    greetUser()

    var firstName string
    var lastName string
    var email string
    var userTickets uint
    var userLocation string
    // create a slice of maps with least value 0, means it can be evolved later
    // Use make to create a slice with initial capacity for 0 elements
    // bookings := make([]map[string]string, 0)

    bookings := make([]Booking, 0)


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

        // Create a new Booking struct and populate the fields
        booking := Booking{
            firstName:   firstName,
            lastName:    lastName,
            email:       email,
            userTickets: userTickets,
        }
        // Append the booking struct to the bookings slice
        bookings = append(bookings, booking)

        // Old way
        // Create a new map object
        // bookingMap := make(map[string]string)
        // bookingMap["firstName"] = firstName
        // bookingMap["lastName"] = lastName
        // bookingMap["email"] = email
        // bookingMap["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
        // bookings = append(bookings, bookingMap)

        wg.Add(1)
        go sendTicket(booking)

        fmt.Printf("Remaining tickets for %v conferece = %v\n", conferenceName, remainingTickets)

        // For data privacy only show first name of the booking holders
        firstNames := [] string {}
        for _, holder := range(bookings){
            firstName := holder.firstName
            firstNames = append(firstNames, firstName)
        }
        fmt.Printf("List of all booking so far.. %v\n", firstNames)


        // CHeck tickets full
        if remainingTickets == 0 {
            fmt.Println("Sorry tickets for this year have been sold out.. come back next year")
        }else{
            fmt.Println("Welcome to next booking....")
        }
        wg.Wait()

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

func sendTicket(booking Booking){
    time.Sleep(50 * time.Second)
    fmt.Println("#####################")
    fmt.Printf("Thankyou %v %v for booking %v tickets\nYou will recieve a confirmation email at %v shortly\n", booking.firstName, booking.lastName, booking.userTickets, booking.email)
    fmt.Println("#####################")
    wg.Done()

}