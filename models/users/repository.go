package users

import "gorm.io/gorm"
import uuid "github.com/satori/go.uuid"
import "golang.org/x/crypto/bcrypt"

// UserManager struct
type userManager struct {
	db *gorm.DB
}

/*
*
 */
func NewUserManager(db *gorm.DB) (userManager, error) {
	db.AutoMigrate(&User{})
	usermgr := userManager{}
	usermgr.db = db

	return usermgr, nil
}

// HasUser - Check if the given username exists.
func (um *userManager) HasUser(username string) bool {
	if err := um.db.Where("username=?", username).Find(&User{}).Error; err != nil {
		return false
	}
	return true
}

// FindUser -
func (um *userManager) FindUser(username string) *User {
	user := User{}
	um.db.Where("username=?", username).Find(&user)
	return &user
}

// FindUserByUUID -
func (um *userManager) FindUserByUUID(uuid string) *User {
	user := User{}
	um.db.Where("uuid=?", uuid).Find(&user)
	return &user
}

// AddUser - Creates a user and hashes the password
func (um *userManager) AddUser(username, password string) *User {
	passwordHash := um.HashPassword(username, password)
	guid := uuid.NewV4()
	user := &User{
		UserName: username,
		Password: passwordHash,
		UUID:     guid.String(),
	}
	um.db.Create(&user)
	return user
}

// HashPassword - Hash the password (takes a username as well, it can be used for salting).
func (um *userManager) HashPassword(username, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}
