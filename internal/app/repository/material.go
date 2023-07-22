package repository

import (
	"database/sql"
	"fmt"
)

type MaterialRepository struct {
	db *sql.DB
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

func (m *MaterialRepository) GetByID(id int) (*sql.Row, error) {
	return nil, nil
}
