package internal

import "math/rand/v2"

type EmailWithHash struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

var letters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890!?")

func NewEmailWithHash(email string) *EmailWithHash {
	return &EmailWithHash{
		Email: email,
		Hash:  GenerateHash(10),
	}
}

func GenerateHash(n int) string {
	runes := make([]rune, n)
	for i, _ := range runes {
		runes[i] = letters[rand.IntN(len(letters))]
	}
	hash := string(runes)
	return hash
}
