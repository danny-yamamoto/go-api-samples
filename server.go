package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	//"github.com/danny-yamamoto/go-api-samples/internal/storage"
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

func (h *Handler) storageHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	fmt.Println("hello")
	ctx := context.Background()
	storage, err := storage.NewService(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		fmt.Printf("credentials not found. %s", err)
	}
	dataSource := os.Getenv("DB_URL")
	driver, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		fmt.Printf("data source not found. %s", err)
	}
	handler := NewHandler(storage, driver)
	http.HandleFunc("/storage", handler.storageHandler)
	// new handler
	port := "8080"
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	http.ListenAndServe(addr, nil)
	// handle func
	// listen
}
