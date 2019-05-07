package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// ErrorValueOfFieldRequired
	ErrorValueOfFieldRequired = "Field name %s is required"
	// ErrorOutOfRange is template for
	ErrorOutOfRange = "Field name %s is out of range"
	// BookFieldRatingMaxValue is max value of field rating
	BookFieldRatingMaxValue = 3
	BookFieldRatingMinValue = 0
)

type (
	Books struct {
		ID          uint64    `json:"id,string,omitempty" gorm:"primary_key"`
		Title       string    `json:"title" `
		Author      string    `json:"author"`
		Publisher   string    `json:"publisher"`
		PublishDate time.Time `json:"publish_date"`
		Rating      uint8     `json:"rating"`
		Status      bool      `json:"status"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

// ModelName describe model name
func (b *Books) ModelName() string {
	return fmt.Sprintf("%T", b)
}

// BeforeCreate is gorm hook
func (b *Books) BeforeCreate(scope *gorm.Scope) error {
	var errText string

	if b.Title == "" {
		errText = fmt.Sprintf(ErrorValueOfFieldRequired, "title")
		return errors.New(errText)
	}

	if b.Author == "" {
		errText = fmt.Sprintf(ErrorValueOfFieldRequired, "author")
		return errors.New(errText)
	}

	if b.Publisher == "" {
		errText = fmt.Sprintf(ErrorValueOfFieldRequired, "publisher")
		return errors.New(errText)
	}

	if b.Rating > BookFieldRatingMaxValue || b.Rating < BookFieldRatingMinValue{
		errText = fmt.Sprintf(ErrorOutOfRange, "rating")
		return errors.New(errText)
	}

	return nil
}

// BeforeUpdate is gorm hook
func (b *Books) BeforeUpdate(scope *gorm.Scope) error {
	var errText string

	if b.Rating > BookFieldRatingMaxValue || b.Rating < BookFieldRatingMinValue{
		errText = fmt.Sprintf(ErrorOutOfRange, "rating")
		return errors.New(errText)
	}

	return nil
}
