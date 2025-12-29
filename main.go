package main

import (
	"fmt"
	"github.com/hpns47/defence1go/Guests"
	"github.com/hpns47/defence1go/reservations"
)

var guestsRegistry = make(map[string]guests.Guest)
var reservationManager = reservations.NewManager()
func main(){
	
	for {
		fmt.Println("Hello, World!")
		fmt.Println("1. add guest regular")
		fmt.Println("2. add guest vip")
		fmt.Println("3. create reservation")
		fmt.Println("4. cancel reservation")
		fmt.Println("5. List Active Reservations")
		fmt.Println("6. List Reservations By Guest")
		fmt.Println("7. show talbe usage")
		fmt.Println("8. exit")

		var choice int
		fmt.Println("Enter your choice: ")
		_,err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}
			
		switch choice {
		case 1:
			addRegularGuest()
		case 2:
			addVIPGuest()
		case 3:
			createReservation()
		case 4:
			cancelReservation()
		case 5:
			listActiveReservations()
		case 6:
			listReservationsByGuest()
		case 7:
			showTableUsage()
		case 8:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

}
func addRegularGuest() {
	var id, name string
	fmt.Print("Enter Guest ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Guest Name: ")
	fmt.Scan(&name)

	if id == "" || name == "" {
		fmt.Println("Guest ID and Name cannot be empty.")
		return
	}
	if _, exists := guestsRegistry[id]; exists {
		fmt.Println("Guest with this ID already exists.")
		return
	}
	guest := guests.RegularGuest{ID: id, Name: name}
	guestsRegistry[id] = guest
	fmt.Printf("Regular guest added: %s\n", guests.FormatGuest(guest))
}

func addVIPGuest() {
	var id, name string
	var level int
	fmt.Print("Enter Guest ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Guest Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter VIP Level: ")
	fmt.Scan(&level)

	if (level < 1 || level > 3) {
		fmt.Println("Invalid VIP level. It should be between 1 and 3.")
		return
	}

	if id == "" || name == "" {
		fmt.Println("Guest ID and Name cannot be empty.")
		return
	}
	if _, exists := guestsRegistry[id]; exists {
		fmt.Println("Guest with this ID already exists.")
		return
	}


	guest := guests.VIPGuest{ID: id, Name: name, Level: level}
	guestsRegistry[id] = guest
	fmt.Printf("VIP guest added: %s\n", guests.FormatGuest(guest))
}

func createReservation() {

	var id, guestID string
	var people, table int

	fmt.Print("Enter Reservation ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Guest ID: ")
	fmt.Scan(&guestID)
	fmt.Print("Enter Number of People: ")
	fmt.Scan(&people)
	fmt.Print("Enter Table Number: ")
	fmt.Scan(&table)
	guest, exists := guestsRegistry[guestID]
	if !exists {
		fmt.Println("Guest not found. Please add the guest first.")
		return
	}
	
	reservation := reservations.Reservation{
		ID:       id,
		GuestID:  guest.GetID(),
		Table:    table,
		People:   people,
		Status:   "ACTIVE",
	}
	err := reservationManager.AddReservation(reservation)
	if err != nil {
		fmt.Printf("Error creating reservation: %s\n", err.Error())
	} else {
		fmt.Println("Reservation created successfully.")
	}
}

func cancelReservation() {
	var id string
	fmt.Print("Enter Reservation ID to cancel: ")
	fmt.Scan(&id)
	
	err := reservationManager.CancelReservation(id)
	if err != nil {
		fmt.Printf("Error cancelling reservation: %s\n", err.Error())
	} else {
		fmt.Println("Reservation cancelled successfully.")
	}
}


func listActiveReservations() {
	activeReservations := reservationManager.ListReservations()
	fmt.Print("Active Reservations:\n")
	for _, res := range activeReservations {
		fmt.Printf("ID: %s | GuestID: %s | Table: %d | People: %d | Status: %s\n",
			res.ID, res.GuestID, res.Table, res.People, res.Status)
	}
}

func listReservationsByGuest() {
	var guestID string
	fmt.Print("Enter Guest ID to list reservations: ")
	fmt.Scan(&guestID)

	reservations := reservationManager.ListByGuest(guestID)
	fmt.Printf("Reservations for Guest ID %s:\n", guestID)
	for _, res := range reservations {
		fmt.Printf("ID: %s | Table: %d | People: %d | Status: %s\n",
			res.ID, res.Table, res.People, res.Status)
	}
}

func showTableUsage() {
	tableUsage := reservationManager.TableUsage()
	fmt.Println("Table Usage:")
	for table, count := range tableUsage {
		fmt.Printf("Table %d: %d reservations\n", table, count)
	}
}