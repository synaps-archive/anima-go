package models

type IssueRequest struct {
	Resource             *IssueResource
	Verification         *IssueVerification
	IssuingAuthorization *IssueAuthorization
}

type IssueVerification struct {
	Schema    string
	Content   string
	Signature string
}

type IssueAuthorization struct {
	Schema    string
	Content   string
	Signature string
}

type IssueResource struct {
	ID         string
	ExpiresAt  int64
	Attributes map[string][]byte
}

type Protocol struct {
	PrivateKey string
	Network    string
	Chain      string
}
