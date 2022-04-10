package main

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/anima-protocol/anima-go/core"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
	"github.com/anima-protocol/anima-go/validators"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

// Issue - Issue new credential to Anima Protocol
func Issue(anima *models.Protocol, issuer *protocol.AnimaIssuer, request *protocol.IssueRequest) error {
	if err := validators.ValidateProtocol(anima); err != nil {
		return err
	}

	request, err := core.SignIssuing(anima, issuer, request, anima.SigningFunc)
	if err != nil {
		return err
	}

	return protocol.Issue(anima, request)
}

// Verify - Verify Sharing Request from Anima Protocol
func Verify(anima *models.Protocol, request *protocol.VerifyRequest) (*protocol.VerifyResponse, error) {
	if err := validators.ValidateProtocol(anima); err != nil {
		return &protocol.VerifyResponse{}, err
	}

	return protocol.Verify(anima, request)
}

// RegisterVerifier - Register Verifier on Anima Protocol
func RegisterVerifier(anima *models.Protocol, request *models.RegisterVerifierRequest) (*protocol.RegisterVerifierResponse, error) {
	if err := validators.ValidateProtocol(anima); err != nil {
		return &protocol.RegisterVerifierResponse{}, err
	}

	req := &protocol.RegisterVerifierRequest{
		Id:            request.Id,
		PublicAddress: request.PublicAddress,
		Chain:         request.Chain,
	}

	res, err := protocol.RegisterVerifier(anima, req)
	if err != nil {
		return &protocol.RegisterVerifierResponse{}, err
	}

	return res, nil
}

type EthereumKey struct {
	Key        string
	EthAddress string
}

func (ek *EthereumKey) EthSign(digest []byte) (string, error) {
	pv, err := ethCrypto.HexToECDSA(ek.Key)
	if err != nil {
		return "", err
	}

	sig, err := ethCrypto.Sign(digest, pv)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sig), nil
}

