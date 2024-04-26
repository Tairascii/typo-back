package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"typo_back"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("something went wrong when getting body from request for sign up")
	}
	var user typo_back.User
	if err := json.Unmarshal(userBody, &user); err != nil {
		log.Fatalf(err.Error())
	}

	userId, err := h.services.Auth.CreateUser(ctx, user)

	log.Println(userId)

	if err != nil {
		log.Println(err.Error())
	}
}
