package middlewares

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		logrus.WithFields(logrus.Fields{
			"Method": method,
			"Path":   path,
		}).Info("Логирование метода и функции.")
		next.ServeHTTP(w, r)
	})
}
