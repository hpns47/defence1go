package guests

import "fmt"

type Guest interface {
	GetID()	 string
	DisplayName() string
}

type RegularGuest struct {
	ID   string
	name string
}

type VIPGuest struct {
	ID        string
	name  string
	level int
}

func (guest RegularGuest) GetID() string {
	return guest.ID
}
func (guest RegularGuest) DisplayName() string {
	return fmt.Sprintf("%s ", guest.name)
}

func (guest VIPGuest) GetID() string {
	return guest.ID
}

func (guest VIPGuest) DisplayName() string {
	return fmt.Sprintf("%s | vip: %s ", guest.name,guest.level)
}

func formatGuest(guest Guest) string {
	switch v := guest.(type) {
	case RegularGuest:
		return fmt.Sprintf("Regular Guest: %s (ID: %s)", v.name, v.ID)
	case VIPGuest:
		return fmt.Sprintf("VIP Guest: %s (ID: %s, Level: %d)", v.name, v.ID, v.level)
	default:
		return "Unknown Guest Type"
	}
}