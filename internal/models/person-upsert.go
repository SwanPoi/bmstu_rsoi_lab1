package models

type PersonUpsert struct {
	Name 	string 	`json:"name" binding:"required,min=1"`
	Age 	int32 	`json:"age" binding:"min=0"`
	Address string 	`json:"address"`
	Work	string	`json:"work"`
}