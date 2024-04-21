package handler

import (
	"log"
	"net/http"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
}
