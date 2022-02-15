package anima

import (
	"github.com/anima-protocol/anima-go/chains/ethereum"
	"github.com/anima-protocol/anima-go/core"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
	"github.com/anima-protocol/anima-go/validators"
)

// Issue - Issue new credential to Anima Protocol
func Issue(anima *models.Protocol, request *models.IssueRequest) error {
	if err := validators.ValidateProtocol(anima); err != nil {
		return err
	}

	signedAttributes, err := core.SignAttributes(anima, request.IssuingAuthorization, request.Resource, anima.SigningFunc)
	if err != nil {
		return err
	}

	proofContent, proofSignature, err := ethereum.SignProof(anima, request.Verification.Content, anima.SigningFunc)
	if err != nil {
		return err
	}

	req := &protocol.IssueRequest{
		Resource: &protocol.IssueResource{
			Id:         request.Resource.ID,
			ExpiresAt:  request.Resource.ExpiresAt,
			Attributes: signedAttributes,
		},
		Verification: &protocol.IssueVerification{
			Content:   proofContent,
			Schema:    request.Verification.Schema,
			Signature: proofSignature,
		},
		IssuingAuthorization: &protocol.IssueAuthorization{
			Content:   request.IssuingAuthorization.Content,
			Schema:    request.IssuingAuthorization.Schema,
			Signature: request.IssuingAuthorization.Signature,
		},
	}

	return protocol.Issue(anima, req)
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
