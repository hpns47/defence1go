package reservations

import (
	"errors"

)

type Reservation struct {
	ID string
	GuestID string
	Table int
	People int
	Status string
}

type Manager struct {
	items map[string]Reservation
}

func NewManager() *Manager {
	return &Manager{
		items: make(map[string]Reservation),
	}
}

func (m *Manager) AddReservation(reservation Reservation) error {
	if reservation.ID == "" {
		return errors.New("reservation ID cannot be empty")
	}

	if _,exists := m.items[reservation.ID]; exists {
		return errors.New("reservation with this ID already exists")
	}
	if reservation.GuestID == "" {
		return errors.New("guest ID cannot be empty")
	}
	if reservation.People <= 0 || reservation.People > 10 {
		return errors.New("number of people must be greater than zero")
	}
	if reservation.Table <= 0 || reservation.Table > 50 {
		return errors.New("table number must be between 1 and 50")
	}
	if reservation.Status == "" {
		reservation.Status = "ACTIVE"
	}
	
	m.items[reservation.ID] = reservation
	return nil
}

func (m *Manager) CancelReservation(reservationID string) error {
	reservation, exists := m.items[reservationID]
	if !exists {
		return errors.New("reservation not found")
	}
	
	if reservation.Status == "CANCELLED" {
		return errors.New("reservation is already cancelled")
	}
	reservation.Status = "CANCELLED"
	m.items[reservationID] = reservation
	return nil
}

func (m *Manager) ListReservations() []Reservation {
	var reservations []Reservation
	for _, reservation := range m.items {
		if reservation.Status != "CANCELLED" {
			reservations = append(reservations, reservation)
		}
	}
	return reservations
}

func (manager *Manager) ListByGuest(guestID string) []Reservation {
	var reservations []Reservation
	for _, reservation := range manager.items {
		if reservation.GuestID == guestID && reservation.Status != "CANCELLED" {
			reservations = append(reservations, reservation)
		}
	}
	return reservations
}


func (m *Manager) TableUsage() map[int]int {
	usage := make(map[int]int)
	for _, reservation := range m.items {
		if reservation.Status != "CANCELLED" {
			usage[reservation.Table]++
		}
	}
	return usage
}