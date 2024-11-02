// models/models.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type LoginDTO struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role" gorm:"default:user"` 
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Event struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date"`
	CreatedBy   uint           `json:"created_by"` 
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Club struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Members   []User         `gorm:"many2many:club_members" json:"members"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
