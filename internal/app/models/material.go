package models

import (
	"time"

	"github.com/azka1415/crud-bootcamp-evermos/db"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/repository"
)

type Material struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Teacher   int       `json:"teacher"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MaterialService struct{}

func (m *MaterialService) GetAll(limit, page int, sort, field string) ([]Material, error) {
	db, err := db.GetDB()

	offset := (page - 1) * limit

	var materials []Material

	if err != nil {
		return nil, err
	}

	materialRepository := repository.MaterialRepository{}
	materialRepository.SetDB(db)

	rows, err := materialRepository.GetAll(field, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var material Material
		rows.Scan(&material.Id, &material.Title, &material.Teacher, &material.CreatedAt, &material.UpdatedAt)
		materials = append(materials, material)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return materials, nil
}
