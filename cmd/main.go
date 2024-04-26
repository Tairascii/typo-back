package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"typo_back"
	"typo_back/pkg/handler"
	"typo_back/pkg/repository"
	"typo_back/pkg/service"
)

var db *mgo.Database
var rnd *renderer.Render

const (
	hostName       string = "localhost:27017"
	dbName         string = "typo"
	collectionName string = "typo"
	port           string = ":9000"
)

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main2() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

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

func main() {
	db, conErr := repository.NewMongoDB(context.Background())
	defer func() {
		if err := db.Disconnect(context.Background()); err != nil {
			log.Fatalf("something went wrong while disconnecting %s", err.Error())
		}
	}()
	if err := initConfigs(); err != nil {
		log.Fatalf("something with configs: %s", err.Error())
	}
	if conErr != nil {
		log.Fatalf("error while connecting to db %s", conErr.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(typo_back.Server)
	log.Println("listening on port", ":8000")
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("something went wrong while running server %s", err.Error())
	}
}
