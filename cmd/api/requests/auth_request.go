package requests

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gt=12,containsany=uppercase,containsany=lowercase,containsany=numeric,containsany=!@#?_[]"`
}

type RegisterUserRequest struct {
	IsAdmin  bool   `json:"is_admin" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gt=12,containsany=uppercase,containsany=lowercase,containsany=numeric,containsany=!@#?_[]"`
}

type RegisterPrePersonRequest struct {
	Register uint   `json:"matricula"`
	Name     string `json:"name" validate:"required"`
	Photo    string `json:"photo" validate:"required"`
	Notes    string `json:"notes"`
	Password string `json:"password"`
	DeviceID string `json:"device_id" validate:"required"`
	Valid    string `json:"valid" validate:"required"`
}

type RegisterFullPersonRequest struct {
	RegisterPrePersonRequest
	TerminalNames []string `json:"terminal_names"`
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
