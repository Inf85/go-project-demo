package users

import "context"

type Storage interface {
	HasUser(ctx context.Context, username string) bool
	FindUser(ctx context.Context, username string) *User
	FindUserByUUID(ctx context.Context, uuid string) *User
	AddUser(ctx context.Context, username, pasword string) *User
	CheckPassword(hashedPassword, password string) bool
}
