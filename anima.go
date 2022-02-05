package animago

import (
	"fmt"

	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/validators"
)

// Issue - Issue a new payload to Anima Protocol
func Issue(protocol *models.Protocol, request *models.IssueRequest) error {
	if err := validators.ValidateProtocol(protocol); err != nil {
		return err
	}

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
