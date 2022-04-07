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
func Verify(anima *models.Protocol, request *models.VerifyRequest) (*protocol.VerifyResponse, error) {
	if err := validators.ValidateProtocol(anima); err != nil {
		return &protocol.VerifyResponse{}, err
	}

	req := &protocol.VerifyRequest{
		Schema:    request.Schema,
		Content:   request.Content,
		Signature: request.Signature,
	}

	res, err := protocol.Verify(anima, req)
	if err != nil {
		return &protocol.VerifyResponse{}, err
	}

	return res, nil
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
	documentExpiration := "2023-01-2021"
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
		Attributes: map[string]*protocol.IssAttribute{
			"firstname": {
				Name:  "firstname",
				Type:  "string",
				Value: firstname,
			},
			"lastname": {
				Name:  "lastname",
				Type:  "string",
				Value: lastname,
			},
			"birth_date": {
				Name:  "birth_date",
				Type:  "string",
				Value: birthDate,
			},
			"document_country": {
				Name:   "document_country",
				Type:   "string",
				Format: "country_iso3",
				Value:  documentCountry,
			},
			"document_number": {
				Name:  "document_number",
				Type:  "string",
				Value: documentNumber,
			},
			"document_expiration": {
				Name:   "document_expiration",
				Type:   "string",
				Format: "date_yyyymmdd",
				Value:  documentExpiration,
			},
			"nationality": {
				Name:   "nationality",
				Type:   "string",
				Format: "country_iso3",
				Value:  nationality,
			},
			"liveness_photo": {
				Name:  "liveness_photo",
				Type:  "file",
				Value: crypto.Hash(livenessPhoto),
			},
			"passport_page": {
				Name:  "passport_page",
				Type:  "file",
				Value: crypto.Hash(passportPage),
			},
		},
		IssuingAuthorization: &protocol.IssAuthorization{
			Specs:     "anima:specs:issuing/authorization/eip712@1.0.0",
			Content:   "eyJkb21haW4iOnsibmFtZSI6ImFuaW1hIiwidmVyc2lvbiI6IjEuMCIsImNoYWluSWQiOiIweDAxIn0sIm1lc3NhZ2UiOnsic3BlY3MiOiJhbmltYTpzcGVjczppc3N1aW5nL2F1dGhvcml6YXRpb25AMS4wLjAiLCJyZXF1ZXN0ZWRfYXQiOjEyMzcyODI5MjMxLCJmaWVsZHMiOnsicGFzc3BvcnRfcGFnZSI6ImFjYTMyMWM1NWU0NGYzYWFjNmEyMTljZGQxMzQ0NDkwNzA2NWQ1YjFlMmU4MWY2MzQ1ODhiNjgyODAyZWYxN2UiLCJsaXZlbmVzc19waG90byI6ImIyNDBjY2NkNzY3MjJiOTE4OWI0OTliNTQ4YjEwMzA1ZTFhZGIxMTBhYTcwZTc4MGZjMjc2MWZiOTQ4NjUxYzEifSwiYXR0cmlidXRlcyI6eyJmaXJzdG5hbWUiOnRydWUsImxhc3RuYW1lIjp0cnVlLCJiaXJ0aF9kYXRlIjp0cnVlLCJkb2N1bWVudF9jb3VudHJ5Ijp0cnVlLCJkb2N1bWVudF9udW1iZXIiOnRydWUsImRvY3VtZW50X2V4cGlyYXRpb24iOnRydWUsIm5hdGlvbmFsaXR5Ijp0cnVlLCJsaXZlbmVzc19waG90byI6dHJ1ZSwicGFzc3BvcnRfcGFnZSI6dHJ1ZX0sIm93bmVyIjp7ImlkIjoiYW5pbWE6b3duZXI6MHhGMDk5Njc1MGU3NDVCOTU4NjRjNmNjN2MzYmU0OThEQjNEODgzMEM1IiwiY2hhaW4iOiJFVEgiLCJ3YWxsZXQiOiJNRVRBTUFTSyIsInB1YmxpY19hZGRyZXNzIjoiMHhGMDk5Njc1MGU3NDVCOTU4NjRjNmNjN2MzYmU0OThEQjNEODgzMEM1IiwicHVibGljX2tleV9lbmNyeXB0aW9uIjoib01nWUlGUHUvRks5c3ByOHZjcnJGLzlZdGxrYmF1eVhWWjVkRExmYllDND0ifSwiaXNzdWVyIjp7ImlkIjoiYW5pbWE6aXNzdWVyOnN5bmFwc0AxLjAuMCIsInB1YmxpY19hZGRyZXNzIjoiMHgwM2FlOTBmMzg5OUIwNjk0RGMzYzQ1RTAxNTQzYzg0ODBiZkI5NUZDIiwiY2hhaW4iOiJFVEgifX0sInByaW1hcnlUeXBlIjoiTWFpbiIsInR5cGVzIjp7Ik1haW4iOlt7Im5hbWUiOiJpc3N1ZXIiLCJ0eXBlIjoiSXNzdWVyIn0seyJuYW1lIjoib3duZXIiLCJ0eXBlIjoiT3duZXIifSx7Im5hbWUiOiJzcGVjcyIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJyZXF1ZXN0ZWRfYXQiLCJ0eXBlIjoidWludDY0In0seyJuYW1lIjoiYXR0cmlidXRlcyIsInR5cGUiOiJBdHRyaWJ1dGVzIn0seyJuYW1lIjoiZmllbGRzIiwidHlwZSI6IkZpZWxkcyJ9XSwiRmllbGRzIjpbeyJuYW1lIjoicGFzc3BvcnRfcGFnZSIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJsaXZlbmVzc19waG90byIsInR5cGUiOiJzdHJpbmcifV0sIkF0dHJpYnV0ZXMiOlt7Im5hbWUiOiJmaXJzdG5hbWUiLCJ0eXBlIjoiYm9vbCJ9LHsibmFtZSI6Imxhc3RuYW1lIiwidHlwZSI6ImJvb2wifSx7Im5hbWUiOiJiaXJ0aF9kYXRlIiwidHlwZSI6ImJvb2wifSx7Im5hbWUiOiJkb2N1bWVudF9jb3VudHJ5IiwidHlwZSI6ImJvb2wifSx7Im5hbWUiOiJkb2N1bWVudF9udW1iZXIiLCJ0eXBlIjoiYm9vbCJ9LHsibmFtZSI6ImRvY3VtZW50X2V4cGlyYXRpb24iLCJ0eXBlIjoiYm9vbCJ9LHsibmFtZSI6Im5hdGlvbmFsaXR5IiwidHlwZSI6ImJvb2wifSx7Im5hbWUiOiJsaXZlbmVzc19waG90byIsInR5cGUiOiJib29sIn0seyJuYW1lIjoicGFzc3BvcnRfcGFnZSIsInR5cGUiOiJib29sIn1dLCJPd25lciI6W3sibmFtZSI6ImlkIiwidHlwZSI6InN0cmluZyJ9LHsibmFtZSI6InB1YmxpY19hZGRyZXNzIiwidHlwZSI6ImFkZHJlc3MifSx7Im5hbWUiOiJjaGFpbiIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJ3YWxsZXQiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoicHVibGljX2tleV9lbmNyeXB0aW9uIiwidHlwZSI6InN0cmluZyJ9XSwiSXNzdWVyIjpbeyJuYW1lIjoiaWQiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoicHVibGljX2FkZHJlc3MiLCJ0eXBlIjoiYWRkcmVzcyJ9LHsibmFtZSI6ImNoYWluIiwidHlwZSI6InN0cmluZyJ9XSwiRUlQNzEyRG9tYWluIjpbeyJuYW1lIjoibmFtZSIsInR5cGUiOiJzdHJpbmcifSx7Im5hbWUiOiJjaGFpbklkIiwidHlwZSI6InVpbnQyNTYifSx7Im5hbWUiOiJ2ZXJzaW9uIiwidHlwZSI6InN0cmluZyJ9XX19",
			Signature: "0xfa9eeb49fc23faba3e9335a6d6863c13ba84e85c19abf66cf1f1bc5597dc13000ac0729944de93cc0a069adf3e0deeb1f467b00d06ac0b28b3c867f69810c3e81c",
		},
	}

	proof := &protocol.IssProof{
		Specs:   "anima:issuer:synaps/proofs/passport@1.0.0",
		Content: "ewogICAgInByb29mIjogInN5bmFwc19pZF9vayIKfQ==",
	}

	issuingRequest := &protocol.IssueRequest{
		Document: document,
		Proof:    proof,
		Attributes: map[string]*protocol.IssAttributeCredential{
			"firstname": {
				Value: []byte(firstname),
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "firstname",
						Type:      "string",
						Hash:      crypto.HashStr(firstname),
					},
				},
			},
			"lastname": {
				Value: []byte(lastname),
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "lastname",
						Type:      "string",
						Hash:      crypto.HashStr(lastname),
					},
				},
			},
			"birth_date": {
				Value: []byte(birthDate),
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "birth_date",
						Type:      "string",
						Hash:      crypto.HashStr(birthDate),
					},
				},
			},
			"nationality": {
				Value: []byte(nationality),
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "nationality",
						Type:      "string",
						Hash:      crypto.HashStr(nationality),
					},
				},
			},
			"document_country": {
				Value: []byte(documentCountry),
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "document_country",
						Type:      "string",
						Hash:      crypto.HashStr(documentCountry),
					},
				},
			},
			"document_expiration": {
				Value: []byte(documentExpiration),
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "document_expiration",
						Type:      "string",
						Hash:      crypto.HashStr(documentExpiration),
					},
				},
			},
			"document_number": {
				Value: []byte(documentNumber),
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "document_number",
						Type:      "string",
						Hash:      crypto.HashStr(documentNumber),
					},
				},
			},
			"passport_page": {
				Value: passportPage,
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "passport_page",
						Type:      "string",
						Hash:      crypto.Hash(passportPage),
					},
				},
			},
			"liveness_photo": {
				Value: livenessPhoto,
				Credential: &protocol.IssCredential{
					Content: &protocol.IssCredentialContent{
						ExpiresAt: time.Now().Unix(),
						Name:      "liveness_photo",
						Type:      "string",
						Hash:      crypto.Hash(livenessPhoto),
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
