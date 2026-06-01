package middlewares

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		logrus.WithFields(logrus.Fields{
			"Method": method,
			"Path":   path,
			"Time":   time.Now(),
		}).Info("Логирование метода и функции.")
		next.ServeHTTP(w, r)
	})
}
