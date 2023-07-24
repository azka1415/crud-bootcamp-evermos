package models

import (
	"errors"
	"time"

	"github.com/azka1415/crud-bootcamp-evermos/db"
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

type TeacherService struct{}

func NewTeacherService() *TeacherService {
	return &TeacherService{}
}

func (t *TeacherService) GetAll(limit, page int, sort, field string) ([]Teacher, error) {

	db, err := db.GetDB()
	offset := (page - 1) * limit
	if err != nil {
		return nil, err
	}
	var teachers []Teacher

	teaRepo := repository.NewTeacherRepository(db)
	rows, err := teaRepo.GetAll(field, sort, limit, offset)
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
	db, err := db.GetDB()
	var teacher Teacher
	if err != nil {
		return teacher, err
	}

	teaRepo := repository.NewTeacherRepository(db)

	row := teaRepo.GetByID(id)
	row.Scan(&teacher.Id, &teacher.Name, &teacher.Position, &teacher.CreatedAt, &teacher.UpdatedAt)

	return teacher, nil
}

func (t *TeacherService) UpdateTeacher(id int, updates PayloadTeacher) (Teacher, error) {
	db, err := db.GetDB()
	if err != nil {
		return Teacher{}, err
	}

	teaRepo := repository.NewTeacherRepository(db)
	var material Teacher

	row := teaRepo.UpdateByID(id, updates.Name, updates.Position)
	row.Scan(&material.Id, &material.Name, &material.Position, &material.CreatedAt, &material.UpdatedAt)

	return material, nil
}

func (t *TeacherService) NewTeacher(newTea PayloadTeacher) (Teacher, error) {
	db, err := db.GetDB()

	teaRepo := repository.NewTeacherRepository(db)
	var tea Teacher
	row := teaRepo.NewTeacher(newTea.Name, newTea.Position)
	row.Scan(&tea.Id, &tea.Name, &tea.Position, &tea.CreatedAt, &tea.UpdatedAt)
	return tea, err
}

func (t *TeacherService) DeleteTeacher(id int) (bool, error) {
	db, err := db.GetDB()
	if err != nil {
		return false, err
	}
	teaRepo := repository.NewTeacherRepository(db)

	err = teaRepo.DeleteTeacher(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *TeacherService) TeacherExistsByID(id int) (bool, error) {
	db, err := db.GetDB()
	if err != nil {
		return false, err
	}

	teaRepo := repository.NewTeacherRepository(db)
	valid, err := teaRepo.TeacherExistsByID(id)
	if err != nil {
		return valid, errors.New("Teacher not found")
	}
	return valid, nil
}
