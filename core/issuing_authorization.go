package core

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/anima-protocol/anima-go/chains/ethereum"
	"github.com/anima-protocol/anima-go/models"
)

func GetIssuingAuthorization(issuingAuthorization *models.IssueAuthorization) (models.IssuingAuthorization, error) {
	switch issuingAuthorization.Schema {
	case "anima:schema:eth_issuing_authorization":
		issBytes, err := base64.StdEncoding.DecodeString(issuingAuthorization.Content)
		if err != nil {
			return models.IssuingAuthorization{}, err
		}

		ethIssAuthorization := models.EthIssuingAuthorization{}
		err = json.Unmarshal(issBytes, &ethIssAuthorization)
		if err != nil {
			return models.IssuingAuthorization{}, err
		}

		valid, err := ethereum.VerifySignature(ethIssAuthorization.Message.Owner.PublicAddress, issBytes, issuingAuthorization.Signature)
		if err != nil {
			return models.IssuingAuthorization{}, err
		}

		if !valid {
			return models.IssuingAuthorization{}, fmt.Errorf("invalid issuing authorization signature")
		}
		return ethIssAuthorization.Message, nil
	}
	return models.IssuingAuthorization{}, fmt.Errorf("invalid issuing authorization")
}
