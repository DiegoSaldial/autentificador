package usuarios

import (
	"auth/graph/model"
	"database/sql"
	"strconv"
)

func Actualizar(db *sql.DB, input model.UpdateUsuario) (*model.Usuario, error) {
	if err := permisos_obligatorios(input.Roles, input.PermisosSueltos); err != nil {
		return nil, err
	}
	us, err := GetById(db, input.ID)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := `
	update usuarios set 
	nombres=?, 
	apellido1=?, 
	apellido2=?, 
	documento=?, 
	celular=?, 
	correo=?, 
	sexo=?, 
	direccion=? 
	where id= ? 
	`
	_, err = tx.Exec(sql,
		input.Nombres,
		input.Apellido1,
		input.Apellido2,
		input.Documento,
		input.Celular,
		input.Correo,
		input.Sexo,
		input.Direccion,
		input.ID,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(input.Password) > 0 && us.OauthID == nil {
		_, err = tx.Exec("update usuarios set password=SHA2(?, 256) where id = ?", input.Password, input.ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if len(input.Username) > 0 && us.OauthID == nil {
		_, err = tx.Exec("update usuarios set username=? where id = ?", input.Username, input.ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	_, err = tx.Exec("delete from rol_usuario where usuario_id = ?", input.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec("delete from usuario_permiso where usuario_id = ?", input.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	id, _ := strconv.ParseInt(input.ID, 10, 64)

	// asignar roles
	if len(input.Roles) > 0 {
		err = asignarRoles(tx, input.Roles, id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// fi asignar roles

	// asignar permisos sueltos
	if len(input.PermisosSueltos) > 0 {
		err = asignarPermisos(tx, input.PermisosSueltos, id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// fin permisos sueltos

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return GetById(db, strconv.FormatInt(id, 10))
}

func UpdatePerfil(db *sql.DB, input model.UpdatePerfil) (*model.Usuario, error) {
	user, err := GetById(db, input.ID)
	if err != nil {
		return nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := `
	update usuarios set 
	nombres=?, 
	apellido1=?, 
	apellido2=?, 
	documento=?, 
	celular=?, 
	correo=?, 
	sexo=?, 
	direccion=? 
	where id= ? 
	`
	_, err = tx.Exec(sql,
		input.Nombres,
		input.Apellido1,
		input.Apellido2,
		input.Documento,
		input.Celular,
		input.Correo,
		input.Sexo,
		input.Direccion,
		input.ID,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if user.OauthID == nil {
		if input.Username != nil && len(*input.Username) > 0 {
			_, err = tx.Exec("update usuarios set username=? where id=?", input.Username, input.ID)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}

		if input.Password != nil && len(*input.Password) > 0 {
			_, err = tx.Exec("update usuarios set password=SHA2(?, 256) where id = ?", input.Password, input.ID)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return GetById(db, input.ID)
}

func SetLastLogin(db *sql.DB, userid string) {
	sql := "update usuarios set last_login=CURRENT_TIMESTAMP where id = ?"
	db.Exec(sql, userid)
}
