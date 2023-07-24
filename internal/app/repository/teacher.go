package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/azka1415/crud-bootcamp-evermos/db"
)

type TeacherRepository struct {
	db *sql.DB
}

func NewTeacherRepository() *TeacherRepository {
	tr := TeacherRepository{}
	db, _ := db.GetDB()
	tr.db = db
	return &tr
}

func (t *TeacherRepository) GetAll(field, sort string, limit, offset int) (*sql.Rows, error) {
	query := fmt.Sprintf(
		"SELECT * FROM teachers ORDER BY %s %s LIMIT %d OFFSET %d",
		field, sort, limit, offset)
	rows, err := t.db.Query(query)
	return rows, err
}

func (t *TeacherRepository) GetByID(id int) *sql.Row {
	query := fmt.Sprintf("SELECT * FROM teachers WHERE id = %d", id)
	row := t.db.QueryRow(query)
	return row
}

func (t *TeacherRepository) UpdateByID(id int, name, position string) *sql.Row {
	currentTime := time.Now()
	query := `
		UPDATE teachers
		SET name = ?,
			updated_at = ?,
			position = ?
		WHERE id = ?
		RETURNING *
	`
	row := t.db.QueryRow(query, name, currentTime.Local(), position, id)
	return row
}

func (t *TeacherRepository) NewTeacher(name, position string) *sql.Row {
	query := `
		INSERT INTO teachers (name, position)
		VALUES (?, ?)
		RETURNING *;
	`

	row := t.db.QueryRow(query, name, position)
	return row
}

func (t *TeacherRepository) DeleteTeacher(id int) error {
	_, err := t.db.Exec("DELETE FROM teachers WHERE id=?", id)
	return err
}

func (t *TeacherRepository) TeacherExistsByID(id int) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM teachers WHERE id = %d ", id)
	var count int
	err := t.db.QueryRow(query).Scan(&count)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
