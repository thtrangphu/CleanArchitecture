package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        ID
	Email     string
	Password  string
	FullName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateNewUser(email, password, fullName string) (*User, error) {
	user := &User{
		ID:        NewID(),
		Email:     email,
		FullName:  fullName,
		CreatedAt: time.Now(),
	}
	if !user.CheckUserInfor() {
		return nil, ErrInvalidEntity
	}
	pass, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = pass

	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) CheckUserInfor() bool {
	if user.Email == "" || user.Password == "" || user.FullName == "" {
		return false
	}
	return true
}
