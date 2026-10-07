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

func SelectAllSubjects(ctx context.Context, conn *pgx.Conn) ([]Subject, error) {
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

func SelectAllTeachers(ctx context.Context, conn *pgx.Conn) ([]Teacher, error) {
	sqlQuery := `
	SELECT id, name FROM teachers
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teachers := make([]Teacher, 0)

	for rows.Next() {
		var t Teacher

		err := rows.Scan(
			&t.ID,
			&t.Name,
		)
		if err != nil {
			return nil, err
		}

		teachers = append(teachers, t)
	}

	return teachers, err
}
