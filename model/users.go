package model

import "time"

type Users struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email        string    `gorm:"column:email" json:"email"`
	Password     string    `gorm:"column:password" json:"-"`
	Name         string    `gorm:"column:name" json:"name"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Users) TableName() string {
	return "users"
}
