package models

import (
	"database/sql"
	"errors"
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
		Rating      uint8     `json:"rating" gorm:"type:uint;not null"`
		Status      bool      `json:"status"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

//
func (p *Books) Get(db *sql.DB) error {
	return errors.New("Not implemented")
}
//
func (p *Books) Update(db *sql.DB) error {
	return errors.New("Not implemented")
}
//
func (p *Books) Delete(db *sql.DB) error {
	return errors.New("Not implemented")
}

//
func (p *Books) Create(db *sql.DB) error {
	return errors.New("Not implemented")
}

// ModelName describe model name
func (b *Books) ModelName() string {
	return fmt.Sprintf("%T", b)
}
