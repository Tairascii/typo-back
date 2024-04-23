package handler

import (
	"log"
	"net/http"
	"typo_back/assets"
)

func generateRandomWords() []string {
	return assets.Words
}

func (h *Handler) fetchWords(w http.ResponseWriter, r *http.Request) {
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
