package xnotificaciones

import (
	"auth/graph_auth/model"
	"context"
	"fmt"
)

func EnviarNotificacion(ctx context.Context, mensaje model.XNotificacionEnvio) (bool, error) {
	cha := GetGlobal()
	cha.Mu.Lock()
	defer cha.Mu.Unlock()

	for userID, channels := range cha.Subscriptores {
		xn := &model.XNotificacion{
			Title:    mensaje.Title,
			DataJSON: mensaje.DataJSON,
		}
		for _, ch := range channels {
			select {
			case ch <- xn:
				// Successfully sent the notification
			default:
				// Channel is full, consider logging this event
				fmt.Printf("Notification channel for user %s is full.\n", userID)
			}
		}
	}

	// fmt.Printf("-> %+v\n", cha.Subscriptores)

	return true, nil
}
