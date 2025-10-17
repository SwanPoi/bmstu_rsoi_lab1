package services

import (
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/converters"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/repositories"
)

type PersonService struct {
	repo repositories.PersonRepositoryInterface
}

func NewPersonService(repo repositories.PersonRepositoryInterface) *PersonService {
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

func (s *PersonService) GetById(id int32) (*models.Person, error) {
	person, err := s.repo.GetById(id)

	return person, err
}

func (s *PersonService) DeletePerson(id int32) (error) {
	return s.repo.DeletePerson(id)
}

func (s *PersonService) UpdatePerson(id int32, person *models.PersonUpsert) (*models.Person, error) {
	return s.repo.UpdatePerson(id, person)
}