package main
import "fmt"

func main(){

    const MaxTickets = 50
    conferenceName := "Go Conference"
    var remainingTickets uint = 25

    fmt.Printf("Welcome to the %v booking app...\n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v are available\n", MaxTickets, remainingTickets)
    fmt.Print("Get your tickets here to attend\n")

    var firstName string
    var lastName string
    var email string
    var userTickets uint

    // uses pointer conecpt to access the memory location of vaiable
    fmt.Println("Enter the First Name: ")
    fmt.Scan(&firstName)

    fmt.Println("Enter the Last Name:")
    fmt.Scan(&lastName)

    fmt.Println("Enter the Email address:")
    fmt.Scan(&email)

    fmt.Println("Enter the number of tickets needed:")
    fmt.Scan(&userTickets)

    remainingTickets = remainingTickets - userTickets

    fmt.Printf("Thankyou %v %v for booking %v tickets\nYou will recieve a confirmation email at %v shortly\n", firstName, lastName, userTickets, email)

    fmt.Printf("Remaining tickets for %v conferece = %v", conferenceName, remainingTickets)




}


