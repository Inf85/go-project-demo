package users

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	UserName   string `gorm:"not null;unique" json:"user_name"`
	Email      string `gorm:"unique" json:"email"`
	Password   string `gorm:"not null" json:"-"`
	UUID       string `gorm:"not null;unique" json:"uuid"`
}
