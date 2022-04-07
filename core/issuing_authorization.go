package core

import (
	"encoding/base64"

	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func GetIssuingAuthorization(request *protocol.IssueRequest) (*models.IssuingAuthorization, error) {
	specs := request.Document.IssuingAuthorization.Specs
	encodedContent := request.Document.IssuingAuthorization.Content
	signature := request.Document.IssuingAuthorization.Signature

	content, err := base64.StdEncoding.DecodeString(encodedContent)
	if err != nil {
		return nil, err
	}

	issuingAuthorization, rErr := ExtractIssuingAuthorization[specs](content, signature)
	if rErr != nil {
		return nil, rErr
	}

	return issuingAuthorization, nil
}
