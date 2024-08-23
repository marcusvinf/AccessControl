package requests

import (
	"time"
)

type RegisterPersonRequest struct {
	Name     string    `json:"name" validate:"required"`
	Photo    []byte    `json:"photo" validate:"required"`
	Valid    time.Time `json:"valid" validate:"required"`
	Notes    string    `json:"notes" validate:"required"`
	Password string    `json:"password" validate:"required"`
	Register int       `json:"matricula" validate:"required"`
	DeviceID string    `json:"device_id" validate:"required"`
}
