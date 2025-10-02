package services

import (
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/repositories"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/converters"
)

type PersonService struct {
	repo *repositories.PersonRepository
}

func NewPersonService(repo *repositories.PersonRepository) *PersonService {
	return &PersonService{ repo: repo }
}

func (s *PersonService) AddPerson(person *models.PersonUpsert) (int32, error) {
	createdPerson := converters.ConvertPersonUpsertToPerson(person)
	id, err := s.repo.AddPerson(createdPerson)

	return id, err
}

func (s *PersonService) GetAll() ([]models.Person, error) {
	return s.repo.GetAll()
}