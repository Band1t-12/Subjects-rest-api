package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertSubject(ctx context.Context, conn *pgx.Conn, s Subject) error {
	var exists bool
	var teacherID int
	err := conn.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM teachers WHERE name = $1)`, s.TeacherName).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		err = conn.QueryRow(ctx, `SELECT id FROM teachers WHERE name = $1`, s.TeacherName).Scan(&teacherID)
	} else {
		err = conn.QueryRow(ctx, `INSERT INTO teachers(name) VALUES($1) RETURNING id`, s.TeacherName).Scan(&teacherID)
	}
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `INSERT INTO subjects(name, teacher_id) VALUES($1, $2)`, s.Name, teacherID)
	return err
}

func InsertTeacher(ctx context.Context, conn *pgx.Conn, t Teacher) error {
	sqlQuery := `
	INSERT INTO teachers(name)
	VALUES($1)
	`

	_, err := conn.Exec(ctx, sqlQuery, t.Name)
	return err

}
