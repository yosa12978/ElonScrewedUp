package models

import (
	"time"

	"github.com/yosa12978/ElonScrewedUp/internal/domain/models"
)

type TweetGorm struct {
	Gid     uint      `gorm:"column:id;primaryKey"`
	UUID    string    `gorm:"uniqueIndex;column:uuid"`
	Url     string    `gorm:"column:url"`
	Pubdate time.Time `gorm:"column:pubdate"`
	Text    string    `gorm:"column:text"`
}

func (t *TweetGorm) Map() *models.Tweet {
	if t == nil {
		return nil
	}
	return &models.Tweet{
		UUID:    t.UUID,
		Text:    t.Text,
		Pubdate: t.Pubdate,
		Url:     t.Url,
	}
}

func (t *TweetGorm) TableName() string {
	return "Tweets"
}