package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty" form:"id"`
	Name      string     `json:"name,omitempty" form:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt *time.Time
}

type Job struct {
	gorm.Model
	ID          uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty" form:"id"`
	Name        string     `json:"name,omitempty" form:"name"`
	Description string     `json:"description,omitempty" form:"description"`
	CreatedAt   *time.Time `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" form:"updated_at"`
}