// /* Issuer */
func main() {
	ek := EthereumKey{
		Key:        "d78a9d14865872880a36cf4f22831e32a00e0b09e8b436fbf9c148cc54e9e52b",
		EthAddress: "0x03ae90f3899B0694Dc3c45E01543c8480bfB95FC",
	}

	firstname := "Riwan"
	lastname := "Lasmi"
	birthDate := "1998-06-08"
	documentCountry := "FRA"
	documentNumber := "823942834"
	documentExpiration := "2021-12-01"
	nationality := "FRA"
	livenessPhoto := []byte{}
	passportPage := []byte{}

	issuer := &protocol.AnimaIssuer{
		Id:            "anima:issuer:synaps@1.0.0",
		PublicAddress: "0x03ae90f3899B0694Dc3c45E01543c8480bfB95FC",
		Chain:         "ETH",
	}

	document := &protocol.IssDocument{
		Specs:     "anima:specs:document/passport@1.0.0",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		Attributes: map[string]*protocol.IssDocumentAttribute{
			"firstname": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:  "firstname",
					Type:  "string",
					Value: firstname,
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
			"lastname": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:  "lastname",
					Type:  "string",
					Value: lastname,
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
			"birth_date": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:   "birth_date",
					Type:   "string",
					Format: "date_yyyymmdd",
					Value:  birthDate,
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
			"document_country": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:   "document_country",
					Type:   "string",
					Format: "country_iso3",
					Value:  documentCountry,
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
			"document_number": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:  "document_number",
					Type:  "string",
					Value: documentNumber,
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
			"document_expiration": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:   "document_expiration",
					Type:   "string",
					Format: "date_yyyymmdd",
					Value:  documentExpiration,
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
			"nationality": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:   "nationality",
					Type:   "string",
					Format: "country_iso3",
					Value:  nationality,
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
			"liveness_photo": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:  "liveness_photo",
					Type:  "file",
					Value: crypto.Hash(livenessPhoto),
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
			"passport_page": {
				Content: &protocol.IssDocumentAttributeContent{
					Name:  "passport_page",
					Type:  "file",
					Value: crypto.Hash(passportPage),
				},
				Credential: &protocol.IssDocumentAttributeCredential{},
			},
		},
		Authorization: &protocol.IssAuthorization{
			Specs:     "anima:specs:issuing/authorization/eip712@1.0.0",
			Content:   "eyJkb21haW4iOnsibmFtZSI6ImFuaW1hIiwidmVyc2lvbiI6IjEuMCIsImNoYWluSWQiOiIweDAxIn0sIm1lc3NhZ2UiOnsic3BlY3MiOiJhbmltYTpzcGVjczpkb2N1bWVudC9wYXNzcG9ydEAxLjAuMCIsInJlcXVlc3RlZF9hdCI6MTY0OTUyNDE4NywiZmllbGRzIjp7InBhc3Nwb3J0X3BhZ2UiOiJhY2EzMjFjNTVlNDRmM2FhYzZhMjE5Y2RkMTM0NDQ5MDcwNjVkNWIxZTJlODFmNjM0NTg4YjY4MjgwMmVmMTdlIiwibGl2ZW5lc3NfcGhvdG8iOiJiMjQwY2NjZDc2NzIyYjkxODliNDk5YjU0OGIxMDMwNWUxYWRiMTEwYWE3MGU3ODBmYzI3NjFmYjk0ODY1MWMxIn0sImF0dHJpYnV0ZXMiOnsiZmlyc3RuYW1lIjp0cnVlLCJsYXN0bmFtZSI6dHJ1ZSwiYmlydGhfZGF0ZSI6dHJ1ZSwiZG9jdW1lbnRfY291bnRyeSI6dHJ1ZSwiZG9jdW1lbnRfbnVtYmVyIjp0cnVlLCJkb2N1bWVudF9leHBpcmF0aW9uIjp0cnVlLCJuYXRpb25hbGl0eSI6dHJ1ZSwibGl2ZW5lc3NfcGhvdG8iOnRydWUsInBhc3Nwb3J0X3BhZ2UiOnRydWV9LCJvd25lciI6eyJpZCI6ImFuaW1hOm93bmVyOjB4RjA5OTY3NTBlNzQ1Qjk1ODY0YzZjYzdjM2JlNDk4REIzRDg4MzBDNSIsImNoYWluIjoiRVRIIiwid2FsbGV0IjoiTUVUQU1BU0siLCJwdWJsaWNfYWRkcmVzcyI6IjB4RjA5OTY3NTBlNzQ1Qjk1ODY0YzZjYzdjM2JlNDk4REIzRDg4MzBDNSIsInB1YmxpY19rZXlfZW5jcnlwdGlvbiI6Im9NZ1lJRlB1L0ZLOXNwcjh2Y3JyRi85WXRsa2JhdXlYVlo1ZERMZmJZQzQ9In0sImlzc3VlciI6eyJpZCI6ImFuaW1hOmlzc3VlcjpzeW5hcHNAMS4wLjAiLCJwdWJsaWNfYWRkcmVzcyI6IjB4MDNhZTkwZjM4OTlCMDY5NERjM2M0NUUwMTU0M2M4NDgwYmZCOTVGQyIsImNoYWluIjoiRVRIIn19LCJwcmltYXJ5VHlwZSI6Ik1haW4iLCJ0eXBlcyI6eyJNYWluIjpbeyJuYW1lIjoiaXNzdWVyIiwidHlwZSI6Iklzc3VlciJ9LHsibmFtZSI6Im93bmVyIiwidHlwZSI6Ik93bmVyIn0seyJuYW1lIjoic3BlY3MiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoicmVxdWVzdGVkX2F0IiwidHlwZSI6InVpbnQ2NCJ9LHsibmFtZSI6ImF0dHJpYnV0ZXMiLCJ0eXBlIjoiQXR0cmlidXRlcyJ9LHsibmFtZSI6ImZpZWxkcyIsInR5cGUiOiJGaWVsZHMifV0sIkZpZWxkcyI6W3sibmFtZSI6InBhc3Nwb3J0X3BhZ2UiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoibGl2ZW5lc3NfcGhvdG8iLCJ0eXBlIjoic3RyaW5nIn1dLCJBdHRyaWJ1dGVzIjpbeyJuYW1lIjoiZmlyc3RuYW1lIiwidHlwZSI6ImJvb2wifSx7Im5hbWUiOiJsYXN0bmFtZSIsInR5cGUiOiJib29sIn0seyJuYW1lIjoiYmlydGhfZGF0ZSIsInR5cGUiOiJib29sIn0seyJuYW1lIjoiZG9jdW1lbnRfY291bnRyeSIsInR5cGUiOiJib29sIn0seyJuYW1lIjoiZG9jdW1lbnRfbnVtYmVyIiwidHlwZSI6ImJvb2wifSx7Im5hbWUiOiJkb2N1bWVudF9leHBpcmF0aW9uIiwidHlwZSI6ImJvb2wifSx7Im5hbWUiOiJuYXRpb25hbGl0eSIsInR5cGUiOiJib29sIn0seyJuYW1lIjoibGl2ZW5lc3NfcGhvdG8iLCJ0eXBlIjoiYm9vbCJ9LHsibmFtZSI6InBhc3Nwb3J0X3BhZ2UiLCJ0eXBlIjoiYm9vbCJ9XSwiT3duZXIiOlt7Im5hbWUiOiJpZCIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJwdWJsaWNfYWRkcmVzcyIsInR5cGUiOiJhZGRyZXNzIn0seyJuYW1lIjoiY2hhaW4iLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoid2FsbGV0IiwidHlwZSI6InN0cmluZyJ9LHsibmFtZSI6InB1YmxpY19rZXlfZW5jcnlwdGlvbiIsInR5cGUiOiJzdHJpbmcifV0sIklzc3VlciI6W3sibmFtZSI6ImlkIiwidHlwZSI6InN0cmluZyJ9LHsibmFtZSI6InB1YmxpY19hZGRyZXNzIiwidHlwZSI6ImFkZHJlc3MifSx7Im5hbWUiOiJjaGFpbiIsInR5cGUiOiJzdHJpbmcifV0sIkVJUDcxMkRvbWFpbiI6W3sibmFtZSI6Im5hbWUiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoiY2hhaW5JZCIsInR5cGUiOiJ1aW50MjU2In0seyJuYW1lIjoidmVyc2lvbiIsInR5cGUiOiJzdHJpbmcifV19fQ==",
			Signature: "0x1c5cff8f72378464154ebef738f90e848e7d978e335d28c37d96717afe6e1ec834d2c29d061a92998336988e93b0836b949fb7b80316ae308a92d60f89f8cfa91c",
		},
	}

	proof := &protocol.IssProof{
		Specs:   "anima:issuer:synaps/proofs/passport@1.0.0",
		Content: "ewogICAgInByb29mIjogInN5bmFwc19pZF9vayIKfQ==",
	}

	issuingRequest := &protocol.IssueRequest{
		Document: document,
		Proof:    proof,
		Attributes: map[string]*protocol.IssAttribute{
			"firstname": {
				Value:   []byte(firstname),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.HashStr(firstname),
						},
					},
				},
			},
			"lastname": {
				Value:   []byte(lastname),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.HashStr(lastname),
						},
					},
				},
			},
			"birth_date": {
				Value:   []byte(birthDate),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.HashStr(birthDate),
						},
					},
				},
			},
			"nationality": {
				Value:   []byte(nationality),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.HashStr(nationality),
						},
					},
				},
			},
			"document_country": {
				Value:   []byte(documentCountry),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.HashStr(documentCountry),
						},
					},
				},
			},
			"document_expiration": {
				Value:   []byte(documentExpiration),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.HashStr(documentExpiration),
						},
					},
				},
			},
			"document_number": {
				Value:   []byte(documentNumber),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.HashStr(documentNumber),
						},
					},
				},
			},
			"passport_page": {
				Value:   []byte(passportPage),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.Hash(passportPage),
						},
					},
				},
			},
			"liveness_photo": {
				Value:   []byte(livenessPhoto),
				Content: &protocol.IssDocumentAttributeContent{},
				Credential: &protocol.IssAttributeCredential{
					Content: &protocol.IssAttributeCredentialContent{
						Attribute: &protocol.IssAttributeCredentialContentAttribute{
							Hash: crypto.Hash(livenessPhoto),
						},
					},
				},
			},
		},
	}

	anima := &models.Protocol{
		Network:     "localhost:8100",
		Chain:       models.CHAIN_ETH,
		Secure:      false,
		SigningFunc: ek.EthSign,
	}

	err := Issue(anima, issuer, issuingRequest)
	if err != nil {
		fmt.Printf("-> err = %v\n", err.Error())
	}
}

