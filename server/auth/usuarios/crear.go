package usuarios

import (
	"auth/auth/utils"
	"auth/graph_auth/model"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Crear(db *sql.DB, input model.NewUsuario, oauth_id *string) (*model.Usuario, error) {
	if err := permisos_obligatorios(input.Roles, input.PermisosSueltos); err != nil {
		return nil, err
	}
	if err := validarCadena(input.Username, "username"); err != nil {
		return nil, err
	}
	if err := validarCadena(input.Password, "password"); err != nil {
		return nil, err
	}
	point, err := ubicacion(input.Latitud, input.Longitud)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := `
	INSERT INTO usuarios(nombres, apellido1, apellido2, documento, celular, correo, sexo, direccion, username, password,oauth_id,ubicacion)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, SHA2(?, 256),?, ST_GeomFromText(?));
	`
	res, err := tx.Exec(sql,
		input.Nombres,
		input.Apellido1,
		input.Apellido2,
		input.Documento,
		input.Celular,
		input.Correo,
		input.Sexo,
		input.Direccion,
		input.Username,
		input.Password,
		oauth_id,
		point,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, _ := res.LastInsertId()
	xid := strconv.FormatInt(id, 10)

	// asignar roles
	if len(input.Roles) > 0 {
		err = asignarRoles(tx, input.Roles, id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// fin asignar roles

	// asignar permisos sueltos
	if len(input.PermisosSueltos) > 0 {
		err = asignarPermisos(tx, input.PermisosSueltos, id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// fin permisos sueltos

	// subir foto perfil
	if input.Foto64 != nil {
		foto_url, err := utils.SubirImagen(*input.Foto64, "perfil", xid)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		sql = `update usuarios set foto_url=? where id=?`
		_, err = tx.Exec(sql, foto_url, id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// fin subir foto perfil

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return GetById(db, strconv.FormatInt(id, 10))
}

func CrearOauth(db *sql.DB, input model.NewUsuarioOauth, isportal bool) (*model.Usuario, error) {

	if err := oauth_emails_permitidos(input.Correo); err != nil {
		return nil, err
	}

	var id, oauth *string
	xsql := "select id, oauth_id from usuarios where oauth_id=?"
	err := db.QueryRow(xsql, input.Username).Scan(&id, &oauth)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if id != nil {
		return GetById(db, *id)
	}

	rol := os.Getenv("DEFAULT_ROL_OAUTH")
	nombres, aps := splitName(input.Nombres)

	data := model.NewUsuario{}
	data.Nombres = cut_string(nombres, 30)
	data.Apellido1 = cut_string(aps, 30)
	data.Celular = input.Celular
	data.Correo = input.Correo
	data.Username = input.Username
	data.Password = input.Password
	data.Roles = []string{rol}

	if isportal {
		id_existe := ""
		xsql := "select id  from usuarios where username=?"
		err := db.QueryRow(xsql, input.Username).Scan(&id_existe)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		if id_existe != "" {
			sqlx := "update usuarios set password=SHA2(?, 256) where id=?"
			_, err = db.Exec(sqlx, input.Password, id_existe)
			if err != nil {
				return nil, err
			}
			return GetById(db, id_existe)
		}
	}

	return Crear(db, data, &data.Username)
}

func asignarRoles(tx *sql.Tx, roles []string, userid int64) error {
	user_rols := "replace into `rol_usuario`(`rol`,`usuario_id`) values %s"
	places := make([]string, len(roles))
	args := make([]interface{}, len(roles)*2)

	for i, r := range roles {
		places[i] = "(?,?)"
		args[i*2] = r
		args[i*2+1] = userid
	}

	user_rols = fmt.Sprintf(user_rols, strings.Join(places, ", "))
	_, err := tx.Exec(user_rols, args...)
	return err
}

func asignarPermisos(tx *sql.Tx, permisosSueltos []string, userid int64) error {
	user_perms := "replace into `usuario_permiso`(`usuario_id`,`metodo`) values %s"
	places2 := make([]string, len(permisosSueltos))
	args2 := make([]interface{}, len(permisosSueltos)*2)

	for i, p := range permisosSueltos {
		places2[i] = "(?,?)"
		args2[i*2] = userid
		args2[i*2+1] = p
	}

	user_perms = fmt.Sprintf(user_perms, strings.Join(places2, ", "))
	_, err := tx.Exec(user_perms, args2...)
	return err
}
