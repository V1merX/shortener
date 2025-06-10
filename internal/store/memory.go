package store

import (
	"errors"
)

type MemoryStorage struct {
	urls map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{urls: make(map[string]string)}
}

func (ms *MemoryStorage) Get(key string) (string, error) {
	value, ok := ms.urls[key]
	if !ok {
		return "", errors.New("error: failed to get url by key")
	}
	return value, nil
}

func (ms *MemoryStorage) Set(key, value string) {
	ms.urls[key] = value
}

