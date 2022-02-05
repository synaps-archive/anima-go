package animago

import (
	"fmt"

	"github.com/anima-protocol/anima-go/core"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/validators"
)

// Issue - Issue a new payload to Anima Protocol
func Issue(anima *models.Protocol, request *models.IssueRequest) error {
	if err := validators.ValidateProtocol(anima); err != nil {
		return err
	}

	signedAttributes, err := core.SignAttributes(anima, request.IssuingAuthorization, request.Resource)
	if err != nil {
		return err
	}

	// Send to server

	return nil
}

// Verify - Verify Request from Anima Protocol
func Verify(protocol *models.Protocol) error {
	fmt.Printf("> Verify Request\n")
	return nil
}

func main() {
	// Issue()
}
