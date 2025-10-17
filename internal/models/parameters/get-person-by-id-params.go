package parameters

type GetPersonByIdParams struct {
	ID int32 `uri:"id" binding:"required,min=0,max=2147483647"`
}
