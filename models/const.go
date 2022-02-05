package models

const (
	MAINNET   = "protocol.anima.io:443"
	TESTNET   = "protocol-tesnet.anima.io:443"
	CHAIN_ETH = "ETH"
)

var AVAILABLE_CHAIN = []string{CHAIN_ETH}
var AVAILABLE_NET = []string{MAINNET, TESTNET}
