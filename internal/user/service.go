package users

import (
	"errors"
	"sync"

	"go.uber.org/zap"

	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
)

type Users struct {
	db     *database.Database
	once   sync.Once
	ready  bool
	logger *zap.Logger
}

var (
	instance *Users
	once     sync.Once
	ready    bool
)

func InitUser(db *database.Database, logger *zap.Logger) {
	once.Do(func() {
		instance = &Users{
			db:     db,
			logger: logger,
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
	return u.createUser(user)
}

func (u *Users) GetOrCreateUser(user *model.User) (*model.User, error) {
	userInDB, err := u.getUserByPhonenumber(*user.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if userInDB != nil {
		return userInDB, nil
	}
	user.Version = 1
	return user, u.createUser(user)
}

func (u *Users) GetUserByPhonenumber(phonenumber string) (*model.User, error) {
	return u.getUserByPhonenumber(phonenumber)
}
