package models

type VerifyRequest struct {
	Schema    string `json:"schema"`
	Content   string `json:"content"`
	Signature string `json:"signature"`
}
