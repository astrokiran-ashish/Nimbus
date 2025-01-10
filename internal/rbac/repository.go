package rbac

import (
	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	models "github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	tables "github.com/astrokiran/nimbus/internal/models/nimbus/public/table"
)

type RBACRepository struct {
	db *database.Database
}

func NewRBACRepository(db *database.Database) *RBACRepository {
	return &RBACRepository{
		db: db,
	}
}

func (r *RBACRepository) GetUserRoles(userID string) ([]model.Role, error) {
	stmt := tables.UserRoles.SELECT(
		tables.Role.RoleName,
		tables.Role.RoleID,
		tables.Role.RoleDescription,
	).FROM(
		tables.Role.
			INNER_JOIN(tables.UserRoles, tables.UserRoles.RoleID.EQ(tables.Role.RoleID)),
	)

	var roles []models.Role
	err := stmt.Query(r.db.Conn, &roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *RBACRepository) GetRolePermissions(roleID int64) ([]model.Permission, error) {
	stmt := tables.Permission.SELECT(
		tables.Permission.PermissionName,
		tables.Permission.PermissionID,
	).FROM(
		tables.Permission.
			INNER_JOIN(tables.Role, tables.Role.RoleID.EQ(tables.Permission.PermissionID)),
	).WHERE(tables.Role.RoleID.EQ(r.db.Dialect.Int(roleID)))

	var permissions []models.Permission
	err := stmt.Query(r.db.Conn, &permissions)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *RBACRepository) CheckPermission(userID, permission string) (bool, error) {
	stmt := tables.Permission.SELECT(
		tables.Permission.PermissionID,
	).FROM(
		tables.Permission.
			INNER_JOIN(tables.Role, tables.Role.RoleID.EQ(tables.Permission.PermissionID)).
			INNER_JOIN(tables.UserRoles, tables.UserRoles.UserID.EQ(r.db.Dialect.String(userID))),
	).WHERE(tables.Permission.PermissionName.EQ(r.db.Dialect.String(permission)))

	var permissionID int
	err := stmt.Query(r.db.Conn, &permissionID)
	if err != nil {
		return false, err
	}

	return permissionID > 0, nil
}
