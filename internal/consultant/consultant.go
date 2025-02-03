package consultant

import (
	"github.com/astrokiran/nimbus/internal/auth"
	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/astrokiran/nimbus/internal/common/services"
	users "github.com/astrokiran/nimbus/internal/user"
)

type Consultant struct {
	db         *database.Database
	auth       *auth.Auth
	user       users.IUsers
	SMSService services.ISMSService
}

func NewConsultant(db *database.Database, auth *auth.Auth, user *users.Users, smsService *services.SMSService) *Consultant {
	return &Consultant{
		db:         db,
		auth:       auth,
		user:       user,
		SMSService: smsService,
	}
}
