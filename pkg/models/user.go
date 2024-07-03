package models

import "time"

type User struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Name       string
	RoleId     string
	Password   string
	MessId     string
	Email      string
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	IsVerified bool
}
