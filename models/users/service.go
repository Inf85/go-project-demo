package users

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository userManager
}

func (u *UserService) FindUser(ctx context.Context, username string) *User {
	return u.userRepository.FindUser(username)
}

func (u *UserService) FindUserByUUID(ctx context.Context, uuid string) *User {
	return u.userRepository.FindUserByUUID(uuid)
}

func (u *UserService) AddUser(ctx context.Context, username, pasword string) *User {
	return u.userRepository.AddUser(username, pasword)
}

func (u *UserService) HasUser(ctx context.Context, username string) bool {

	return u.userRepository.HasUser(username)
}

// CheckPassword - compare a hashed password with a possible plaintext equivalent
func (u *UserService) CheckPassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false
	}
	return true
}

func NewService(rep userManager) Storage {
	return &UserService{
		userRepository: rep,
	}
}
