package roles

import (
	"auth/database/auth/permisos"
	"auth/graph_auth/model"
	"database/sql"
	"fmt"
	"strings"
)

func Actualizar(db *sql.DB, input model.NewRol) (*model.Rol, error) {
	sql := `update roles set descripcion=?,jerarquia=? where nombre=?`
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(sql, input.Descripcion, input.Jerarquia, input.Nombre)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec("delete from rol_permiso where rol = ?", input.Nombre)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	sql = "replace into rol_permiso(rol, metodo) values %s"
	places := make([]string, len(input.Permisos))
	args := make([]interface{}, len(input.Permisos)*2)

	for i, p := range input.Permisos {
		places[i] = "(?,?)"
		args[i*2] = input.Nombre
		args[i*2+1] = p
	}

	sql = fmt.Sprintf(sql, strings.Join(places, ", "))
	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// ====

	_, err = tx.Exec("delete from rol_menus where rol = ?", input.Nombre)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	sql = "replace into rol_menus(rol, menu_id) values %s"
	places = make([]string, len(input.Menus))
	args = make([]interface{}, len(input.Menus)*2)

	for i, p := range input.Menus {
		places[i] = "(?,?)"
		args[i*2] = input.Nombre
		args[i*2+1] = p
	}

	sql = fmt.Sprintf(sql, strings.Join(places, ", "))
	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// ====

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	r, err := GetRolById(db, input.Nombre)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	res := model.Rol{}
	res.Nombre = r.Nombre
	res.Descripcion = r.Descripcion
	res.Jerarquia = r.Jerarquia
	res.FechaRegistro = r.FechaRegistro
	res.Permisos, err = permisos.GetPermisosByRol(db, input.Nombre)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &res, nil
}
