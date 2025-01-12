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

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		log.Println(err)
		respondError(w, 400, fmt.Sprintln(err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parameters.Name,
	})
	if err != nil {
		log.Println(err)
		respondError(w, 400, fmt.Sprintln(err))
		return
	}
	respondJSON(w, 201, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondJSON(w, 200, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handleGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		ID:    user.ID,
		Limit: 20,
	})
	if err != nil {
		log.Println("Error fetching posts for user: ", err)
		respondError(w, 400, fmt.Sprintln("Error fetching posts for user: ", err))
		return
	}
	respondJSON(w, 200, databasePostsToPosts(posts))
}
