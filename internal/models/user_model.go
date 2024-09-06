package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	IsAdmin  bool   `gorm:"type:boolean;not null" json:"is_admin"`
	Username string `gorm:"type:varchar(15);not null;unique" json:"username"`
	Email    string `gorm:"type:varchar(320);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

type Local struct {
	LocalID   uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"local_id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"local_name"`
	Terminals []Terminal `gorm:"foreignKey:LocalID;" json:"-"`
}

type Person struct {
	PersonID  uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Photo     []byte     `gorm:"type:bytea"`
	Valid     time.Time  `gorm:"type:timestamptz"`
	Notes     string     `gorm:"type:text"`
	Password  string     `gorm:"type:varchar(255)"`
	Register  uint       `gorm:"type:integer;unique"`
	DeviceID  string     `gorm:"type:varchar(32);not null;unique"`
	Terminals []Terminal `gorm:"many2many:person_terminal;"`
}

type Terminal struct {
	TerminalID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string    `gorm:"type:varchar(255);not null" json:"person_name"`
	IPv4       string    `gorm:"type:varchar(15);not null;unique" json:"ipv4"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Password   string    `gorm:"type:varchar(255);not null" json:"password"`
	PortSDK    uint16    `gorm:"type:integer;default:37777" json:"port_sdk"`
	LocalID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	PortCGI    uint16    `gorm:"type:integer;default:80" json:"port_cgi"`
	PortRTSP   uint16    `gorm:"type:integer;default:554" json:"port_rtsp"`
	Https      bool      `gorm:"type:boolean;not null" json:"https"`
	Person     []Person  `gorm:"many2many:person_terminal;" json:"-"`
}

type PersonTerminal struct {
	PersonID   uuid.UUID `gorm:"type:uuid;not null"`
	TerminalID uuid.UUID `gorm:"type:uuid;not null"`
	Person     Person    `gorm:"foreignKey:PersonID;references:PersonID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Terminal   Terminal  `gorm:"foreignKey:TerminalID;references:TerminalID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	BaseModel
}

func (Person) TableName() string {
	return "person"
}

func (PersonTerminal) TableName() string {
	return "person_terminal"
}

func (PersonTerminal) PrimaryKey() []string {
	return []string{"PersonID", "TerminalID"}
}
