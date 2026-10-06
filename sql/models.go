package sql

type Subject struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TeacherName string `json:"teacher_name"`
}
