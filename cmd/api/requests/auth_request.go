package requests

import (
	"time"
)

type RegisterUserRequest struct {
	IsAdmin  bool   `json:"is_admin" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gt=12,containsany=uppercase,containsany=lowercase,containsany=numeric,containsany=!@#?_[]"`
}

type RegisterPersonRequest struct {
	Register uint      `json:"matricula" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Photo    string    `json:"photo" validate:"required"`
	Notes    string    `json:"notes" validate:"required"`
	Password string    `json:"password" validate:"required"`
	DeviceID string    `json:"device_id" validate:"required"`
	Valid    time.Time `json:"valid" validate:"required"`
}

type RegisterTerminalRequest struct {
	Https        *bool  `json:"https" validate:"required"`
	PortCGI      uint16 `json:"port_cgi"`
	PortRTSP     uint16 `json:"port_rtsp"`
	PortSDK      uint16 `json:"port_sdk"`
	TerminalName string `json:"terminal_name" validate:"required"`
	LocalName    string `json:"local_name" validate:"required"`
	IPv4         string `json:"ipv4" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Username     string `json:"username" validate:"required"`
}

type RegisterLocalRequest struct {
	Name string `json:"name" validate:"required"`
}
