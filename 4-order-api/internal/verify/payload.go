package verify

type SessionVerifRequest struct {
	Session string `json:"session"`
	Code    int    `json:"code"`
}
