package main

import (
	"fmt"
	"log"
	"main/internal/database"
	"main/internal/database/auth"
	"net/http"
)

type handleAuth func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler handleAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			log.Println(err)
			respondError(w, 400, fmt.Sprintln(err))
			return
		}
		user, err := apiCfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			log.Println(err)
			respondError(w, 400, fmt.Sprintln(err))
			return
		}
		handler(w, r, user)
	}
}
