package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte("pong"))
		if err != nil {
			return
		}
	})

	logrus.Info("Challenge App Starting...")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}
}
