package models

import (
	"errors"
	"time"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/repository"
)

type Material struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Teacher_id int       `json:"teacher_id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type PayloadMaterial struct {
	Title      string `json:"title"`
	Teacher_id int    `json:"teacher"`
}

type MaterialService struct {
	repo *repository.MaterialRepository
}

func NewMaterialService(repo *repository.MaterialRepository) *MaterialService {
	ms := MaterialService{}
	ms.repo = repo
	return &ms
}

func (m *MaterialService) GetAll(limit, page int, sort, field string, teacher_id int) ([]Material, error) {
	offset := (page - 1) * limit
	var materials []Material

	rows, err := m.repo.GetAll(field, sort, limit, offset, teacher_id)
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
	var material Material
	row := m.repo.GetByID(matID)
	row.Scan(&material.Id, &material.Title, &material.Teacher_id, &material.CreatedAt, &material.UpdatedAt)
	return material, nil
}

func (m *MaterialService) UpdateMaterial(matID int, updatedMaterial PayloadMaterial) (Material, error) {
	var material Material
	row := m.repo.UpdateByID(matID, updatedMaterial.Title, updatedMaterial.Teacher_id)
	row.Scan(&material.Id, &material.Title, &material.Teacher_id, &material.CreatedAt, &material.UpdatedAt)
	return material, nil
}

func (m *MaterialService) NewMaterial(newMat PayloadMaterial) Material {
	var mat Material
	row := m.repo.NewMaterial(newMat.Title, newMat.Teacher_id)
	row.Scan(&mat.Id, &mat.Title, &mat.Teacher_id, &mat.CreatedAt, &mat.UpdatedAt)
	return mat
}

func (m *MaterialService) DeleteMaterial(matID int) (bool, error) {
	err := m.repo.DeleteMaterial(matID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *MaterialService) MaterialExistsByID(id int) (bool, error) {
	valid, err := m.repo.MaterialExistsByID(id)
	if err != nil {
		return valid, errors.New("Material not found")
	}
	return valid, nil
}
