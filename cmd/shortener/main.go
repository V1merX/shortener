package main

import (
	"github.com/V1merX/shorter/internal/app"
	"github.com/V1merX/shorter/internal/store"
)

const (
	host string = "localhost:8080"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	storage := store.NewMemoryStorage()
	h := app.NewHandler(storage)
	
	if err := app.StartHTTPServer(host, h); err != nil {
		return err
	}
	return nil
}