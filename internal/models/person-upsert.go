package models

type PersonUpsert struct {
	Name 	string 	`json:"name" binding:"required,min=1"`
	Age 	*int32 	`json:"age,omitempty"`
	Address *string 	`json:"address,omitempty"`
	Work	*string	`json:"work,omitempty"`
}