package models

import (
	"errors"
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

type UpdateMaterial struct {
	Title   string `json:"title"`
	Teacher string `json:"teacher"`
}

type MaterialService struct{}

func NewMaterialService() MaterialService {
	return MaterialService{}
}

func (m *MaterialService) GetAll(limit, page int, sort, field string) ([]Material, error) {
	db, err := db.GetDB()

	offset := (page - 1) * limit

	var materials []Material

	if err != nil {
		return nil, err
	}

	materialRepository := repository.NewMaterialRepository(db)

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

func (m *MaterialService) GetByID(matID int) (Material, error) {
	db, err := db.GetDB()
	var material Material
	if err != nil {
		return material, err
	}

	materialRepository := repository.NewMaterialRepository(db)
	row := materialRepository.GetByID(matID)
	row.Scan(&material.Id, &material.Title, &material.Teacher, &material.CreatedAt, &material.UpdatedAt)

	return material, nil
}

func (m *MaterialService) Update(matID int) (Material, error) {
	return Material{}, nil
}

func (m *MaterialService) MaterialExistsByID(id int) (bool, error) {
	db, err := db.GetDB()
	if err != nil {
		return false, err
	}
	materialRepo := repository.NewMaterialRepository(db)
	valid, err := materialRepo.MaterialExistsByID(id)

	if err != nil {
		return valid, errors.New("Material not found")
	}
	return valid, nil
}
