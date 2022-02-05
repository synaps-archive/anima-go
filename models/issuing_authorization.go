package models

import (
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type IssuingAuthorization struct {
	Request IssuingAuthorizationRequest `json:"request"`
	Owner   AnimaOwner                  `json:"owner"`
	Issuer  AnimaIssuer                 `json:"issuer"`
}

type IssuingAuthorizationRequest struct {
	Resource    string                 `json:"resource"`
	RequestedAt string                 `json:"requested_at"`
	Fields      map[string]interface{} `json:"fields"`
}

/* Ethereum */
type EthIssuingAuthorization struct {
	Domain  apitypes.TypedDataDomain `json:"domain"`
	Message IssuingAuthorization     `json:"message"`
	Types   apitypes.Types           `json:"types"`
}
