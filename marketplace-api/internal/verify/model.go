package verify

import (
	"math/rand/v2"
	"strconv"
)

type EmailWithHash struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type SessionWithCode struct {
	Phone   string `json:"phone"`
	Session string `json:"session"`
	Code    int    `json:"code"`
}

var letters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")
var nums = []rune("0123456789")

func NewEmailWithHash(adrress string) *EmailWithHash {
	return &EmailWithHash{
		Email: adrress,
		Hash:  GenerateHash(10),
	}
}
func NewSessionWithCode(session, phone string) *SessionWithCode {
	return &SessionWithCode{
		Phone:   phone,
		Session: session,
		Code:    GenerateCode(4),
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

func GenerateCode(n int) int {
	runes := make([]rune, n)
	for i, _ := range runes {
		runes[i] = nums[rand.IntN(len(nums))]
	}
	codeStr := string(runes)
	code, _ := strconv.Atoi(codeStr)
	return code
}
