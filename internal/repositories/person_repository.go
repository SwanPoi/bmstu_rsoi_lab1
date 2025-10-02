package repositories

import (
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
	"gorm.io/gorm"
)

type PersonRepository struct {
	DB *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{DB: db}
}

func (r *PersonRepository) GetAll() ([]models.Person, error) {
	var persons []models.Person
	err := r.DB.Find(&persons).Error

	return persons, err
}

func (r *PersonRepository) AddPerson(person *models.Person) (int32, error) {
	err := r.DB.Create(person).Error

	if err != nil {
		return 0, err
	}

	return person.Id, nil
}