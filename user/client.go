package user

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/skamenetskiy/jsonapi"
)

var (
	client = jsonapi.NewClient(Addr)
)

// Create is creating a new user
func Create(u *User) (*User, error) {
	if u == nil {
		return nil, errors.New("user is nil")
	}
	u.ID = 0
	res, err := client.Post("/", u)
	if err != nil {
		return nil, err
	}
	newUser := new(User)
	if err := res.ReadJSON(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

// GetByID returns a user by ID or error
func GetByID(id uint64) (*User, error) {
	return getBy("id", strconv.FormatUint(id, 10))
}

// GetByLogin returns a user by login or error
func GetByLogin(login string) (*User, error) {
	return getBy("login", login)
}

func getBy(param string, value string) (*User, error) {
	res, err := client.Get("/" + param + "/" + value)
	if err != nil {
		return nil, err
	}
	u := new(User)
	if err := res.ReadJSON(u); err != nil {
		return nil, err
	}
	return u, nil
}
