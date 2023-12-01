package api

import (
	"github.com/Inf85/go-project-demo/models/users"
	"gorm.io/gorm"
)

type API struct {
	users users.Storage
}

func NewAPI(db *gorm.DB) *API {
	userRep, _ := users.NewUserManager(db)
	return &API{
		users: users.NewService(userRep),
	}
}
