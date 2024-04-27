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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user typo_back.User
	if err := json.Unmarshal(userBody, &user); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := h.services.Auth.CreateUser(ctx, user)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := struct {
		Id string `json:"id"`
	}{Id: userId.Hex()}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
