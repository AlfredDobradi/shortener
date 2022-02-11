package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AlfredDobradi/shortener/internal/logging"
	"github.com/sirupsen/logrus"
)

type RecordingResponseWriter struct {
	http.ResponseWriter

	Length int
	Status int
}

func NewRecordingResponseWriter(w http.ResponseWriter) RecordingResponseWriter {
	return RecordingResponseWriter{
		ResponseWriter: w,
		Length:         0,
		Status:         200,
	}
}

func (r *RecordingResponseWriter) Write(d []byte) (int, error) {
	n, err := r.ResponseWriter.Write(d)
	r.Length = n
	return n, err
}

func (r *RecordingResponseWriter) WriteHeader(statusCode int) {
	r.Status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func Logging(next http.Handler) http.Handler {
	log := logging.Get()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewRecordingResponseWriter(w)

		start := time.Now()
		next.ServeHTTP(&rw, r)

		log.WithFields(logrus.Fields{
			"user-agent":     r.UserAgent(),
			"user":           r.URL.User,
			"remote-address": r.RemoteAddr,
			"proto":          r.Proto,
			"method":         r.Method,
			"request":        r.RequestURI,
			"bytes-sent":     rw.Length,
			"duration":       time.Since(start).String(),
			"status":         rw.Status,
		}).Info(fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	})
}
