package graph_auth

import "database/sql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *sql.DB

	/* Mu            sync.Mutex
	Subscriptores map[string]chan *model.XNotificacion */
}
