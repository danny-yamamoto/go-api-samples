package users

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"testing"
)

func setupMockDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "../../unit_test.db")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func TestGetUsers(t *testing.T) {
	db := setupMockDB(t)

	tests := []struct {
		name         string
		userId       int64
		expectedJSON string
		want         bool
	}{
		{"Normal pattern a", 10000, `{"user_id":10000,"email_address":"marc@example.com","created_at":0,"deleted":1,"settings":""}`, true},
		{"Normal pattern b", 100, `{"user_id":100,"email_address":"marc@example.com","created_at":1,"deleted":1,"settings":""}`, false},
		{"Failed pattern a", 100, `{"user_id":100,"email_address":"alex@example.com","created_at":1,"deleted":0,"settings":""}`, true},
	}
	for _, tc := range tests {
		user, err := GetUsers(db, UserQuery{UserId: strconv.FormatInt(tc.userId, 10)})
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		userJson, err := json.Marshal(user)
		if err != nil {
			t.Errorf("Failed to marshal user data: %s", err)
		}
		if tc.want && string(userJson) != tc.expectedJSON {
			t.Errorf("Expected JSON %s for user ID %v, got %s", tc.expectedJSON, tc.userId, string(userJson))
		}
	}

}
