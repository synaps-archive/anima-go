package ethereum

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/anima-protocol/anima-go/models"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func Sign(privateKey string, data []byte) (string, error) {
	hashChallenge, err := GetEIP712Message(data)
	if err != nil {
		return "", err
	}

	var key *ecdsa.PrivateKey
	var bytes []byte

	if b, err := hex.DecodeString(privateKey); err != nil {
		return "", err
	} else {
		bytes = b
	}
	if pk, err := crypto.ToECDSA(bytes); err != nil {
		return "", err
	} else {
		key = pk
	}
	sig, err := crypto.Sign(hashChallenge, key)
	if err != nil {
		return "", err
	}

	sig[64] += 27

	return "0x" + hex.EncodeToString(sig), nil
}

func SignIssueAttribute(protocol *models.Protocol, issAttr *models.IssueAttribute) ([]byte, string, error) {
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

	signature, err := Sign(protocol.PrivateKey, c)
	if err != nil {
		return []byte{}, "", err
	}

	return c, signature, nil
}

func SignRequest(protocol *models.Protocol, req interface{}) (string, error) {
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

	signature, err := Sign(protocol.PrivateKey, c)
	if err != nil {
		return "", err
	}

	return signature, nil
}

func SignProof(protocol *models.Protocol, proof string) (string, string, error) {
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

	signature, err := Sign(protocol.PrivateKey, c)
	if err != nil {
		return "", "", err
	}

	sigRequestEncoded := base64.StdEncoding.EncodeToString(c)

	return sigRequestEncoded, signature, nil
}
