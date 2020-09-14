package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type contextKey string

// Logging logs when a request arrives
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithField("path", r.RequestURI).
			WithField("method", r.Method).
			WithField("request_id", GetRequestID(r.Context())).
			Debug("Request arrived")
		next.ServeHTTP(w, r)
	})
}

// RequestID adds a request id to the request context
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		nctx := context.WithValue(r.Context(), contextKey("requestID"), requestID)
		next.ServeHTTP(w, r.WithContext(nctx))
	})
}

// GetRequestID returns the context request id value, if existent
func GetRequestID(ctx context.Context) string {
	id, ok := ctx.Value(contextKey("requestID")).(string)
	if !ok {
		logrus.Warn("requestID not found in context")
		return ""
	}
	return id
}

// Timing logs the request duration
func Timing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		d := time.Since(start)
		logrus.WithField("duration", d).
			WithField("request_id", GetRequestID(r.Context())).
			Debug("Request timing")
	})
}
