package main
import "fmt"

func main(){

    const MaxTickets = 50
    conferenceName := "Go Conference"
    var remainingTickets uint = 25

    fmt.Printf("Welcome to the %v booking app...\n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v are available\n", MaxTickets, remainingTickets)
    fmt.Print("Get your tickets here to attend")

    var userName string = "Anurag"
    var userTickets = 2
    fmt.Printf("\nUser %v has booked %v ticket (s)\n", userName, userTickets)


}


