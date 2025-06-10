package store

import (
	"errors"
	"sync"
)

type MemoryStorage struct {
	mx sync.RWMutex
	urls map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{urls: make(map[string]string)}
}

func (ms *MemoryStorage) Get(key string) (string, error) {
	ms.mx.RLock()
    defer ms.mx.RUnlock()

	value, ok := ms.urls[key]
	if !ok {
		return "", errors.New("error: failed to get url by key")
	}
	return value, nil
}

func (ms *MemoryStorage) Set(key, value string) {
	ms.mx.RLock()
    defer ms.mx.RUnlock()

	ms.urls[key] = value
}

