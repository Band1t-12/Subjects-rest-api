package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Select(ctx context.Context, conn *pgx.Conn, id int) (Subject, error) {
	var s Subject

	sqlQuery := `
	SELECT subjects.id, subjects.name, teachers.name FROM subjects
	JOIN teachers ON teachers.id = subjects.teacher_id
	WHERE subjects.id = $1
	`

	err := conn.QueryRow(ctx, sqlQuery, id).Scan(s.ID, s.Name, s.TeacherName)
	if err != nil {
		return Subject{}, err
	}

	return s, err
}

func SelectAll(ctx context.Context, conn *pgx.Conn) ([]Subject, error) {
	sqlQuery := `
	SELECT subjects.id, subjects.name, teachers.name FROM subjects
	JOIN teachers ON teachers.id = subjects.teacher_id
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subjects := make([]Subject, 0)

	for rows.Next() {
		var s Subject

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.TeacherName,
		)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, s)
	}

	return subjects, err
}
