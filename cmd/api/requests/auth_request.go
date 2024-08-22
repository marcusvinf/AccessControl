package requests

import (
	"time"
)

type RegisterPersonRequest struct {
	Name     string    `json:"name"`
	Photo    []byte    `json:"photo"`
	Valid    time.Time `json:"valid"`
	Notes    string    `json:"notes"`
	Password string    `json:"password"`
	Register int       `json:"matricula"`
	DeviceID string    `json:"device_id"`
}
