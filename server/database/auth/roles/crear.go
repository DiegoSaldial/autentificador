package roles

import (
	"auth/database/auth/permisos"
	"auth/graph_auth/model"
	"database/sql"
	"fmt"
	"strings"
)

func Crear(db *sql.DB, input model.NewRol) (*model.Rol, error) {
	sql := `insert into roles(nombre,descripcion,jerarquia) values (?,?,?)`
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(sql, input.Nombre, input.Descripcion, input.Jerarquia)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	sql = "insert into rol_permiso(rol, metodo) values %s"
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

	// ================================================
	sql = "insert into rol_menus(rol, menu_id) values %s"
	places = make([]string, len(input.Menus))
	args = make([]interface{}, len(input.Menus)*2)

	for i, m := range input.Menus {
		places[i] = "(?,?)"
		args[i*2] = input.Nombre
		args[i*2+1] = m
	}

	sql = fmt.Sprintf(sql, strings.Join(places, ", "))
	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// =================================================

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
