package repositories

import "github.com/yosa12978/ElonScrewedUp/internal/domain/models"

type TweetRepsoitory interface {
	Create(tweet *models.Tweet) error
	GetAll() ([]models.Tweet, error)
	GetById(id uint) (*models.Tweet, error)
	Update(tweet *models.Tweet) error
	Delete(id uint) error
}
