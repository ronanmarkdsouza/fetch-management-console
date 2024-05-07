package models

type Host struct {
	HID          int    `json:"HID"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	PhoneNumber  string `json:"phone_number"`
}
