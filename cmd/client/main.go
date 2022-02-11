package main

import (
	"fmt"

	"github.com/AlfredDobradi/shortener/internal/logging"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func main() {
	log.Printf("client")
}

func init() {
	if err := logging.InitLogger(); err != nil {
		panic(fmt.Sprintf("Error creating logger: %v", err))
	}

	log = logging.Get()
}
