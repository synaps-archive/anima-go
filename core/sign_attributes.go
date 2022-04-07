package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/anima-protocol/anima-go/chains/evm"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func SignIssuing(anima *models.Protocol, issuer *protocol.AnimaIssuer, request *protocol.IssueRequest, signingFunc func([]byte) (string, error)) (*protocol.IssueRequest, error) {
	issuingAuthorization, err := GetIssuingAuthorization(request)
	if err != nil {
		return nil, err
	}

	owner := &protocol.AnimaOwner{
		Id:            issuingAuthorization.Owner.ID,
		PublicAddress: issuingAuthorization.Owner.PublicAddress,
		Chain:         issuingAuthorization.Owner.Chain,
		Wallet:        issuingAuthorization.Owner.Wallet,
	}

	documentBytes, err := json.Marshal(request.Document)
	if err != nil {
		return nil, err
	}

	documentContentBytes := new(bytes.Buffer)
	err = json.Compact(documentContentBytes, documentBytes)
	if err != nil {
		return nil, err
	}

	documentHash := crypto.Hash(documentContentBytes.Bytes())
	documentIdentifier := fmt.Sprintf("anima:document:%s", documentHash)
	documentSpecs := request.Document.Specs

	// Sign Attributes
	for name := range request.Attributes {
		request.Attributes[name].Credential.Content.Source = &protocol.IssCredentialSource{
			Id:    documentIdentifier,
			Specs: documentSpecs,
		}

		request.Attributes[name].Credential.Content.Owner = owner
		request.Attributes[name].Credential.Content.Issuer = issuer

		switch anima.Chain {
		case models.CHAIN_ETH:
			signature, err := evm.SignCredential(anima, request.Attributes[name].Credential.Content, signingFunc)
			if err != nil {
				return nil, err
			}

			request.Attributes[name].Credential.Signature = "0x" + signature
		}
	}

	// Sign Proof
	proofContent, err := base64.StdEncoding.DecodeString(request.Proof.Content)
	if err != nil {
		return nil, err
	}

	proofValue := map[string]interface{}{}
	err = json.Unmarshal(proofContent, &proofValue)
	if err != nil {
		return nil, err
	}

	switch anima.Chain {
	case models.CHAIN_ETH:
		proofSignature, err := evm.SignCredential(anima, proofValue, signingFunc)
		if err != nil {
			return nil, err
		}

		request.Proof.Signature = "0x" + proofSignature
	}

	return request, nil
}
