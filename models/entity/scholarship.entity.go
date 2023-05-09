package entity

import "time"

// Scholarship is a representation of a scholarship
type Scholarship struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:text"`
	Deadline    string    `json:"deadline" gorm:"type:varchar(255)"`
	Link        string    `json:"link" gorm:"type:text"`
	Thumbnail   string    `json:"thumbnail" gorm:"type:text"`
	UserID      uint64    `json:"user_id"`
	CategoryID  uint64    `json:"category_id"`
	User        User      `gorm:"foreignKey:UserID"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"type:timestamp"`
}
