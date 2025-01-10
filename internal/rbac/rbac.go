package rbac

import (
	"sync"

	"github.com/astrokiran/nimbus/internal/common/database"
	models "github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
)

type RBACI interface {
	GetUserRoles(userID string) ([]models.Role, error)
	GetRolePermissions(roleID int) ([]models.Permission, error)
	CheckPermission(userID, permission string) (bool, error)
	GetUserPermissions(userID string) ([]models.Permission, error)
	GetUserRolesAndPermissions(userID string) (map[models.Role][]models.Permission, error)
	GetRolePermissionsByRole(role string) ([]models.Permission, error)
	GetRolePermissionsByUser(userID string) ([]models.Permission, error)
	GetActionsByPermission(permission string) ([]models.Action, error)
	GetActionsByRole(role string) ([]models.Action, error)
	GetActionsByUser(userID string) ([]string, error)
}

type RBAC struct {
	db   *database.Database
	repo *RBACRepository
}

var (
	instance *RBAC
	once     sync.Once
	ready    bool
)

func Init(db *database.Database) {
	once.Do(func() {
		instance = &RBAC{
			db:   db,
			repo: NewRBACRepository(db),
		}
		ready = true
	})
}

func GetInstance() *RBAC {
	if !ready {
		panic("RBAC not initialized")
	}
	return instance
}

func (r *RBAC) GetUserRoles(userID string) ([]models.Role, error) {
	return r.repo.GetUserRoles(userID)
}
