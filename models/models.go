package models

import (
	"time"
)

type Author struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Word struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	AuthorID  uint      `json:"-"`
	Author    Author    `json:"author,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
