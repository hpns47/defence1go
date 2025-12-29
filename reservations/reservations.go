package reservations

import (
	"errors"

)

type Reservation struct {
	ID string
	guestID string
	table int
	people int
	status string
}

type Manager struct {
	items map[string]Reservation
}

func NewManager() *Manager {
	return &Manager{
		items: make(map[string]Reservation),
	}
}

func (m *Manager) addReservation(reservation Reservation) error {
	if reservation.ID == "" {
		return errors.New("reservation ID cannot be empty")
	}

	if _,exists := m.items[reservation.ID]; exists {
		return errors.New("reservation with this ID already exists")
	}
	if reservation.guestID == "" {
		return errors.New("guest ID cannot be empty")
	}
	if reservation.people <= 0 || reservation.people > 10 {
		return errors.New("number of people must be greater than zero")
	}
	if reservation.table <= 0 || reservation.table > 50 {
		return errors.New("table number must be between 1 and 50")
	}
	if reservation.status == "" {
		reservation.status = "ACTIVE"
	}
	
	m.items[reservation.ID] = reservation
	return nil
}

func (m *Manager) cancelReservation(reservationID string) error {
	reservation, exists := m.items[reservationID]
	if !exists {
		return errors.New("reservation not found")
	}
	
	if reservation.status == "CANCELLED" {
		return errors.New("reservation is already cancelled")
	}
	reservation.status = "CANCELLED"
	m.items[reservationID] = reservation
	return nil
}

func (m *Manager) listReservations() []Reservation {
	var reservations []Reservation
	for _, reservation := range m.items {
		if reservation.status != "CANCELLED" {
			reservations = append(reservations, reservation)
		}
	}
	return reservations
}

func (manager *Manager) listByGuest(guestID string) []Reservation {
	var reservations []Reservation
	for _, reservation := range manager.items {
		if reservation.guestID == guestID && reservation.status != "CANCELLED" {
			reservations = append(reservations, reservation)
		}
	}
	return reservations
}


func (m *Manager) tableUsage() map[int]int {
	usage := make(map[int]int)
	for _, reservation := range m.items {
		if reservation.status != "CANCELLED" {
			usage[reservation.table]++
		}
	}
	return usage
}