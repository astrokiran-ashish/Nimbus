package users

import (
	"errors"
	"sync"

	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
)

type UsersI interface {
	CreateUser(user *model.User) error
	// GetUserByEmail(email string) (*model.User, error)
	// GetUserByID(id string) (*model.User, error)
	// UpdateUser(user *model.User) error
	// DeleteUser(id string) error
	GetOrCreateUser(user *model.User) (*model.User, error)
}

type Users struct {
	db    *database.Database
	once  sync.Once
	ready bool
}

var (
	instance *Users
	once     sync.Once
	ready    bool
)

func Init(db *database.Database) {
	once.Do(func() {
		instance = &Users{
			db: db,
		}
		ready = true
	})
}

func GetInstance() (*Users, error) {
	if !ready {
		return nil, errors.New("Users instance not initialized")
	}
	return instance, nil
}

func (u *Users) CreateUser(user *model.User) error {
	return nil
}

func (u *Users) GetOrCreateUser(user *model.User) (*model.User, error) {
	return &model.User{}, nil
}
