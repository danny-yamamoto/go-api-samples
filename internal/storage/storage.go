package storage

import (
	"encoding/json"
	"io"
	"net/http"

	"google.golang.org/api/storage/v1"
)

type StorageQuery struct {
	Bucket string `json:"bucket"`
	Object string `json:"object"`
}

type StorageResponse struct {
	Content string `json:"content"`
}

func GetObject(client *storage.Service, query StorageQuery) (*StorageResponse, error) {
	rc, err := client.Objects.Get(query.Bucket, query.Object).Download()
	if err != nil {
		return nil, err
	}
	defer rc.Body.Close()
	data, err := io.ReadAll(rc.Body)
	if err != nil {
		return nil, err
	}
	ret := &StorageResponse{Content: string(data)}
	return ret, nil
}

// Factory Function
func New(storage *storage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := StorageQuery{Bucket: r.URL.Query().Get("bucket"), Object: r.URL.Query().Get("object")}
		data, err := GetObject(storage, query)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(data)
	}
}