// /* Verify */
// func main() {
// 	ek := EthereumKey{
// 		Key:        "5c18b87c4ae6c645da7be527d2d959c7213af2a14100b8dfbf802649658bb725",
// 		EthAddress: "0xcE07939863Bcbf38c4FaAf7604680924E87a6115",
// 	}

// 	request := &protocol.VerifyRequest{
// 		Authorization: &protocol.SharingAuthorization{
// 			Specs:     "anima:specs:sharing/authorization/eip712@1.0.0",
// 			Content:   "eyJkb21haW4iOnsibmFtZSI6ImFuaW1hIiwidmVyc2lvbiI6IjEuMCIsImNoYWluSWQiOiIxIn0sIm1lc3NhZ2UiOnsic3BlY3MiOiJhbmltYTpzcGVjczpzaGFyaW5nL2F1dGhvcml6YXRpb25AMS4wLjAiLCJzaGFyZWRfYXQiOjE2NDk1MDUxNTYsImF0dHJpYnV0ZXMiOnsiYmlydGhfZGF0ZSI6ImFuaW1hOmNyZWRlbnRpYWw6MmI1MDkxYmI5N2IyMzFlZmY4YTkyZTI5NTQ1ZWQ5ZDc5NDMzM2U0YTNkZGIzZjM5ZTYwNGRiMzIxMDk0MmNhYiIsImRvY3VtZW50X2NvdW50cnkiOiJhbmltYTpjcmVkZW50aWFsOjFjNDU0ZmFjM2UyYjVhZjY4NjExODcwYzQxZDA2MjRjNDRmNjg4YmMxYjEzMDJmYWM2ZWIwYWFlMWJkOGI4OWEiLCJkb2N1bWVudF9leHBpcmF0aW9uIjoiYW5pbWE6Y3JlZGVudGlhbDoyYzcwYmUyZjhlOTJhODQ3Y2ZiYTNlMzc4MTdkZmNmMTQzZjFlZWRiYzc3YzUwNDJiMzc2MDRlNGJiNTBhZWY4IiwiZG9jdW1lbnRfbnVtYmVyIjoiYW5pbWE6Y3JlZGVudGlhbDoxNzExY2NkZDEyNDA0NmIxODM0OGI0ZWZlZGUyZTlkY2ZiM2E4N2UyMzc3YjliNDg4MDMzM2M1MTE3N2RhY2ZmIiwiZmlyc3RuYW1lIjoiYW5pbWE6Y3JlZGVudGlhbDowNjA5Y2ZkMGE3OWIyNDY4ZGZkOTA1NTg1YTgwZjcyY2Q0YzFmMGEzZDdkNjkyZmJmZDI3ZDkxYTg5OTc1YTgzIiwibGFzdG5hbWUiOiJhbmltYTpjcmVkZW50aWFsOjE4NWE0N2NiMWY3N2VhM2U0N2NhYjI4NWE5ZjNmZjAzMWE4YzY2MzlmNmEzOTc2MTlmMWY4NjE3ZTMyYTc2ZjUiLCJsaXZlbmVzc19waG90byI6ImFuaW1hOmNyZWRlbnRpYWw6NDBkZGY2OTY3YzNkYzAzOGFjZDhlMzYyYjBmNDk0N2QxZjEwNTQ5OTk1ODZjYzA5MmYxYmIyYzljMjJlNTA3ZiIsIm5hdGlvbmFsaXR5IjoiYW5pbWE6Y3JlZGVudGlhbDo3MTRkYzM2NmE0NjkwNDEzMjY4MmZmOWUyZjJhZTJlNjMxNjJlN2ZkMzViNWI0YzRjZjQwOTdjYTk0OTgyN2I3IiwicGFzc3BvcnRfcGFnZSI6ImFuaW1hOmNyZWRlbnRpYWw6NzcyMzM5ZWU5MzRlYjBjM2JlODY3YWQzOGVmYTcwOTkxMWZhMGM5ODQ4NWI1OTNhNTk0MjdkNDZiZDFkZTRiMiJ9LCJvd25lciI6eyJpZCI6ImFuaW1hOm93bmVyOjB4RjA5OTY3NTBlNzQ1Qjk1ODY0YzZjYzdjM2JlNDk4REIzRDg4MzBDNSIsImNoYWluIjoiRVRIIiwid2FsbGV0IjoiTUVUQU1BU0siLCJwdWJsaWNfYWRkcmVzcyI6IjB4RjA5OTY3NTBlNzQ1Qjk1ODY0YzZjYzdjM2JlNDk4REIzRDg4MzBDNSJ9LCJ2ZXJpZmllciI6eyJpZCI6ImFuaW1hOnZlcmlmaWVyOnN5bmFwc19iaW5hbmNlQDEuMC4wIiwicHVibGljX2FkZHJlc3MiOiIweGNFMDc5Mzk4NjNCY2JmMzhjNEZhQWY3NjA0NjgwOTI0RTg3YTYxMTUiLCJjaGFpbiI6IkVUSCJ9fSwicHJpbWFyeVR5cGUiOiJNYWluIiwidHlwZXMiOnsiTWFpbiI6W3sibmFtZSI6InZlcmlmaWVyIiwidHlwZSI6IlZlcmlmaWVyIn0seyJuYW1lIjoib3duZXIiLCJ0eXBlIjoiT3duZXIifSx7Im5hbWUiOiJzcGVjcyIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJzaGFyZWRfYXQiLCJ0eXBlIjoidWludDY0In0seyJuYW1lIjoiYXR0cmlidXRlcyIsInR5cGUiOiJBdHRyaWJ1dGVzIn1dLCJBdHRyaWJ1dGVzIjpbeyJuYW1lIjoiZmlyc3RuYW1lIiwidHlwZSI6InN0cmluZyJ9LHsibmFtZSI6Imxhc3RuYW1lIiwidHlwZSI6InN0cmluZyJ9LHsibmFtZSI6ImJpcnRoX2RhdGUiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoiZG9jdW1lbnRfY291bnRyeSIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJkb2N1bWVudF9udW1iZXIiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoiZG9jdW1lbnRfZXhwaXJhdGlvbiIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJuYXRpb25hbGl0eSIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJsaXZlbmVzc19waG90byIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJwYXNzcG9ydF9wYWdlIiwidHlwZSI6InN0cmluZyJ9XSwiT3duZXIiOlt7Im5hbWUiOiJpZCIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJwdWJsaWNfYWRkcmVzcyIsInR5cGUiOiJhZGRyZXNzIn0seyJuYW1lIjoiY2hhaW4iLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoid2FsbGV0IiwidHlwZSI6InN0cmluZyJ9XSwiVmVyaWZpZXIiOlt7Im5hbWUiOiJpZCIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJwdWJsaWNfYWRkcmVzcyIsInR5cGUiOiJhZGRyZXNzIn0seyJuYW1lIjoiY2hhaW4iLCJ0eXBlIjoic3RyaW5nIn1dLCJFSVA3MTJEb21haW4iOlt7Im5hbWUiOiJuYW1lIiwidHlwZSI6InN0cmluZyJ9LHsibmFtZSI6ImNoYWluSWQiLCJ0eXBlIjoidWludDI1NiJ9LHsibmFtZSI6InZlcnNpb24iLCJ0eXBlIjoic3RyaW5nIn1dfX0=",
// 			Signature: "0xcd6a67f1edafd2c0a698c761d76bce55c4858fc2834b4819c93f3147ad67e994721687588164afe5a36586355f299c4b5cb405940f1280f4b386e5949f67c8801c",
// 		},
// 	}

// 	anima := &models.Protocol{
// 		Network:     "localhost:8100",
// 		Chain:       models.CHAIN_ETH,
// 		Secure:      false,
// 		SigningFunc: ek.EthSign,
// 	}

// 	res, err := Verify(anima, request)
// 	if err != nil {
// 		fmt.Printf("-> err = %v\n", err.Error())
// 	}

// 	b, _ := json.Marshal(res)
// 	ioutil.WriteFile("result.json", b, 0755)

// 	fmt.Printf("-> res = %v\n", res)
// }
