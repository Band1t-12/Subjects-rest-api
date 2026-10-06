package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	DELETE FROM Subjects
	WHERE id = $1
	`
	_, err := conn.Exec(ctx, sqlQuery, id)
	return err
}
