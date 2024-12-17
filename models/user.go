package models

import (
	"time"
)

type User struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(300)" json:"nama" binding:"max=300,required"`
	Email     string    `gorm:"type:varchar(255)" json:"email" binding:"required,email"`
	Password  string    `gorm:"type:varchar(255)" json:"password" binding:"min=3"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Products  []Product `gorm:"foreignKey:UserId" json:"products"`
}
