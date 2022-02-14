package ethereum

import (
	"encoding/base64"
	"encoding/json"

	"github.com/anima-protocol/anima-go/models"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func SignIssueAttribute(protocol *models.Protocol, issAttr *models.IssueAttribute, signingFunc func([]byte) (string, error)) ([]byte, string, error) {
	message := make(map[string]interface{})

	b, err := json.Marshal(issAttr)
	if err != nil {
		return []byte{}, "", err
	}

	if err := json.Unmarshal(b, &message); err != nil {
		return []byte{}, "", err
	}

	content := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    models.PROTOCOL_NAME,
			Version: models.PROTOCOL_VERSION,
			ChainId: math.NewHexOrDecimal256(1),
		},
		PrimaryType: "Main",
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "version",
					Type: "string",
				},
			},
			"Issuer": []apitypes.Type{
				{
					Name: "id",
					Type: "string",
				},
				{
					Name: "public_address",
					Type: "string",
				},
				{
					Name: "chain",
					Type: "string",
				},
			},
			"Owner": []apitypes.Type{
				{
					Name: "id",
					Type: "string",
				},
				{
					Name: "public_address",
					Type: "string",
				},
				{
					Name: "chain",
					Type: "string",
				},
			},
			"Resource": []apitypes.Type{
				{
					Name: "id",
					Type: "string",
				},
				{
					Name: "expires_at",
					Type: "string",
				},
			},
			"Attribute": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "type",
					Type: "string",
				},
				{
					Name: "value",
					Type: "string",
				},
			},
			"Main": []apitypes.Type{
				{
					Name: "resource",
					Type: "Resource",
				},
				{
					Name: "owner",
					Type: "Owner",
				},
				{
					Name: "issuer",
					Type: "Issuer",
				},
				{
					Name: "attribute",
					Type: "Attribute",
				},
			},
		},
		Message: message,
	}

	c, err := json.Marshal(content)
	if err != nil {
		return []byte{}, "", err
	}

	digest, err := GetEIP712Message(c)
	if err != nil {
		return []byte{}, "", err
	}

	signature, err := signingFunc(digest)
	if err != nil {
		return []byte{}, "", err
	}

	return c, signature, nil
}

func SignRequest(protocol *models.Protocol, req interface{}, signingFunc func([]byte) (string, error)) (string, error) {
	message := make(map[string]interface{})

	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	content := base64.StdEncoding.EncodeToString(b)

	message["content"] = content

	sigRequest := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    models.PROTOCOL_NAME,
			Version: models.PROTOCOL_VERSION,
			ChainId: math.NewHexOrDecimal256(1),
		},
		PrimaryType: "Main",
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "version",
					Type: "string",
				},
			},
			"Main": []apitypes.Type{
				{
					Name: "content",
					Type: "string",
				},
			},
		},
		Message: message,
	}

	c, err := json.Marshal(sigRequest)
	if err != nil {
		return "", err
	}

	signature, err := signingFunc(c)
	if err != nil {
		return "", err
	}

	return signature, nil
}

func SignProof(protocol *models.Protocol, proof string, signingFunc func([]byte) (string, error)) (string, string, error) {
	message := make(map[string]interface{})

	message["content"] = proof

	sigRequest := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    models.PROTOCOL_NAME,
			Version: models.PROTOCOL_VERSION,
			ChainId: math.NewHexOrDecimal256(1),
		},
		PrimaryType: "Main",
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "version",
					Type: "string",
				},
			},
			"Main": []apitypes.Type{
				{
					Name: "content",
					Type: "string",
				},
			},
		},
		Message: message,
	}

	c, err := json.Marshal(sigRequest)
	if err != nil {
		return "", "", err
	}

	signature, err := signingFunc(c)
	if err != nil {
		return "", "", err
	}

	sigRequestEncoded := base64.StdEncoding.EncodeToString(c)

	return sigRequestEncoded, signature, nil
}
