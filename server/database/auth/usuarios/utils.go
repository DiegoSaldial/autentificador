package usuarios

import (
	"auth/graph/model"
	"database/sql"
	"errors"
	"strings"
)

func parseRow(row *sql.Row, t *model.Usuario) error {
	return row.Scan(
		&t.ID,
		&t.Nombres,
		&t.Apellido1,
		&t.Apellido2,
		&t.Documento,
		&t.Celular,
		&t.Correo,
		&t.Sexo,
		&t.Direccion,
		&t.Estado,
		&t.Username,
		&t.LastLogin,
		&t.OauthID,
		&t.FechaRegistro,
		&t.FechaUpdate,
	)
}

func parseRows(rows *sql.Rows, t *model.Usuario) error {
	return rows.Scan(
		&t.ID,
		&t.Nombres,
		&t.Apellido1,
		&t.Apellido2,
		&t.Documento,
		&t.Celular,
		&t.Correo,
		&t.Sexo,
		&t.Direccion,
		&t.Estado,
		&t.Username,
		&t.LastLogin,
		&t.OauthID,
		&t.FechaRegistro,
		&t.FechaUpdate,
	)
}

func permisos_obligatorios(roles, permisosueltos []string) error {
	if len(roles) == 0 && len(permisosueltos) == 0 {
		return errors.New("selecciona al menos un rol o un permiso")
	}
	return nil
}

func splitName(name string) (string, string) {
	words := strings.Split(name, " ")
	if len(words) <= 2 {
		return name, ""
	}

	firstPart := strings.Join(words[:2], " ")
	secondPart := strings.Join(words[2:], " ")
	return firstPart, secondPart
}

func cut_string(name string, max int) string {
	if len(name) > 30 {
		return name[:max]
	}
	return name
}
