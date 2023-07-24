package models

import (
	"errors"
	"time"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/repository"
)

type Teacher struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Position  string    `json:"position"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PayloadTeacher struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

type TeacherService struct {
	repo *repository.TeacherRepository
}

func NewTeacherService(repo *repository.TeacherRepository) *TeacherService {
	ts := TeacherService{}
	ts.repo = repo
	return &ts
}

func (t *TeacherService) GetAll(limit, page int, sort, field string) ([]Teacher, error) {
	offset := (page - 1) * limit
	var teachers []Teacher
	rows, err := t.repo.GetAll(field, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var teacher Teacher
		rows.Scan(&teacher.Id, &teacher.Name, &teacher.Position, &teacher.CreatedAt, &teacher.UpdatedAt)
		teachers = append(teachers, teacher)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (t *TeacherService) GetByID(id int) (Teacher, error) {
	var teacher Teacher
	row := t.repo.GetByID(id)
	row.Scan(&teacher.Id, &teacher.Name, &teacher.Position, &teacher.CreatedAt, &teacher.UpdatedAt)

	return teacher, nil
}

func (t *TeacherService) UpdateTeacher(id int, updates PayloadTeacher) (Teacher, error) {
	var material Teacher

	row := t.repo.UpdateByID(id, updates.Name, updates.Position)
	row.Scan(&material.Id, &material.Name, &material.Position, &material.CreatedAt, &material.UpdatedAt)

	return material, nil
}

func (t *TeacherService) NewTeacher(newTea PayloadTeacher) Teacher {
	var tea Teacher
	row := t.repo.NewTeacher(newTea.Name, newTea.Position)
	row.Scan(&tea.Id, &tea.Name, &tea.Position, &tea.CreatedAt, &tea.UpdatedAt)
	return tea
}

func (t *TeacherService) DeleteTeacher(id int) (bool, error) {
	err := t.repo.DeleteTeacher(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *TeacherService) TeacherExistsByID(id int) (bool, error) {
	valid, err := t.repo.TeacherExistsByID(id)
	if err != nil {
		return valid, errors.New("Teacher not found")
	}
	return valid, nil
}
