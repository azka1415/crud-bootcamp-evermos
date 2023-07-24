package models

import (
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
