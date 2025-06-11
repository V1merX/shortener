package app

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/V1merX/shorter/internal/store"
)

type Handler struct {
	MemoryStorage *store.MemoryStorage
}

func NewHandler(memoryStorage *store.MemoryStorage) *Handler {
	return &Handler{
		MemoryStorage: memoryStorage,
	}
}

var alph string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func (h *Handler) getLink(w http.ResponseWriter, r *http.Request) {
	targetURL, err := h.MemoryStorage.Get(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, targetURL, http.StatusTemporaryRedirect)
}

func (h *Handler) createLink(w http.ResponseWriter, r *http.Request) {
	rand.New(rand.NewSource(time.Now().Unix()))

	var randomString string

	for i := 0; i < 8; i++ {
		randomString += string(alph[rand.Intn(len(alph))])
	}
	w.Header().Add(
		"Content-Type", "text/plain",
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.MemoryStorage.Set(randomString, string(body))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", randomString)))
}
