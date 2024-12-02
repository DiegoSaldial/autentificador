package roles

import (
	"auth/graph_auth/model"
	"database/sql"
)

func parse(rows *sql.Rows, t *model.ResponseRolMe) error {
	return rows.Scan(
		&t.Nombre,
		&t.Descripcion,
		&t.Jerarquia,
		&t.FechaRegistro,
		&t.UsuarioID,
		&t.FechaAsignado,
	)
}

func parseRes(rows *sql.Rows, t *model.ResponseRoles) error {
	return rows.Scan(
		&t.Nombre,
		&t.Descripcion,
		&t.Jerarquia,
		&t.FechaRegistro,
		&t.Menus,
		&t.Permisos,
		&t.Usuarios,
	)
}

func parseRow(row *sql.Row, t *model.Rol) error {
	return row.Scan(
		&t.Nombre,
		&t.Descripcion,
		&t.Jerarquia,
		&t.FechaRegistro,
	)
}

/*
func parseRows(rows *sql.Rows, t *model.Rol) error {
	return rows.Scan(
		&t.Nombre,
		&t.Descripcion,
		&t.Jerarquia,
		&t.FechaRegistro,
	)
}
*/
