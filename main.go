package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"typo_back/assets"
)

var db *mgo.Database
var rnd *renderer.Render

const (
	hostName       string = "localhost:27017"
	dbName         string = "typo"
	collectionName string = "typo"
	port           string = ":9000"
)

type (
	resultModel struct {
		Id        bson.ObjectId `bson:"_id,omitempty"`
		WPM       int16         `bson:"wpm"`
		Accuracy  int8          `bson:"accuracy"`
		Timestamp time.Time     `bson:"timestamp"`
	}
	result struct {
		ID        string    `json:"id"`
		WPM       int16     `json:"WPM"`
		Accuracy  int8      `json:"accuracy"`
		Timestamp time.Time `json:"timestamp"`
	}
)

func init() {
	//sess, err := mgo.Dial(hostName)
	rnd = renderer.New()
	//checkError(err, "something with db")

	//sess.SetMode(mgo.Monotonic, true)
	//db = sess.DB(dbName)
}

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/words", wordsHandlers())

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		log.Println("listening on port", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen:%s\n", err)
		}
	}()

	<-stopChan
	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("something went wrong while shutting down server")
	}

	defer cancel()
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

	err := rnd.JSON(w, http.StatusOK, renderer.M{
		"data": words,
	})

	if err != nil {
		log.Println("something wrong with words")
	}
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}
