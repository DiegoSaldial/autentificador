package menu

import (
	"auth/graph/model"
	"database/sql"
	"fmt"
)

func Listar(db *sql.DB) ([]*model.Menus, error) {
	sql := `select id,label,path,icon,color,grupo from menus order by grupo,id`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mns := []*model.Menus{}
	for rows.Next() {
		m := model.Menus{}
		er := parseRows(rows, &m)
		if er != nil {
			return nil, er
		}
		mns = append(mns, &m)
	}
	return mns, nil
}

func ListarAsignados(db *sql.DB, userid string, only_user bool) ([]*model.Menus, error) {
	/* sql := `
	select m.id,m.label,m.path,m.icon,m.color,m.grupo
		from menus m
		inner join menus_usuario mu on mu.menu_id = m.id
		left join rol_usuario ru on ru.usuario_id = mu.usuario_id
		left join rol_menus rm on rm.rol = ru.rol
		where mu.usuario_id = ?
		group by m.id
		order by m.grupo,m.id
	` */

	xsql := `
	SELECT DISTINCT m.id, m.label, m.path, m.icon, m.color, m.grupo
	FROM menus m
	LEFT JOIN menus_usuario mu ON mu.menu_id = m.id AND mu.usuario_id = ?
	LEFT JOIN rol_menus rm ON rm.menu_id = m.id
	LEFT JOIN roles r ON r.nombre = rm.rol
	LEFT JOIN rol_usuario ru ON ru.rol = r.nombre AND ru.usuario_id = ?
	WHERE (mu.usuario_id IS NOT NULL OR ru.usuario_id IS NOT NULL) %s
	ORDER BY m.grupo, m.id
	`

	from_user := ""
	if only_user {
		from_user = "and mu.id is not null"
	}

	xsql = fmt.Sprintf(xsql, from_user)

	rows, err := db.Query(xsql, userid, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mns := []*model.Menus{}
	for rows.Next() {
		m := model.Menus{}
		er := parseRows(rows, &m)
		if er != nil {
			return nil, er
		}
		mns = append(mns, &m)
	}
	return mns, nil
}

func GetMenusbyRol(db *sql.DB, rol string) ([]*model.Menus, error) {
	sql := `
	select m.id,m.label,m.path,m.icon,m.color,m.grupo 
	from menus m
	inner join rol_menus rm on rm.menu_id = m.id
	where rm.rol = ?
	`
	rows, err := db.Query(sql, rol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mns := []*model.Menus{}
	for rows.Next() {
		m := model.Menus{}
		er := parseRows(rows, &m)
		if er != nil {
			return nil, er
		}
		mns = append(mns, &m)
	}
	return mns, nil
}
