package converters

import (
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
)

func ConvertPersonUpsertToPerson(person *models.PersonUpsert) *models.Person {
	return &models.Person{
		Name: person.Name,
		Age: person.Age,
		Address: person.Address,
		Work: person.Work,
	}
}
