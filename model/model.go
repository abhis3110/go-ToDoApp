package model

import (
	"time"
)

type User struct{
	Info
	Username string     `json:"username"`
	Password string     `json:"password"`
	Todos    []Task     `json:"todos"`
}

type Task struct{
	Info
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	UserID        uint         `json:"userid"`
}

type Info struct{
	ID        uint         `gorm:"primary_key" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt *time.Time   `json:"deleted_at"`
}