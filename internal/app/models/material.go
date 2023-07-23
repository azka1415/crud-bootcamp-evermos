package models

import (
	"errors"
	"time"

	"github.com/azka1415/crud-bootcamp-evermos/db"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/repository"
)

type Material struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Teacher_id int       `json:"teacher"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UpdateMaterial struct {
	Title      string `json:"title"`
	Teacher_id int    `json:"teacher"`
}

type MaterialService struct{}

func NewMaterialService() *MaterialService {
	return &MaterialService{}
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
		rows.Scan(&material.Id, &material.Title, &material.Teacher_id, &material.CreatedAt, &material.UpdatedAt)
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
	row.Scan(&material.Id, &material.Title, &material.Teacher_id, &material.CreatedAt, &material.UpdatedAt)

	return material, nil
}

func (m *MaterialService) UpdateMaterial(matID int, updatedMaterial UpdateMaterial) (Material, error) {
	db, err := db.GetDB()
	if err != nil {
		return Material{}, err
	}

	matRepo := repository.NewMaterialRepository(db)
	var material Material

	row := matRepo.UpdateByID(matID, updatedMaterial.Title, updatedMaterial.Teacher_id)
	row.Scan(&material.Id, &material.Title, &material.Teacher_id, &material.CreatedAt, &material.UpdatedAt)

	return material, nil
}

func (m *MaterialService) NewMaterial(newMat UpdateMaterial) (Material, error) {
	db, err := db.GetDB()

	matRepo := repository.NewMaterialRepository(db)
	var mat Material
	row := matRepo.NewMaterial(newMat.Title, newMat.Teacher_id)
	row.Scan(&mat.Id, &mat.Title, &mat.Teacher_id, &mat.CreatedAt, &mat.UpdatedAt)
	return mat, err
}

func (m *MaterialService) DeleteMaterial(matID int) (bool, error) {
	db, err := db.GetDB()
	if err != nil {
		return false, err
	}
	matRepo := repository.NewMaterialRepository(db)

	err = matRepo.DeleteMaterial(matID)
	if err != nil {
		return false, err
	}
	return true, nil
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
