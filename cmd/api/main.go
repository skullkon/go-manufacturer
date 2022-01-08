package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/skullkon/go-manufacturer/internal/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorln("Error loading .env file in main: " + err.Error())
		return
	}
	router := api.NewRouter()

	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		logrus.Fatal("Error starting server: " + err.Error())
		return
	}
}
