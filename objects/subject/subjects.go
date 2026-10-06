package objects

import (
	"StudyApi/sql"
	"context"

	"github.com/jackc/pgx/v5"
)

func (d *Database) AddSubject(name string, teacher_name string) error {
	t := sql.Subject{
		ID:          0,
		Name:        name,
		TeacherName: teacher_name,
	}

	err := sql.Insert(d.ctx, d.conn, t)
	return err
}

func (d *Database) GetSubject(id int) (sql.Subject, error) {
	subject, err := sql.Select(d.ctx, d.conn, id)

	return subject, err
}

func (d *Database) GetAllSubjects() ([]sql.Subject, error) {
	subjects, err := sql.SelectAll(d.ctx, d.conn)
	return subjects, err
}

func (d *Database) UpdateSubject(id int, name string, teacherName string) error {
	t := sql.Subject{
		ID:          id,
		Name:        name,
		TeacherName: teacherName,
	}

	err := sql.Update(d.ctx, d.conn, t)
	return err
}

func (d *Database) DeleteSubject(id int) error {
	err := sql.Delete(d.ctx, d.conn, id)
	return err
}

func NewDatabase(ctx context.Context, conn *pgx.Conn) *Database {
	db := Database{
		ctx:  ctx,
		conn: conn,
	}
	return &db
}
