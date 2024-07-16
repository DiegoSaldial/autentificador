package xnotificaciones

import (
	"auth/graph/model"
	"context"
	"fmt"
)

func EnviarNotificacion(ctx context.Context, mensaje model.XNotificacionEnvio) (bool, error) {

	// mu.Lock()

	cha := GetGlobal()
	cha.Mu.Lock()
	// defer ch.Mu.Unlock()

	for _, ch := range cha.Subscriptores {
		xn := &model.XNotificacion{
			Title:    mensaje.Title,
			DataJSON: mensaje.DataJSON,
		}
		ch <- xn
	}

	// mu.Unlock()
	cha.Mu.Unlock()

	fmt.Printf("-> %+v\n", cha.Subscriptores)

	return true, nil
}
