package database

import "fmt"

const (
	UserTable       = "users"
	RoleTable       = "roles"
	PermissionTable = "permissions"

	UserRoleTable       = "user_role_relationships"
	RolePermissionTable = "role_permission_relationships"
)

func LoginCheck(name, pass string) (*User, error) {
	query := fmt.Sprintf(`select id from %s where name=? and password=?`, UserTable)
	row := db.QueryRow(query, name, pass)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	u := &User{
		ID:   id,
		Name: name,
	}
	return u, nil
}

type User struct {
	ID   int
	Name string
}

func (u *User) Roles() ([]*Role, error) {
	query := fmt.Sprintf(`select r.id, r.name
from %s ur left join %s r on ur.role_id=r.id
where ur.user_id=?`, UserRoleTable, RoleTable)

	row, err := db.Query(query, u.ID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var roles []*Role
	for row.Next() {
		r := &Role{}
		if err := row.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return roles, nil
}

type Role struct {
	ID   int
	Name string
}
