package repository

import (
	"database/sql"
	"fmt"
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

func (m *MaterialRepository) UpdateByID(id int) (*sql.Row, error) {
	return nil, nil
}

func (m *MaterialRepository) MaterialExistsByID(id int) (bool, error) {
	query := fmt.Sprintf("SELECT * FROM materials WHERE id = %d", id)
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
