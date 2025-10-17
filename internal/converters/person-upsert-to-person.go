package converters

import (
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
)

func ConvertPersonUpsertToPerson(upsert *models.PersonUpsert) *models.Person {
	var person models.Person

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

	return &person
}
