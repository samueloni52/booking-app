package main

import (
	"Strings"
	"fmt"
	"time"
)

const ConferenceTicket = 50
var ConferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([] userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numerOfTickets uint
}
func main() {

	//greetings user
	greetUsers()

	//asking for user input

	for {
		firstName, lastName, email, userTickets :=  getuserInputs()

		 //user input validation or validating the input

		 isVaidName, isValidEmail, isValidTicketNumber :=  userInputValidation(firstName, lastName, email, userTickets,)

		if isVaidName && isValidEmail && isValidTicketNumber {

			bookTicket (firstName, lastName, email, userTickets,)
			go sendTicket(userTickets, firstName, lastName, email)
			

			firstNames := getFirstNames ()
			fmt.Printf("these are all the bookings: %v\n", firstNames)

			if remainingTickets == 0 {
			// end program
			fmt.Println("our conference is booked out. come back next year")
			break
			}

		}else{
			if !isVaidName{
				fmt.Println("your first name or last name you entered is not correct")
			}
			if !isValidEmail{
				fmt.Println("your email address is not correct")
			}
			if !isValidTicketNumber{
				fmt.Println("number of ticket is not valid")
			}
			
		}	
 
	}
}

func greetUsers (){
	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v left\n", ConferenceTicket, remainingTickets)
	fmt.Println("get your ticket here to attend")
}

func getFirstNames () []string{
	var firstNames []string // firstName :=[]string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
	
}
func userInputValidation (firstName string, lastName string, email string, userTickets uint,) (bool, bool, bool) {
	isVaidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isVaidName, isValidEmail, isValidTicketNumber	
}
func getuserInputs () (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("please provide the following details for your booking...")

	fmt.Println("Enter your First Name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter your Last Name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter your Email address:")
	fmt.Scanln(&email)
	fmt.Println("Enter numbers of ticket:")
	fmt.Scanln(&userTickets)
	return firstName, lastName,email, userTickets
}
func bookTicket (firstName string, lastName string, email string, userTickets uint,){
	remainingTickets -= userTickets
 // creating a struct for a user
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numerOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a Confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainiing for %v\n", remainingTickets, ConferenceName)
}

func sendTicket( userTickets uint, FirstNames string, lastName string, email string){
	time.Sleep(10* time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, FirstNames, lastName)
	fmt.Println("##################")
	fmt.Printf("sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
}