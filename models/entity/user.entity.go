package entity

import "time"

// User is a representation of a user
type User struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:timestamp"`
}
