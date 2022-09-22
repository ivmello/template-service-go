package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID  `gorm:"type:uuid" json:"id"`
	Name      string     `validate:"required" gorm:"type:string" json:"name"`
	Username  string     `validate:"required" gorm:"type:string" json:"username"`
	Email     string     `validate:"required" gorm:"type:string" json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
