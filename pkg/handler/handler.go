package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"typo_back/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/words", wordsHandlers(h))
	r.Mount("/auth", authHandlers(h))
	return r
}

func authHandlers(h *Handler) http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Post("/sign-in", func(w http.ResponseWriter, r *http.Request) {
			h.signIn(w, r)
		})
		r.Post("/sign-up", func(w http.ResponseWriter, r *http.Request) {
			h.signUp(w, r)
		})
	})

	return rg
}

func wordsHandlers(h *Handler) http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			h.fetchWords(w, r)
		})
	})

	return rg
}
