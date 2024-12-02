package xauth

import (
	"auth/auth/permisos"
	"auth/auth/usuarios"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func UaserIDMiddleware(db *sql.DB) transport.WebsocketInitFunc {

	return func(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
		auth, ok := initPayload["Authorization"].(string)
		pay := &transport.InitPayload{}

		// fmt.Println("auth \n", auth, ok)
		if !ok {
			return ctx, pay, nil
		}

		if auth == "" || len(auth) <= 7 {
			return ctx, pay, nil
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := JwtValidate(auth)
		if err != nil {
			return ctx, pay, nil
		}

		customClaim, _ := validate.Claims.(*JwtCustomClaim)

		data := AuthData{}
		data.Clains = customClaim
		data.TOKEN = auth
		us, er := usuarios.GetById(db, customClaim.USERID)
		if er == nil {
			data.Usuario = us
		}

		ctxx := context.WithValue(ctx, authString("auth"), &data)

		return ctxx, pay, nil
	}
}

// =======================================================
// =======================================================
// headers websockets
/*func UaserIDMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")

			fmt.Println("auth", auth)

			if auth == "" || len(auth) <= 7 {
				next.ServeHTTP(w, r)
				return
			}

			bearer := "Bearer "
			auth = auth[len(bearer):]

			validate, err := JwtValidate(auth)
			if err != nil {
				txt := err.Error()
				if !strings.HasPrefix(txt, "token is expired by") {
					next.ServeHTTP(w, r)
					return
				}
			}

			customClaim, _ := validate.Claims.(*JwtCustomClaim)

			data := AuthData{}
			data.Clains = customClaim
			data.TOKEN = auth
			us, er := usuarios.GetById(db, customClaim.USERID)
			if er == nil {
				data.Usuario = us
			}

			ctx := context.WithValue(r.Context(), authString("auth"), &data)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}*/

func CtxValueWs(ctx context.Context, db *sql.DB, metodo string) (*AuthData, error) {
	str := authString("auth")
	algo := ctx.Value(str)
	if algo == nil {
		return nil, errors.New("proporcione un token ws")
	}
	clains, _ := algo.(*AuthData)
	if clains == nil {
		return nil, errors.New("debes iniciar session")
	}
	validate, err := JwtValidate(clains.TOKEN)
	if err != nil || !validate.Valid {
		txt := err.Error()
		if !strings.HasPrefix(txt, "token is expired by") {
			return nil, errors.New(txt)
		}
	}
	if !clains.Usuario.Estado {
		return nil, errors.New("tu cuenta se encuentra suspendida")
	}

	if len(metodo) > 0 {
		err = permisos.VerificarPermiso(db, clains.Usuario.ID, metodo)
		if err != nil {
			return nil, err
		}
	}

	return clains, nil
}
