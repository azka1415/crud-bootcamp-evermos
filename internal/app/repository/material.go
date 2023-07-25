package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type MaterialRepositoryImpl struct {
	db *sql.DB
}

type MaterialRepository interface {
	GetAll(field, sort string, limit, offset int, teacher_id int) (*sql.Rows, error)
	GetByID(id int) *sql.Row
	UpdateByID(matID int, title string, teacher int, currTime time.Time) *sql.Row
	NewMaterial(title string, teacher_id int) *sql.Row
	DeleteMaterial(matID int) error
	MaterialExistsByID(matID int) (bool, error)
}

func NewMaterialRepository(db *sql.DB) MaterialRepository {
	mr := MaterialRepositoryImpl{db: db}
	var matRep MaterialRepository = &mr
	return matRep
}

func (m *MaterialRepositoryImpl) GetAll(field, sort string, limit, offset int, teacher_id int) (*sql.Rows, error) {
	if teacher_id == 0 {
		query := fmt.Sprintf(
			"SELECT * FROM materials ORDER BY %s %s LIMIT %d OFFSET %d",
			field, sort, limit, offset)
		rows, err := m.db.Query(query)
		return rows, err
	}
	query := fmt.Sprintf(
		"SELECT * FROM materials WHERE teacher_id = %d ORDER BY %s %s LIMIT %d OFFSET %d",
		teacher_id, field, sort, limit, offset)
	rows, err := m.db.Query(query)
	return rows, err
}

func (m *MaterialRepositoryImpl) GetByID(id int) *sql.Row {
	query := fmt.Sprintf("SELECT * FROM materials WHERE id = %d", id)
	row := m.db.QueryRow(query)
	return row
}

func (m *MaterialRepositoryImpl) UpdateByID(matID int, title string, teacher int, currTime time.Time) *sql.Row {
	query := `
		UPDATE materials
		SET title = ?,
			updated_at = ?,
			teacher_id = ?
		WHERE id = ?
		RETURNING *
	`
	row := m.db.QueryRow(query, title, currTime.UTC(), teacher, matID)
	return row
}

func (m *MaterialRepositoryImpl) NewMaterial(title string, teacher_id int) *sql.Row {
	query := `
		INSERT INTO materials (title, teacher_id)
		VALUES (?, ?)
		RETURNING *;
	`

	row := m.db.QueryRow(query, title, teacher_id)
	return row
}

func (m *MaterialRepositoryImpl) DeleteMaterial(matID int) error {
	_, err := m.db.Exec("DELETE FROM materials WHERE id=?", matID)
	return err
}

func (m *MaterialRepositoryImpl) MaterialExistsByID(matID int) (bool, error) {
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
