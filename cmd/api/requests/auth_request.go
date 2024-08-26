package requests

import (
	"time"

	"github.com/google/uuid"
)

type RegisterPersonRequest struct {
	Register int       `json:"matricula" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Photo    string    `json:"photo" validate:"required"`
	Notes    string    `json:"notes" validate:"required"`
	Password string    `json:"password" validate:"required"`
	DeviceID string    `json:"device_id" validate:"required"`
	Valid    time.Time `json:"valid" validate:"required"`
}

type RegisterTerminalRequest struct {
	Https    *bool  `json:"https" validate:"required"`
	PortCGI  int16  `json:"port_cgi"`
	PortRTSP int16  `json:"port_rtsp"`
	PortSDK  int16  `json:"port_sdk"`
	Name     string `json:"name" validate:"required"`
	IPv4     string `json:"ipv4" validate:"required"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type RegisterLocalRequest struct {
	Name string `json:"name" validate:"required"`
}

type RegisterLocalResponse struct {
	Name    string    `json:"local_name"`
	LocalID uuid.UUID `json:"local_id"`
}
