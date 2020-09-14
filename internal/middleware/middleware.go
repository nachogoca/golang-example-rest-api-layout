package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithField("path", r.RequestURI).
			WithField("method", r.Method).
			Debug("Request arrived")
		next.ServeHTTP(w, r)
	})
}

// TODO Request ID middleware