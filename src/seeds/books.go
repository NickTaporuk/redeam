package seeds

import (
	"math/rand"
	"time"

	"github.com/NickTaporuk/redeam/src/models"
	"github.com/bxcodec/faker/v3"
)

// Seeds is fake database data for initiate db
func Seeds() models.RedeamModels {

	seeds := []models.RedeamModel{
		&models.Books{
			Author:      faker.Username(),
			Title:       faker.Paragraph(),
			Publisher:   faker.Username(),
			PublishDate: time.Now(),
			Rating:      uint8(rand.Intn(3)),
			Status:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		&models.Books{
			Author:      faker.Username(),
			Title:       faker.Paragraph(),
			Publisher:   faker.Username(),
			PublishDate: time.Now(),
			Rating:      uint8(rand.Intn(3)),
			Status:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		&models.Books{
			Author:      faker.Username(),
			Title:       faker.Paragraph(),
			Publisher:   faker.Username(),
			PublishDate: time.Now(),
			Rating:      uint8(rand.Intn(3)),
			Status:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		&models.Books{
			Author:      faker.Username(),
			Title:       faker.Paragraph(),
			Publisher:   faker.Username(),
			PublishDate: time.Now(),
			Rating:      uint8(rand.Intn(3)),
			Status:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return seeds
}
