package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentials == "" {
		log.Fatal("Environment variable GOOGLE_APPLICATION_CREDENTIALS is not set.")
	}
	storage, err := storage.NewService(ctx, option.WithCredentialsFile(credentials))
	if err != nil {
		log.Fatalf("Failed to initialize Google Storage service: %s", err)
	}

	dataSource := os.Getenv("DATABASE_URL")
	if dataSource == "" {
		log.Fatal("Environment Variable DATABASE_URL is not set.")
	}
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Minute * 5)
	defer db.Close()

	handler := NewHandler(storage, db)
	http.HandleFunc("/storage", cloudstorage.New(handler.storage))
	http.HandleFunc("/users", users.New(handler.db))
	port := "8080"
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Listening on http://%s ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
