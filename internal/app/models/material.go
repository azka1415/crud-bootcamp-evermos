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
	Teacher_id int    `json:"teacher_id"`
}

type MaterialServiceImpl struct {
	repo repository.MaterialRepository
}

type MaterialService interface {
	GetAll(limit, page int, sort, field string, teacher_id int) ([]Material, error)
	GetByID(matID int) (Material, error)
	UpdateMaterial(matID int, updatedMaterial PayloadMaterial) (Material, error)
	NewMaterial(newMat PayloadMaterial) Material
	DeleteMaterial(matID int) (bool, error)
	MaterialExistsByID(id int) (bool, error)
}

func NewMaterialService(repo repository.MaterialRepository) MaterialService {
	ms := MaterialServiceImpl{repo: repo}
	var matService MaterialService = &ms

	return matService
}

func (m *MaterialServiceImpl) GetAll(limit, page int, sort, field string, teacher_id int) ([]Material, error) {
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

func (m *MaterialServiceImpl) GetByID(matID int) (Material, error) {
	var material Material
	row := m.repo.GetByID(matID)
	row.Scan(&material.Id, &material.Title, &material.Teacher_id, &material.CreatedAt, &material.UpdatedAt)
	return material, nil
}

func (m *MaterialServiceImpl) UpdateMaterial(matID int, updatedMaterial PayloadMaterial) (Material, error) {
	var material Material
	row := m.repo.UpdateByID(matID, updatedMaterial.Title, updatedMaterial.Teacher_id, time.Now())
	row.Scan(&material.Id, &material.Title, &material.Teacher_id, &material.CreatedAt, &material.UpdatedAt)
	return material, nil
}

func (m *MaterialServiceImpl) NewMaterial(newMat PayloadMaterial) Material {
	var mat Material
	row := m.repo.NewMaterial(newMat.Title, newMat.Teacher_id)
	row.Scan(&mat.Id, &mat.Title, &mat.Teacher_id, &mat.CreatedAt, &mat.UpdatedAt)
	return mat
}

func (m *MaterialServiceImpl) DeleteMaterial(matID int) (bool, error) {
	err := m.repo.DeleteMaterial(matID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *MaterialServiceImpl) MaterialExistsByID(id int) (bool, error) {
	valid, err := m.repo.MaterialExistsByID(id)
	if err != nil {
		return valid, errors.New("Material not found")
	}
	return valid, nil
}
