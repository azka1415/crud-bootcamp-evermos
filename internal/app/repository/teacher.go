package repository

import (
	"database/sql"
	"fmt"
)

type TeacherRepository struct {
	db *sql.DB
}

func NewTeacherRepository(db *sql.DB) *TeacherRepository {
	tr := TeacherRepository{}
	tr.SetDB(db)
	return &tr
}

func (t *TeacherRepository) SetDB(db *sql.DB) {
	t.db = db
}

func (t *TeacherRepository) GetAll(field, sort string, limit, offset int) (*sql.Rows, error) {
	query := fmt.Sprintf(
		"SELECT * FROM teachers ORDER BY %s %s LIMIT %d OFFSET %d",
		field, sort, limit, offset)
	rows, err := t.db.Query(query)
	return rows, err
}
