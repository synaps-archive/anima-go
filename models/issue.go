package models

type IssueRequest struct {
	Resource             *IssueResource
	Verification         *IssueVerification
	IssuingAuthorization *IssueAuthorization
}

type IssueVerification struct {
	Schema  string `json:"schema"`
	Content string `json:"content"`
}

type IssueAuthorization struct {
	Schema    string `json:"schema"`
	Content   string `json:"content"`
	Signature string `json:"signature"`
}

type IssueResource struct {
	ID         string               `json:"id"`
	ExpiresAt  string               `json:"expires_at"`
	Attributes map[string]Attribute `json:"attributes"`
}

type Attribute struct {
	Value []byte `json:"value"`
	Type  string `json:"type"`
}

type IssueAttribute struct {
	Attribute IssueAttributeAttr     `json:"attribute"`
	Resource  IssueAttributeResource `json:"resource"`
	Owner     AnimaOwner             `json:"owner"`
	Issuer    AnimaIssuer            `json:"issuer"`
}

type IssueAttributeResource struct {
	ID        string `json:"id"`
	ExpiresAt string `json:"expires_at"`
}

type IssueAttributeAttr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}
