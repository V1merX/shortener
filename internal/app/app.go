package app

import (
	"net/http"
)

func StartHTTPServer(host string, h *Handler) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{id}", h.getLink)
	mux.HandleFunc("POST /", h.createLink)

	if err := http.ListenAndServe(host, mux); err != nil {
		return err
	}

	return nil
}