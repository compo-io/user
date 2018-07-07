package user

import (
	"github.com/pkg/errors"
)

const (
	Addr = "user:9500"
)

// User model
//go:generate reform
//go:generate easyjson
//reform:user
//easyjson:json
type User struct {
	ID           uint64 `reform:"id,pk" json:"id"`
	Login        string `reform:"login" json:"login"`
	Password     string `reform:"-" json:"password"`
	PasswordHash string `reform:"password" json:"-"`
}

func (u *User) ValidateOnRegistration() error {
	if u.Login == "" {
		return errors.New("bad login")
	}
	if u.Password == "" {
		return errors.New("bad password")
	}
	return nil
}

func (u *User) ValidateOnLogin() error {
	if u.Login == "" {
		return errors.New("bad login")
	}
	if u.Password == "" {
		return errors.New("bad password")
	}
	return nil
}
