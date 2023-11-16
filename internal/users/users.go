package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserId       int64  `json:"user_id"`
	EmailAddress string `json:"email_address"`
	CreatedAt    int64  `json:"created_at"`
	Deleted      int64  `json:"deleted"`
	Settings     string `json:"settings"`
}

type UserQuery struct {
	UserId string `json:"user_id"`
}

func GetUsers(db *sql.DB, query UserQuery) (*User, error) {
	userId := query.UserId
	var user User
	err := db.QueryRow("select * from users where user_id = ?", userId).Scan(&user.UserId, &user.EmailAddress, &user.CreatedAt, &user.Deleted, &user.Settings)
	if err != nil {
		log.Printf("Query Error: %s", err)
		return nil, err
	}
	return &user, nil
}

// Factory Function
func New(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := UserQuery{UserId: r.URL.Query().Get("user_id")}
		data, err := GetUsers(db, query)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(data)
	}
}
