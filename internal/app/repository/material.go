package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type MaterialRepository struct {
	db *sql.DB
}

func NewMaterialRepository(db *sql.DB) MaterialRepository {
	mr := MaterialRepository{}
	mr.SetDB(db)
	return mr
}

func (m *MaterialRepository) SetDB(db *sql.DB) {
	m.db = db
}

func (m *MaterialRepository) GetAll(field, sort string, limit, offset int) (*sql.Rows, error) {
	query := fmt.Sprintf(
		"SELECT * FROM materials ORDER BY %s %s LIMIT %d OFFSET %d",
		field, sort, limit, offset)
	rows, err := m.db.Query(query, limit, offset)
	return rows, err
}

func (m *MaterialRepository) GetByID(id int) *sql.Row {
	query := fmt.Sprintf("SELECT * FROM materials WHERE id = %d", id)
	row := m.db.QueryRow(query)
	return row
}

func (m *MaterialRepository) UpdateByID(matID int, title string, teacher int) *sql.Row {
	currentTime := time.Now()
	query := `
		UPDATE materials
		SET title = ?,
			updated_at = ?,
			teacher_id = ?
		WHERE id = ?
		RETURNING *
	`
	row := m.db.QueryRow(query, title, currentTime.Local(), teacher, matID)
	return row
}

func (m *MaterialRepository) NewMaterial(title string, teacher_id int) *sql.Row {
	query := `
		INSERT INTO materials (title, teacher_id )
		VALUES (?, ?)
		RETURNING *;
	`

	row := m.db.QueryRow(query, title, teacher_id)
	return row
}

func (m *MaterialRepository) DeleteMaterial(matID int) error {
	_, err := m.db.Exec("DELETE FROM materials WHERE id=?", matID)
	return err
}

func (m *MaterialRepository) MaterialExistsByID(matID int) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM materials WHERE id = %d ", matID)
	var count int
	err := m.db.QueryRow(query).Scan(&count)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
