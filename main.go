package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/Boobuh/golang-school-project/dal"

	"github.com/Boobuh/golang-school-project/handler"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	logger.SetOutput(f)
	repo := dal.NewRepository()
	router := handler.NewRouter(repo, logger)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:4040",
	}
	log.Fatal(srv.ListenAndServe())
}
