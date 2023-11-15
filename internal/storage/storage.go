package storage

import (
	"io"

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
