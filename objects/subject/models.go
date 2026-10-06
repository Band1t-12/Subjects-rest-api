package objects

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	ctx  context.Context
	conn *pgx.Conn
}
