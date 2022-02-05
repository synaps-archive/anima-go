package core

import (
	"encoding/json"

	"github.com/anima-protocol/anima-go/chains/ethereum"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func SignAttributes(anima *models.Protocol, issuingAuthorization *models.IssueAuthorization, resource *models.IssueResource) (map[string]*protocol.IssueAttribute, error) {
	signedAttributes := make(map[string]*protocol.IssueAttribute)

	issAuthorization, err := GetIssuingAuthorization(issuingAuthorization)
	if err != nil {
		return nil, err
	}

	for name := range resource.Attributes {
		value := resource.Attributes[name]
		issAttr := models.IssueAttribute{
			Attribute: models.IssueAttributeAttr{
				Name:  name,
				Value: crypto.Hash(value),
			},
			Resource: models.IssueAttributeResource{
				ID:        issAuthorization.Request.Resource,
				ExpiresAt: resource.ExpiresAt,
			},
			Owner:  issAuthorization.Owner,
			Issuer: issAuthorization.Issuer,
		}

		content, err := json.Marshal(issAttr)
		if err != nil {
			return nil, err
		}

		switch anima.Chain {
		case models.CHAIN_ETH:
			signature, err := ethereum.SignIssueAttribute(anima, &issAttr)
			if err != nil {
				return nil, err
			}

			signedAttributes[name] = &protocol.IssueAttribute{
				Value:     value,
				Content:   content,
				Signature: signature,
			}
		}
	}

	return signedAttributes, nil
}
