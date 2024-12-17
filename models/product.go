package models

import (
	"time"
)

type Product struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	NamaProduct string    `gorm:"type:varchar(300)" json:"nama_product" binding:"max=300,required"`
	Deskripsi   string    `gorm:"type:text" json:"deskripsi"`
	UserId      int64     `json:"user_id"`
	User        User      `gorm:"foreignKey:UserId"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
