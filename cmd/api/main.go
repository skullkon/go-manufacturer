package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/skullkon/go-manufacturer/internal/api"
)

func main() {

	router := api.NewRouter()

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router,
	}
	logrus.Info("Starting server on port: " + os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		logrus.Fatal("Error starting server: " + err.Error())
		return
	}
}
