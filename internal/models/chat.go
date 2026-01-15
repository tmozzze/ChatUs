package models

import "time"

type Chat struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	CreatedAt time.Time `json:"created_at"`

	Messages []Message `gorm:"constraint:OnDelete:CASCADE;" json:"messages,omitempty"`
}
