package middlewares

import (
	"Email-API/config"
	"Email-API/packages/jwt"
	"context"
	"net/http"
)

func IsAdminAuth(next http.Handler, config *config.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			http.Error(w, "get cookie problem", http.StatusUnauthorized)
			return
		}
		token := cookie.Value

		jwtWithSercet := jwt.NewJWT(config.AuthToken)
		claims, err := jwtWithSercet.AccessToken(token)
		if err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		userID := claims.UserID
		role := claims.Subject
		if role != "admin" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(context.Background(), "userId", userID)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func IsUserAuth(next http.Handler, config *config.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			http.Error(w, "get cookie problem", http.StatusUnauthorized)
			return
		}
		token := cookie.Value

		jwtWithSercet := jwt.NewJWT(config.AuthToken)
		claims, err := jwtWithSercet.AccessToken(token)
		if err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		userID := claims.UserID

		ctx := context.WithValue(context.Background(), "userId", userID)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
