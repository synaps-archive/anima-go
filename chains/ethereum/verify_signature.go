package ethereum

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(publicAddress string, data []byte, userSignature string) (bool, error) {
	if len(userSignature) < 3 {
		return false, fmt.Errorf("invalid signature length: %d", len(userSignature))
	}

	if userSignature[0:2] == "0x" {
		userSignature = userSignature[2:]
	}

	message, err := GetEIP712Message(data)
	if err != nil {
		return false, err
	}

	signature, err := hex.DecodeString(userSignature)
	if err != nil {
		return false, err
	}

	if len(signature) != 65 {
		return false, fmt.Errorf("invalid signature length: %d", len(signature))
	}

	if signature[64] != 27 && signature[64] != 28 {
		return false, fmt.Errorf("invalid recovery id: %d", signature[64])
	}
	signature[64] -= 27

	pubKeyRaw, err := crypto.Ecrecover(message, signature)
	if err != nil {
		return false, fmt.Errorf("invalid signature: %s", err.Error())
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if !strings.EqualFold(recoveredAddr.String(), publicAddress) {
		return false, fmt.Errorf("public address and signer address does not match")
	}

	return true, nil
}
