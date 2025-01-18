package menu

import (
	"auth/graph_auth/model"
	"database/sql"
)

func AsignarMenusUsuario(db *sql.DB, input model.AsignarMenusUsuario) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}

	_, err = tx.Exec("delete from menus_usuario where usuario_id=?", input.UserID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if len(input.Menus) > 0 {
		values := []interface{}{}
		query := `insert into menus_usuario(usuario_id,menu_id) values `

		for _, m := range input.Menus {
			query += "(?,?),"
			values = append(values, input.UserID, m)
		}

		query = query[:len(query)-1]

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return "ok", nil
}
