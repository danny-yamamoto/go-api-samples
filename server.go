package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	cloudstorage "github.com/danny-yamamoto/go-api-samples/internal/storage"
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

func (h *Handler) storageHandler(w http.ResponseWriter, r *http.Request) {
	data, err := cloudstorage.GetObject(h.storage, cloudstorage.StorageQuery{Bucket: r.URL.Query().Get("bucket"), Object: r.URL.Query().Get("object")})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, data)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	//response, _ := json.Marshal(payload)
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
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
	port := "8080"
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	http.ListenAndServe(addr, nil)
}
