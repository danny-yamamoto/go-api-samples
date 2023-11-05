package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/storage/v1"
)

type StorageQuery struct {
	Bucket string `json:"bucket"`
	Object string `json:"object"`
}

type StorageResponse struct {
	Content string `json:"content"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) storageHandler(w http.ResponseWriter, r *http.Request) {
	bucket := r.URL.Query().Get("bucket")
	object := r.URL.Query().Get("object")

	rc, err := h.client.Objects.Get(bucket, object).Download()
	if err != nil {
		respondWithError(w, fmt.Sprintf("Failed to read object: %v", err))
		return
	}
	defer rc.Body.Close()

	data, err := io.ReadAll(rc.Body)
	if err != nil {
		respondWithError(w, fmt.Sprintf("Failed to read object data: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, StorageResponse{Content: string(data)})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, message string) {
	respondWithJSON(w, http.StatusInternalServerError, ErrorResponse{Error: message})
}

type Handler struct {
	client *storage.Service
}

func NewHandler(client *storage.Service) *Handler {
	return &Handler{client: client}
}

func main() {
	ctx := context.Background()
	client, err := storage.NewService(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		fmt.Printf("Failed to create client: %s", err)
		return
	}

	handler := NewHandler(client)
	http.HandleFunc("/storage", handler.storageHandler)

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("Listening on http://%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
