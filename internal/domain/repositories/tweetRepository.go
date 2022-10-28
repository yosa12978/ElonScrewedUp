package repositories

import (
	"github.com/yosa12978/ElonScrewedUp/internal/domain/models"
)

type TweetRepsoitory interface {
	Create(tweet *models.Tweet) error
	GetAll() ([]models.Tweet, error)
	GetByUUID(uuid string) (*models.Tweet, error)
	Update(tweet *models.Tweet) error
	Delete(uuid string) error
}
