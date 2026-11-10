package objects

import "StudyApi/sql"

func (d *Database) AddTeacher(name string) error {
	t := sql.Teacher{
		ID:   0,
		Name: name,
	}

	return sql.InsertTeacher(d.ctx, d.conn, t)
}

func (d *Database) UpdateTacher(id int, name string) error {
	t := sql.Teacher{
		ID:   id,
		Name: name,
	}

	return sql.UpdateTeacher(d.ctx, d.conn, t)
}

func (d *Database) GetTeacher(id int) (sql.Teacher, error) {
	teacher, err := sql.SelectTeacher(d.ctx, d.conn, id)

	return teacher, err
}

func (d *Database) GetAllTeachers() ([]sql.Teacher, error) {
	arr, err := sql.SelectAllTeachers(d.ctx, d.conn)

	return arr, err
}
