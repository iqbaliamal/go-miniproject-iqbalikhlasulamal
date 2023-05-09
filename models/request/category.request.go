package request

type CategoryRequest struct {
	// name required and minimum length 3 and unique
	Name string `json:"name" validate:"required,min=3"`
}
