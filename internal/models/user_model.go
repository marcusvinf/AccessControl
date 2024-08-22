package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Local struct {
	LocalID   uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"type:varchar(100);not null"`
	Terminals []Terminal `gorm:"foreignKey:LocalID;"`
}

type Person struct {
	PersonID  uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Photo     []byte     `gorm:"type:bytea"`
	Valid     time.Time  `gorm:"type:timestamptz"`
	Notes     string     `gorm:"type:text"`
	Password  string     `gorm:"type:varchar(255)"`
	Register  int        `gorm:"type:integer;unique"`
	DeviceID  string     `gorm:"type:varchar(32);not null;unique"`
	Terminals []Terminal `gorm:"many2many:person_terminal;"`
}

type Terminal struct {
	TerminalID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string    `gorm:"type:varchar(255);not null"`
	IPv4       string    `gorm:"type:varchar(15);not null;unique"`
	Username   string    `gorm:"type:varchar(255);not null"`
	Password   string    `gorm:"type:varchar(255);not null"`
	PortSDK    int16     `gorm:"type:smallint;default:37777"`
	LocalID    uuid.UUID `gorm:"type:uuid;not null"`
	PortCGI    int16     `gorm:"type:smallint;default:80"`
	PortRTSP   int16     `gorm:"type:smallint;default:554"`
	Https      bool      `gorm:"type:boolean;not null"`
	Person     []Person  `gorm:"many2many:person_terminal;"`
}

type PersonTerminal struct {
	PersonID   uuid.UUID `gorm:"type:uuid;not null"`
	TerminalID uuid.UUID `gorm:"type:uuid;not null"`
	Person     Person    `gorm:"foreignKey:PersonID;references:PersonID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Terminal   Terminal  `gorm:"foreignKey:TerminalID;references:TerminalID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	gorm.Model
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
