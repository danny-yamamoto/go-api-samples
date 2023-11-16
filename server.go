package main

import (
	"context"
	"database/sql"
	"encoding/json"
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

func (h *Handler) storageHandler(w http.ResponseWriter, r *http.Request) {
	data, err := cloudstorage.GetObject(h.storage, cloudstorage.StorageQuery{Bucket: r.URL.Query().Get("bucket"), Object: r.URL.Query().Get("object")})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, data)
}

func (h *Handler) userHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	query := users.UserQuery{UserId: userId}
	data, err := users.GetUsers(h.db, query)
	if err != nil {
		log.Printf("Internal Server Error: %s", err)
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, data)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func main() {
	ctx := context.Background()
	storage, err := storage.NewService(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		fmt.Printf("credentials not found. %s", err)
	}
	dataSource := os.Getenv("DATABASE_URL")
	driver, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		fmt.Printf("data source not found. %s", err)
	}
	handler := NewHandler(storage, driver)
	http.HandleFunc("/storage", handler.storageHandler)
	http.HandleFunc("/users", handler.userHandler)
	port := "8080"
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Listening on http://%s ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
