package xnotificaciones

import (
	"auth/graph/model"
	"context"
	"sync"
)

func NotificacionesSubs(ctx context.Context, userid string, mu *sync.Mutex, subs map[string]chan *model.XNotificacion) (<-chan *model.XNotificacion, error) {
	ch := make(chan *model.XNotificacion, 1)
	mu.Lock()

	subs[userid] = ch

	mu.Unlock()

	go func() {
		<-ctx.Done()
		mu.Lock()
		delete(subs, userid)
		mu.Unlock()
	}()

	return ch, nil

}
