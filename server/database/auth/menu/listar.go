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
	sql := `
	select m.id,m.label,m.path,m.icon,m.color,m.grupo 
		from menus m 
		inner join menus_usuario mu on mu.menu_id = m.id
		where mu.usuario_id = ?
		order by m.grupo,m.id
	`
	rows, err := db.Query(sql, userid)
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
