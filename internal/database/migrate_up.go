package main

import (
	"gitlab.bd.com/new-argos-be/common"
	"gitlab.bd.com/new-argos-be/internal/models"
)

func main() {
	db, err := common.NewDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Local{}, &models.Person{}, &models.Terminal{})

	if err != nil {
		panic(err)

	}

	err = db.AutoMigrate(&models.PersonTerminal{})
	if err != nil {
		panic(err)

	}
}
