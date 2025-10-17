package repositories

import (
	"errors"

	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
	"gorm.io/gorm"
)

type PersonRepositoryInterface interface {
	GetAll() ([]models.Person, error)
	GetById(id int32) (*models.Person, error)
	AddPerson(person *models.Person) (int32, error)
	DeletePerson(id int32) (error)
	UpdatePerson(id int32, upsert *models.PersonUpsert) (*models.Person, error)
}

type PersonRepository struct {
	DB *gorm.DB
}

const NotFoundError = "record not found"

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{DB: db}
}

func (r *PersonRepository) GetAll() ([]models.Person, error) {
	var persons []models.Person
	err := r.DB.Find(&persons).Error

	return persons, err
}

func (r *PersonRepository) GetById(id int32) (*models.Person, error) {
	var person models.Person
	
	if err := r.DB.First(&person, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(models.ErrorNotFound)
		}

		return nil, err
	}

	return &person, nil
}

func (r *PersonRepository) AddPerson(person *models.Person) (int32, error) {
	err := r.DB.Create(person).Error

	if err != nil {
		return 0, err
	}

	return person.Id, nil
}

func (r *PersonRepository) DeletePerson(id int32) (error) {
	result := r.DB.Delete(&models.Person{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(models.ErrorNotFound)
	}

	return nil
}

func (r *PersonRepository) UpdatePerson(id int32, upsert *models.PersonUpsert) (*models.Person, error) {
	var person models.Person

	if err := r.DB.First(&person, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(models.ErrorNotFound)
		}

		return nil, err
	}

	person.Name = upsert.Name

	if upsert.Address != nil {
		person.Address = *upsert.Address
	}

	if upsert.Age != nil {
		person.Age = *upsert.Age
	}

	if upsert.Work != nil {
		person.Work = *upsert.Work
	}
	

	if err := r.DB.Save(&person).Error; err != nil {
		return nil, err
	}

	return &person, nil
}