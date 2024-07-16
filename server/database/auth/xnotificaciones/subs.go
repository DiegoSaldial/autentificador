package xnotificaciones

import (
	"auth/graph/model"
	"context"
)

func NotificacionesSubs(ctx context.Context, userid string) (<-chan *model.XNotificacion, error) {

	cha := GetGlobal()
	cha.Mu.Lock()

	ch := make(chan *model.XNotificacion, 1)
	// mu.Lock()

	// subs[userid] = ch
	cha.Subscriptores[userid] = ch

	// mu.Unlock()
	cha.Mu.Unlock()

	go func() {
		<-ctx.Done()
		cha.Mu.Lock()
		delete(cha.Subscriptores, userid)
		cha.Mu.Unlock()
	}()

	return ch, nil

}
