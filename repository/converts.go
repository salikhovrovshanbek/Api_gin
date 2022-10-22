package repository

import "github.com/google/uuid"

type Phone struct {
	PhoneID   uuid.UUID `json:"phone_id"`
	PhoneName string    `json:"phone_name"`
	PhoneYear int       `json:"phone_year"`
}
