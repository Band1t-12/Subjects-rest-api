package sql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func UpdateSubject(ctx context.Context, conn *pgx.Conn, s Subject) error {
	var exists bool
	err := conn.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM teachers WHERE name = $1)`, s.TeacherName).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		_, err = conn.Exec(ctx,
			`UPDATE subjects 
			SET name = $1, 
			teacher_id = (SELECT id FROM teachers WHERE name = $2)
			WHERE id = $3`,
			s.Name, s.TeacherName, s.ID)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Teacher Not Found!")
	}

	return err
}

func UpdateTeacher(ctx context.Context, conn *pgx.Conn, t Teacher) error {
	sqlQuery := `
	UPDATE teachers
	SET name = $1
	WHERE id = $2
	`

	_, err := conn.Exec(ctx, sqlQuery, t.Name, t.ID)
	return err
}
