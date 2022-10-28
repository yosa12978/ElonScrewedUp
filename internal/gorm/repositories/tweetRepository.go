package repositories

import (
	"github.com/yosa12978/ElonScrewedUp/internal/domain/models"
	"github.com/yosa12978/ElonScrewedUp/internal/domain/repositories"
	db "github.com/yosa12978/ElonScrewedUp/internal/gorm"
	gmodels "github.com/yosa12978/ElonScrewedUp/internal/gorm/models"
	"gorm.io/gorm"
)

type tweetRepository struct {
	db *gorm.DB
}

func NewTweetRepositoryGorm() repositories.TweetRepsoitory {
	return &tweetRepository{db: db.GetDB()}
}

func (tr *tweetRepository) Create(tweet *models.Tweet) error {
	gtweet := &gmodels.TweetGorm{
		UUID:    tweet.UUID,
		Text:    tweet.Text,
		Pubdate: tweet.Pubdate,
		Url:     tweet.Url,
	}
	return tr.db.Create(gtweet).Error
}

func (tr *tweetRepository) GetAll() ([]models.Tweet, error) {
	var tweets []models.Tweet
	err := tr.db.Raw("SELECT * FROM Tweets ORDER BY pubdate DESC").Scan(&tweets).Error
	return tweets, err
}

func (tr *tweetRepository) GetByUUID(uuid string) (*models.Tweet, error) {
	var gtweet *gmodels.TweetGorm
	err := tr.db.First(&gtweet, "uuid = ?", uuid).Error
	return gtweet.Map(), err
}

func (tr *tweetRepository) Update(tweet *models.Tweet) error {
	var gtweet gmodels.TweetGorm
	err := tr.db.First(&gtweet, "uuid = ?", tweet.UUID).Error
	if err != nil {
		return err
	}
	gtweet.Text = tweet.Text
	gtweet.Url = tweet.Url
	tr.db.Save(&gtweet)
	return nil
}

func (tr *tweetRepository) Delete(uuid string) error {
	return tr.db.Where("uuid = ?", uuid).Delete(&gmodels.TweetGorm{}).Error
}
