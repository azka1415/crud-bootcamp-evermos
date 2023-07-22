package models

import (
	"time"

	"github.com/azka1415/crud-bootcamp-evermos/db"
)

type Material struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Teacher   int       `json:"teacher"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MaterialService struct{}

func (m *MaterialService) GetAll(limit, offset int) ([]Material, error) {
	db, dbErr := db.GetDB()
	var materials []Material
	if dbErr != nil {
		return nil, dbErr
	}
	rows, err := db.Query("SELECT * FROM materials")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var material Material
		rows.Scan(&material.Id, &material.Title, &material.Teacher, &material.CreatedAt, &material.UpdatedAt)
		materials = append(materials, material)
	}
	return materials, nil
}
