package menu

import (
	"auth/graph_auth/model"
	"database/sql"
)

func parseRows(rows *sql.Rows, t *model.Menus) error {
	return rows.Scan(
		&t.ID,
		&t.Label,
		&t.Path,
		&t.Icon,
		&t.Color,
		&t.Grupo,
		&t.Orden,
	)
}
