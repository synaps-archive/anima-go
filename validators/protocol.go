package validators

import (
	"fmt"

	"github.com/anima-protocol/anima-go/chains/ethereum"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/utils"
)

func ValidateProtocol(protocol *models.Protocol) error {
	if !utils.InArray(protocol.Chain, models.AVAILABLE_CHAIN) {
		return fmt.Errorf("chain unavailable")
	}

	if !utils.InArray(protocol.Network, models.AVAILABLE_NET) {
		return fmt.Errorf("invalid network")
	}

	switch protocol.Chain {
	case models.CHAIN_ETH:
		return ethereum.RecoverAccount(protocol.PrivateKey)
	}
	return nil
}
