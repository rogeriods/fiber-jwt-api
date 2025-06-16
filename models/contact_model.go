package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model        // adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name" gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	Email      string `json:"email" gorm:"type:varchar(200)"`
	Phone      string `json:"phone" gorm:"type:varchar(20)"`
	UserID     uint   `json:"user_id" gorm:"not null;index"` // Foreign key to User
	User       User   `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
