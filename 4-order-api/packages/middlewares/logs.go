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
		time := time.Now().Format("2006-01-02 15:04:05")
		logrus.WithFields(logrus.Fields{
			"Method": method,
			"Path":   path,
			"Time":   time,
		}).Info("Логирование метода и функции.")
		next.ServeHTTP(w, r)
	})
}
