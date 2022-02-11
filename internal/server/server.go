package server

import (
	"net/http"

	"github.com/AlfredDobradi/shortener/internal/server/middleware"
	"github.com/gorilla/mux"
)

type Listener struct {
	*http.Server
}

func New() *Listener {
	handler := mux.NewRouter()
	handler.Use(middleware.Logging)

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK")) // nolint
	})

	s := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: handler,
	}

	return &Listener{s}
}
