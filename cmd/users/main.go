package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserId       int64  `json:"user_id"`
	EmailAddress string `json:"email_address"`
	CreatedAt    int64  `json:"created_at"`
	Deleted      int64  `json:"deleted"`
	Settings     string `json:"settings"`
}

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (h Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	var user User
	err := h.db.QueryRow("select * from users where user_id = ?", userId).Scan(&user.UserId, &user.EmailAddress, &user.CreatedAt, &user.Deleted, &user.Settings)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func main() {
	key := "DATABASE_URL"
	dbUrl := os.Getenv(key)
	client, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		fmt.Printf("Failed to create connection: %s", err)
		return
	}
	defer client.Close()
	handler := NewHandler(client)
	http.HandleFunc("/users", handler.UserHandler)
	port := "8080"
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("Listening on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
