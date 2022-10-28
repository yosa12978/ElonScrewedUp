package endpoints

import (
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yosa12978/ElonScrewedUp/internal/domain/models"
	"github.com/yosa12978/ElonScrewedUp/internal/dtos"
	"github.com/yosa12978/ElonScrewedUp/internal/gorm/repositories"
	"github.com/yosa12978/ElonScrewedUp/pkg/helpers"
)

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.JsonDump(w, dtos.BadRequest)
		return
	}
	tweet, err := helpers.JsonLoad(string(body))
	if err != nil {
		helpers.JsonDump(w, dtos.ServerError)
		return
	}
	newTweet := models.Tweet{
		UUID:    uuid.NewString(),
		Text:    tweet.(dtos.CreateTweetRequest).Text,
		Url:     tweet.(dtos.CreateTweetRequest).Url,
		Pubdate: time.Now(),
	}
	repo := repositories.NewTweetRepositoryGorm()
	err = repo.Create(&newTweet)
	if err != nil {
		helpers.JsonDump(w, dtos.BadRequest)
		return
	}
	helpers.JsonDump(w, dtos.Created)
}

func GetTweets(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewTweetRepositoryGorm()
	tweets, err := repo.GetAll()
	if err != nil {
		helpers.JsonDump(w, dtos.ServerError)
		return
	}
	helpers.JsonDump(w, tweets)
}

func GetTweetByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	repo := repositories.NewTweetRepositoryGorm()
	tweet, err := repo.GetByUUID(vars["uuid"])
	if err != nil {
		helpers.JsonDump(w, dtos.NotFound)
		return
	}
	helpers.JsonDump(w, tweet)
}

func UpdateTweet(w http.ResponseWriter, r *http.Request) {

}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {

}
