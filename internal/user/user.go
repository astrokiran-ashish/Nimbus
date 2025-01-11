package users

import (
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
)

type IUsers interface {
	CreateUser(user *model.User) error
	GetUserByPhonenumber(phonenumber string) (*model.User, error)
	// GetUserByID(id string) (*model.User, error)
	// UpdateUser(user *model.User) error
	// DeleteUser(id string) error
	GetOrCreateUser(user *model.User) (*model.User, error)
}
