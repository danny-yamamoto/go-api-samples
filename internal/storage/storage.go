package storage

import "google.golang.org/api/storage/v1"

type StorageQuery struct {
	Bucket string `json:"bucket"`
	Object string `json:"object"`
}

func GetObject(client *storage.Service, query StorageQuery) {}
