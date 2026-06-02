package jwt

import "Email-API/config"

type JWTData struct {
	Email string
	Phone uint
}

type JWTDeps struct {
	Secret *config.Config
}

type JWT struct {
	Secret string
}

func NewJWT(deps *JWTDeps) *JWT {
	return &JWT{
		Secret: deps.Secret.AuthToken.AuthToken,
	}
}
