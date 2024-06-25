package menu

import (
	"auth/graph/model"
	"database/sql"
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

func ListarAsignados(db *sql.DB, userid string) ([]*model.Menus, error) {
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
	sql := `
	SELECT DISTINCT m.id, m.label, m.path, m.icon, m.color, m.grupo
	FROM menus m
	INNER JOIN menus_usuario mu ON mu.menu_id = m.id
	LEFT JOIN rol_usuario ru ON ru.usuario_id = ?
	LEFT JOIN rol_menus rm ON rm.rol = ru.rol
	WHERE mu.usuario_id = ? OR m.id = rm.menu_id
	ORDER BY m.grupo, m.id;
	`
	rows, err := db.Query(sql, userid, userid)
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
