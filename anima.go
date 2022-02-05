package animago

import (
	"fmt"

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

	signedAttributes, err := core.SignAttributes(anima, request.IssuingAuthorization, request.Resource)
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
			Content:   request.Verification.Content,
			Schema:    request.Verification.Schema,
			Signature: request.Verification.Signature,
		},
		IssuingAuthorization: &protocol.IssueAuthorization{
			Content:   request.IssuingAuthorization.Content,
			Schema:    request.IssuingAuthorization.Schema,
			Signature: request.IssuingAuthorization.Signature,
		},
	}

	return protocol.Issue(anima, req)
}

// Verify - Verify Request from Anima Protocol
func Verify(protocol *models.Protocol) error {
	fmt.Printf("> Verify Request\n")
	return nil
}

func main() {
	// Issue()
}
