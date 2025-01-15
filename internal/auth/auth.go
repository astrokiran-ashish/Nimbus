package auth

import (
	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/astrokiran/nimbus/internal/common/services"
	users "github.com/astrokiran/nimbus/internal/user"
	"go.uber.org/zap"
)

type Auth struct {
	db         *database.Database
	Users      users.IUsers
	SMSService services.ISMSService
	logger     *zap.Logger
}

func NewAuth(db *database.Database, users users.IUsers, smsService services.ISMSService, logger *zap.Logger) *Auth {
	return &Auth{
		db:         db,
		Users:      users,
		SMSService: smsService,
		logger:     logger,
	}
}
