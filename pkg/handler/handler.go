package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"typo_back/assets"
	"typo_back/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/words", wordsHandlers())
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

func wordsHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fetchWords(w, r)
		})
	})

	return rg
}

func generateRandomWords() []string {
	return assets.Words
}

func fetchWords(w http.ResponseWriter, r *http.Request) {
	words := generateRandomWords()

	log.Println(words)
	//err := rnd.JSON(w, http.StatusOK, renderer.M{
	//	"data": words,
	//})
	//
	//if err != nil {
	//	log.Println("something wrong with words")
	//}
}
