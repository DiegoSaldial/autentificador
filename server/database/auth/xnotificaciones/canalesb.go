package xnotificaciones

import (
	"auth/graph_auth/model"
	"sync"
)

type Chan struct {
	Mu            sync.Mutex
	Subscriptores map[string][]chan *model.XNotificacion
}

var global *Chan
var once sync.Once

// InitializeGlobal ensures that the global variable is initialized only once.
func InitializeGlobal() {
	once.Do(func() {
		global = &Chan{
			Subscriptores: make(map[string][]chan *model.XNotificacion),
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

// AddSubscriber adds a new channel for a given user ID.
func (c *Chan) AddSubscriber(userID string, ch chan *model.XNotificacion) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Subscriptores[userID] = append(c.Subscriptores[userID], ch)
}

// RemoveSubscriber removes a channel for a given user ID.
func (c *Chan) RemoveSubscriber(userID string, ch chan *model.XNotificacion) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	channels := c.Subscriptores[userID]
	for i, subscriber := range channels {
		if subscriber == ch {
			c.Subscriptores[userID] = append(channels[:i], channels[i+1:]...)
			break
		}
	}
	// Remove the userID key if there are no more channels
	if len(c.Subscriptores[userID]) == 0 {
		delete(c.Subscriptores, userID)
	}
}

// Broadcast sends a notification to all channels for a given user ID.
func (c *Chan) Broadcast(userID string, notification *model.XNotificacion) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for _, ch := range c.Subscriptores[userID] {
		ch <- notification
	}
}
