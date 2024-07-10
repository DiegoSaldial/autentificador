package xnotificaciones

import (
	"auth/graph/model"
	"context"
	"fmt"
	"sync"
)

func EnviarNotificacion(ctx context.Context, mu *sync.Mutex, subs map[string]chan *model.XNotificacion, mensaje model.XNotificacionEnvio) (bool, error) {

	mu.Lock()

	for _, ch := range subs {
		xn := model.XNotificacion{
			Title:    mensaje.Title,
			DataJSON: mensaje.DataJSON,
		}
		ch <- &xn
	}

	mu.Unlock()

	fmt.Printf("%+v\n\n", subs)

	return true, nil
}
