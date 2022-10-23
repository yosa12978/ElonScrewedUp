package repositories

import (
	"database/sql"

	"github.com/yosa12978/ElonScrewedUp/internal/db"
	"github.com/yosa12978/ElonScrewedUp/internal/domain/models"
	"github.com/yosa12978/ElonScrewedUp/internal/domain/repositories"
)

type tweetRepository struct {
	db *sql.DB
}

func NewTweetRepository() repositories.TweetRepsoitory {
	return &tweetRepository{db: db.GetDB()}
}

func (tr *tweetRepository) Create(tweet *models.Tweet) error {
	stmt, err := tr.db.Prepare("INSERT INTO Tweets (Id, Url, Text, Date) VALUES(? ? ? ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tweet.Id, tweet.Url, tweet.Text, tweet.Date)
	return err
}

func (tr *tweetRepository) GetAll() ([]models.Tweet, error) {
	rows, err := tr.db.Query("SELECT * FROM Tweets ORDER BY Date DESC")
	if err != nil {
		return nil, nil
	}
	var tweets []models.Tweet
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(tweets[i].Id, tweets[i].Url, tweets[i].Text, tweets[i].Date)
		if err != nil {
			return nil, err
		}
	}
	return tweets, nil
}

func (tr *tweetRepository) GetById(id uint) (*models.Tweet, error) {
	return nil, nil
}

func (tr *tweetRepository) Update(tweet *models.Tweet) error {
	return nil
}

func (tr *tweetRepository) Delete(id uint) error {
	return nil
}
