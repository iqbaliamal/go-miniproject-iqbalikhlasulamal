package response

import "time"

type ScholarshipResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    string    `json:"deadline"`
	Link        string    `json:"link"`
	Thumbnail   string    `json:"thumbnail"`
	UserID      uint64    `json:"user_id"`
	Author      string    `json:"author"`
	CategoryID  uint64    `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
