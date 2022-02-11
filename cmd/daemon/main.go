package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/AlfredDobradi/shortener/internal/logging"
	"github.com/AlfredDobradi/shortener/internal/server"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func main() {
	listener := server.New()

	log.WithFields(logrus.Fields{
		"address": listener.Addr,
	}).Info("Web service starting")

	if err := listener.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Listener error")
	}
}

func init() {
	if err := logging.InitLogger(); err != nil {
		panic(fmt.Sprintf("Error creating logger: %v", err))
	}

	log = logging.Get()
}
