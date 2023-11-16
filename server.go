package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	cloudstorage "github.com/danny-yamamoto/go-api-samples/internal/storage"
	users "github.com/danny-yamamoto/go-api-samples/internal/users"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/api/option"
	"google.golang.org/api/storage/v1"
)

type Handler struct {
	storage *storage.Service
	db      *sql.DB
}

func NewHandler(storage *storage.Service, db *sql.DB) *Handler {
	return &Handler{storage: storage, db: db}
}

func main() {
	ctx := context.Background()
	storage, err := storage.NewService(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Fatalf("Failed to initialize Google Storage service: %s", err)
	}
	dataSource := os.Getenv("DATABASE_URL")
	driver, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	handler := NewHandler(storage, driver)
	http.HandleFunc("/storage", cloudstorage.New(handler.storage))
	http.HandleFunc("/users", users.New(handler.db))
	port := "8080"
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Listening on http://%s ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
