package request

type ScholarshipRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link"`
	Deadline    string `json:"deadline"`
	CategoryID  uint64 `json:"category_id" validate:"required"`
}
