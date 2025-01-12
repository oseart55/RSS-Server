package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		Url string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		log.Println(err)
		respondError(w, 400, fmt.Sprintln(err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       parameters.Url,
		UserID:    user.ID,
	})
	if err != nil {
		log.Println(err)
		respondError(w, 400, fmt.Sprintln(err))
		return
	}
	respondJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feed, err := apiCfg.DB.GetFeeds(r.Context(), user.ID)
	if err != nil {
		log.Println(err)
		respondError(w, 400, fmt.Sprintln(err))
		return
	}
	respondJSON(w, 201, databaseFeedsToFeeds(feed))
}
