package terminalactions

import "time"

type TerminalData struct {
	UserID            string
	TermUser          string
	Password          string
	Scheme            string
	TermAdminUser     string
	TermAdminPassword string
	IpAddress         string
	ValidFrom         time.Time
	ValidTo           time.Time
}
