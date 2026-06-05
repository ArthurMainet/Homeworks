package middlewares

import (
	"Email-API/config"
	"Email-API/packages/jwt"
	"net/http"
)

func IsAuth(next http.Handler, config *config.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			http.Error(w, "take cookie problem", http.StatusUnauthorized)
			return
		}
		token := cookie.Value

		jwtWithSercet := jwt.NewJWT(config.AuthToken)
		claims, err := jwtWithSercet.AccessToken(token)
		if err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		// //		email := claims.Email
		role := claims.Subject
		if role != "admin" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		// 		ctx := context.WithValue(r.Context(), "role", role)
		// 		req := r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
