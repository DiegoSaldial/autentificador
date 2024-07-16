package usuarios

import (
	"auth/graph/model"
	"database/sql"
	"errors"
	"os"
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

func oauth_emails_permitidos(email *string) error {
	emails := os.Getenv("OAUTH_EMAILS_PERM")
	perms := strings.Split(emails, ",")

	if len(perms) == 0 {
		return nil
	}

	if email == nil {
		return errors.New("el correo no debe ser vacio")
	}

	parts := strings.Split(*email, "@")
	if len(parts) != 2 {
		return errors.New("email no válido")
	}
	domain := parts[1]

	for _, perm := range perms {
		if strings.TrimSpace(perm) == domain {
			return nil
		}
	}

	return errors.New("utilice su correo de estos dominios: " + emails)
}
