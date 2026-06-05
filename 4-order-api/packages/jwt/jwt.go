package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type JWT struct {
	Secret string
}

type Claims struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) GenerateToken(email, role, phone, secret string) (string, error) {

	claims := Claims{
		Email: email,
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "my-api",
			Subject:   role,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(360 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", errors.New("server error")
	}
	return tokenStr, nil
}

func (j *JWT) AccessToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, jwt.ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}
	if ok := token.Valid; !ok {
		return nil, jwt.ErrTokenInvalidAudience
	}

	claims := token.Claims.(*Claims)

	return claims, nil
}
