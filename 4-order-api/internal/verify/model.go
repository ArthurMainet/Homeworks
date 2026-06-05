package verify

import "math/rand/v2"

type EmailWithHash struct {
	Method           string `json:"method"`
	AdrressOrSession string `json:"email"`
	Hash             string `json:"hash"`
}

var letters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890!?")
var nums = []rune("0123456789")

func NewEmailWithHash(method, adrress string) *EmailWithHash {

	var hash string
	switch method {
	case "phone":
		hash = GenerateCode(4)
	case "email":
		hash = GenerateHash(10)
	default:
		return nil
	}

	return &EmailWithHash{
		Method:           method,
		AdrressOrSession: adrress,
		Hash:             hash,
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

func GenerateCode(n int) string {
	runes := make([]rune, n)
	for i, _ := range runes {
		runes[i] = letters[rand.IntN(len(letters))]
	}
	hash := string(runes)
	return hash
}
