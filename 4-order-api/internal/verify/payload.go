package verify

type SessionVerifRequest struct {
	Session string `json:"session"`
	Code    string `json:"code"`
}
