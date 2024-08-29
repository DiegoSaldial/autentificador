package roles

import (
	"auth/database/auth/menu"
	"auth/database/auth/permisos"
	"auth/graph/model"
	"database/sql"
	"errors"
)

func GetRolesByUsuario(db *sql.DB, userid string, show_permisos bool) ([]*model.ResponseRolMe, error) {
	sql := `
	select r.nombre, r.descripcion, r.jerarquia, r.fecha_registro,ru.usuario_id, ru.fecha_registro as fecha_asignado 
	from roles r
	left join rol_usuario ru on ru.rol = r.nombre
	where ru.usuario_id = ?;
	`
	rows, err := db.Query(sql, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*model.ResponseRolMe{}

	for rows.Next() {
		r := model.ResponseRolMe{}
		er := parse(rows, &r)
		if er != nil {
			return nil, er
		}
		if show_permisos {
			r.Permisos, er = permisos.GetPermisosByRol(db, r.Nombre)
			if er != nil {
				return nil, er
			}
		}
		roles = append(roles, &r)
	}

	return roles, nil
}

func GetRolById(db *sql.DB, rol string) (*model.Rol, error) {
	sq := "select nombre,descripcion,jerarquia,fecha_registro from roles where nombre = ?"
	row := db.QueryRow(sq, rol)
	r := model.Rol{}
	err := parseRow(row, &r)
	if err == sql.ErrNoRows {
		return nil, errors.New("rol no existente")
	}

	r.Permisos, err = permisos.GetPermisosByRol(db, r.Nombre)
	if err != nil {
		return nil, err
	}
	r.Menus, err = menu.GetMenusbyRol(db, r.Nombre)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func GetRoles(db *sql.DB) ([]*model.ResponseRoles, error) {
	sql := `
	SELECT  
		r.nombre,r.descripcion,r.jerarquia,r.fecha_registro,
		COUNT(DISTINCT rm.id) AS total_menus,
		COUNT(DISTINCT rp.metodo) AS total_permisos,
		COUNT(DISTINCT ru.usuario_id) AS total_usuarios
	FROM
		roles r
	LEFT JOIN 
		rol_menus rm ON r.nombre = rm.rol
	LEFT JOIN 
		rol_permiso rp ON r.nombre = rp.rol
	LEFT JOIN 
		rol_usuario ru ON r.nombre = ru.rol
	GROUP BY 
		r.nombre
	order by r.jerarquia asc;
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roles := []*model.ResponseRoles{}

	for rows.Next() {
		r := model.ResponseRoles{}
		er := parseRes(rows, &r)
		if er != nil {
			return nil, er
		}
		roles = append(roles, &r)
	}

	return roles, nil
}
