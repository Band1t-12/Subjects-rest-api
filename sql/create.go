package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXIST Subjects(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	teacher_id INTEGER NOT NULL,
	);`

	_, err := conn.Exec(ctx, sqlQuery)
	return err

}
