package app

import (
	"net/http"
)

func StartHTTPServer(host string, h *Handler) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.indexHandler)

	if err := http.ListenAndServe(host, mux); err != nil {
		return err
	}

	return nil
}