package sql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, connString)
}
