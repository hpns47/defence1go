package guests

import "fmt"

type Guest interface {
	GetID()	 string
	DisplayName() string
}

type RegularGuest struct {
	ID   string
	Name string
}

type VIPGuest struct {
	ID        string
	Name  string
	Level int
}

func (guest RegularGuest) GetID() string {
	return guest.ID
}
func (guest RegularGuest) DisplayName() string {
	return fmt.Sprintf("%s ", guest.Name)
}

func (guest VIPGuest) GetID() string {
	return guest.ID
}

func (guest VIPGuest) DisplayName() string {
	return fmt.Sprintf("%s | vip: %s ", guest.Name,guest.Level)
}

func FormatGuest(guest Guest) string {
	switch v := guest.(type) {
	case RegularGuest:
		return fmt.Sprintf("Regular Guest: %s (ID: %s)", v.Name, v.ID)
	case VIPGuest:
		return fmt.Sprintf("VIP Guest: %s (ID: %s, Level: %d)", v.Name, v.ID, v.Level)
	default:
		return "Unknown Guest Type"
	}
}