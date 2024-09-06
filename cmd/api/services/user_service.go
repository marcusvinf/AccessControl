package services

import (
	"errors"

	"github.com/google/uuid"
	"gitlab.bd.com/new-argos-be/cmd/api/requests"
	"gitlab.bd.com/new-argos-be/common"
	"gitlab.bd.com/new-argos-be/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (u UserService) RegisterTerminal(terminalRequest *requests.RegisterTerminalRequest, localID uuid.UUID) (*models.Terminal, error) {
	hashedPassword, err := common.HashPassword(terminalRequest.Password)
	if err != nil {
		return nil, err
	}
	createdTerminal := models.Terminal{
		Name:     terminalRequest.TerminalName,
		IPv4:     terminalRequest.IPv4,
		Username: terminalRequest.Username,
		Password: hashedPassword,
		Https:    *terminalRequest.Https,
		LocalID:  localID,
	}
	result := u.db.Create(&createdTerminal)
	if result.Error != nil {
		return nil, errors.New("registro de local falhou")
	}

	return &createdTerminal, nil
}

func (u UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := u.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u UserService) GetTerminalByIp(ipv4 string) (*models.Terminal, error) {
	var terminal models.Terminal
	result := u.db.Where("ipv4 = ?", ipv4).First(&terminal)
	if result.Error != nil {
		return nil, result.Error
	}
	return &terminal, nil
}

func (u UserService) GetLocalByName(name string) (*models.Local, error) {
	var local models.Local
	result := u.db.Where("lower(name) = lower(?)", name).First(&local)
	if result.Error != nil {
		return nil, result.Error
	}
	return &local, nil
}

func (u UserService) GetAllLocations() (*[]models.Local, error) {
	var locations []models.Local
	results := u.db.Find(&locations)

	if results.Error != nil {
		return nil, results.Error
	}
	return &locations, nil

}

func (u UserService) RegisterLocal(localRequest *requests.RegisterLocalRequest) (*models.Local, error) {
	createdLocal := models.Local{
		Name: localRequest.Name,
	}
	result := u.db.Create(&createdLocal)
	if result.Error != nil {
		return nil, errors.New("registro de local falhou")
	}

	return &createdLocal, nil
}
