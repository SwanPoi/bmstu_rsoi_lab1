package services

import (
	"errors"
	"testing"

	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/converters"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPersonRepository struct {
	mock.Mock
}

func (m *MockPersonRepository) GetAll() ([]models.Person, error) {
	args := m.Called()

	return args.Get(0).([]models.Person), args.Error(1)
}

func (m *MockPersonRepository) GetById(id int32) (*models.Person, error) {
	args := m.Called(id)

	if person := args.Get(0); person != nil {
		return person.(*models.Person), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockPersonRepository) AddPerson(person *models.Person) (int32, error) {
	args := m.Called(person)

	return args.Get(0).(int32), args.Error(1)
}

func (m *MockPersonRepository) DeletePerson(id int32) (error) {
	args := m.Called(id)

	return args.Error(0)
}

func (m *MockPersonRepository) UpdatePerson(id int32, upsert *models.PersonUpsert) (*models.Person, error) {
	args := m.Called(id, upsert)

	if person := args.Get(0); person != nil {
		return person.(*models.Person), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestPersonService_GetPersonById_NotFound(t *testing.T) {
	mockRepo := new(MockPersonRepository)
	service := NewPersonService(mockRepo)

	id := int32(1)
	expectedError := errors.New(models.ErrorNotFound)

	mockRepo.On("GetById", id).Return((*models.Person)(nil), expectedError)

	_, err := service.GetById(id)

	assert.True(t, errors.Is(err, expectedError))
	mockRepo.AssertExpectations(t)
}

func TestPersonService_GetPersonById_Success(t *testing.T) {
	mockRepo := new(MockPersonRepository)
	service := NewPersonService(mockRepo)

	id := int32(1)
	expectedPerson := &models.Person{
		Id: id,
		Name: "Bob",
	}

	mockRepo.On("GetById", id).Return(expectedPerson, nil)

	person, err := service.GetById(id)

	assert.Nil(t, err)
	assert.Equal(t, expectedPerson, person)
	mockRepo.AssertExpectations(t)
}

func TestPersonService_AddPerson_Success(t *testing.T) {
	mockRepo := new(MockPersonRepository)
	service := NewPersonService(mockRepo)

	id := int32(1)
	name := "Bob"
	age := int32(20)
	address := "Moscow"
	work := "2Gis"

	upsertPerson := &models.PersonUpsert{
		Name: name,
		Age: &age,
		Address: &address,
		Work: &work,
	}

	newPerson := converters.ConvertPersonUpsertToPerson(upsertPerson)

	mockRepo.On("AddPerson", newPerson).Return(id, nil)

	personId, err := service.AddPerson(upsertPerson)

	assert.Equal(t, personId, id)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPersonService_UpdatePerson_Success(t *testing.T) {
	mockRepo := new(MockPersonRepository)
	service := NewPersonService(mockRepo)

	id := int32(1)
	name := "Bob"
	age := int32(20)
	address := "Moscow"
	work := "2Gis"

	upsertPerson := &models.PersonUpsert{
		Name: name,
		Age: &age,
		Address: &address,
		Work: &work,
	}

	expectedPerson := &models.Person{
		Id: id,
		Name: name,
		Age: age,
		Address: address,
		Work: work,
	}

	mockRepo.On("UpdatePerson", id, upsertPerson).Return(expectedPerson, nil)

	person, err := service.UpdatePerson(id, upsertPerson)

	assert.Equal(t, person, expectedPerson)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPersonService_UpdatePerson_NotFound(t *testing.T) {
	mockRepo := new(MockPersonRepository)
	service := NewPersonService(mockRepo)

	id := int32(1)
	name := "Bob"
	age := int32(20)
	address := "Moscow"
	work := "2Gis"

	upsertPerson := &models.PersonUpsert{
		Name: name,
		Age: &age,
		Address: &address,
		Work: &work,
	}

	expectedError := errors.New(models.ErrorNotFound)

	mockRepo.On("UpdatePerson", id, upsertPerson).Return((*models.Person)(nil), expectedError)

	_, err := service.UpdatePerson(id, upsertPerson)

	assert.True(t, errors.Is(err, expectedError))
	mockRepo.AssertExpectations(t)
}
