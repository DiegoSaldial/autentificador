package xlogin

import (
	"auth/database/auth/usuarios"
	"auth/database/auth/xauth"
	"auth/graph/model"
	"context"
	"database/sql"
	"fmt"
)

func CreateExternal(db *sql.DB, data model.NewLogin) (*model.ResponseLogin, error) {
	tok, err := LoginPortal(data.Username, data.Password)
	if err != nil {
		return nil, err
	}

	me, err := GetMe(tok)
	if err != nil {
		return nil, err
	}
	dp := me.Data.Me.DatosPersonales
	input := model.NewUsuarioOauth{
		Nombres:  fmt.Sprintf("%s %s", dp.Nombres, dp.PriApellido),
		Username: data.Username,
		Password: data.Password,
	}
	us, err := usuarios.CrearOauth(db, input)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	token, err := xauth.GenerateToken(ctx, us.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := xauth.GenerateRefreshToken(ctx, us.ID)
	if err != nil {
		return nil, err
	}

	res := model.ResponseLogin{
		Token:        token,
		RefreshToken: refreshToken,
	}

	usuarios.SetLastLogin(db, us.ID)

	return &res, nil
}
