package models

type Host struct {
	HID          int    `json:"HID"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	PhoneNumber  string `json:"phone_number"`
}

type VendingMachine struct {
	VID             int     `json:"VID"`
	HID             int     `json:"HID"`
	TID             int     `json:"TID"`
	Location        string  `json:"location"`
	MachineName     string  `json:"machine_name"`
	Status          string  `json:"status"`
	LastServiceDate string  `json:"last_service_date"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
}

type MachineType struct {
	TID           int    `json:"TID"`
	TypeName      string `json:"type_name"`
	Capacity      int    `json:"capacity"`
	Refrigeration bool   `json:"refrigeration"`
}
