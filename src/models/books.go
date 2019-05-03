package models

import (
	"fmt"
	"time"
)

type (
	Books struct {
		ID          uint      `json:"ID" gorm:"primary_key"`
		Title       string    `json:"title" `
		Author      string    `json:"author"`
		Publisher   string    `json:"publisher"`
		PublishDate time.Time `json:"publish_date"`
		Rating      uint8     `json:"rating" gorm:"type:varchar(100);unique;not null"`
		Status      bool      `json:"status"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

// ModelName describe model name
func (b *Books) ModelName() string {
	return fmt.Sprintf("%T", b)
}
