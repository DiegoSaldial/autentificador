package xnotificaciones

import (
	"auth/graph/model"
	"context"
)

func NotificacionesSubs(ctx context.Context, userid string) (<-chan *model.XNotificacion, error) {

	cha := GetGlobal()
	cha.Mu.Lock()

	ch := make(chan *model.XNotificacion, 1)
	cha.Subscriptores[userid] = append(cha.Subscriptores[userid], ch)

	cha.Mu.Unlock()

	go func() {
		<-ctx.Done()
		cha.Mu.Lock()
		defer cha.Mu.Unlock()
		channels := cha.Subscriptores[userid]
		for i, subscriber := range channels {
			if subscriber == ch {
				cha.Subscriptores[userid] = append(channels[:i], channels[i+1:]...)
				break
			}
		}
		// Remove the userID key if there are no more channels
		if len(cha.Subscriptores[userid]) == 0 {
			delete(cha.Subscriptores, userid)
		}
		close(ch)
	}()

	return ch, nil
}
