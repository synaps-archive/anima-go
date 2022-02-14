package core

import (
	"encoding/base64"

	"github.com/anima-protocol/anima-go/chains/ethereum"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func SignAttributes(anima *models.Protocol, issuingAuthorization *models.IssueAuthorization, resource *models.IssueResource, signingFunc func([]byte) (string, error)) (map[string]*protocol.IssueAttribute, error) {
	signedAttributes := make(map[string]*protocol.IssueAttribute)

	issAuthorization, err := GetIssuingAuthorization(issuingAuthorization)
	if err != nil {
		return nil, err
	}

	for name := range resource.Attributes {
		attr := resource.Attributes[name]
		issAttr := models.IssueAttribute{
			Resource: models.IssueAttributeResource{
				ID:        issAuthorization.Request.Resource,
				ExpiresAt: resource.ExpiresAt,
			},
			Attribute: models.IssueAttributeAttr{
				Name:  name,
				Value: crypto.Hash(attr.Value),
				Type:  attr.Type,
			},
			Owner: models.AnimaOwner{
				ID:            issAuthorization.Owner.ID,
				PublicAddress: issAuthorization.Owner.PublicAddress,
				Chain:         issAuthorization.Owner.Chain,
			},
			Issuer: models.AnimaIssuer{
				ID:            issAuthorization.Issuer.ID,
				PublicAddress: issAuthorization.Issuer.PublicAddress,
				Chain:         issAuthorization.Issuer.Chain,
			},
		}

		switch anima.Chain {
		case models.CHAIN_ETH:
			signedContent, signature, err := ethereum.SignIssueAttribute(anima, &issAttr, signingFunc)
			if err != nil {
				return nil, err
			}

			content := base64.StdEncoding.EncodeToString(signedContent)

			signedAttributes[name] = &protocol.IssueAttribute{
				Value:     attr.Value,
				Content:   content,
				Signature: signature,
			}
		}
	}

	return signedAttributes, nil
}
