package xnotificaciones

import (
	"auth/graph/model"
	"sync"
)

type Chan struct {
	Mu            sync.Mutex
	Subscriptores map[string]chan *model.XNotificacion
}

var global *Chan
var once sync.Once

// InitializeGlobal ensures that the global variable is initialized only once.
func InitializeGlobal() {
	once.Do(func() {
		global = &Chan{
			Subscriptores: make(map[string]chan *model.XNotificacion),
		}
	})
}

// GetGlobal returns the initialized global variable.
func GetGlobal() *Chan {
	if global == nil {
		InitializeGlobal()
	}
	return global
}
