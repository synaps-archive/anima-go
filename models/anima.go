package models

type Protocol struct {
	PrivateKey string
	Network    string
	Chain      string
}

type AnimaOwner struct {
	ID                  string `json:"id"`
	PublicAddress       string `json:"public_address"`
	Chain               string `json:"chain"`
	Wallet              string `json:"wallet"`
	PublicKeyEncryption string `json:"public_key_encryption,omitempty"`
}

type AnimaIssuer struct {
	ID            string `json:"id"`
	PublicAddress string `json:"public_address"`
	Chain         string `json:"chain"`
}

type AnimaVerifier struct {
	ID            string `json:"id"`
	PublicAddress string `json:"public_address"`
	Chain         string `json:"chain"`
}

type AnimaProtocol struct {
	ID            string `json:"id"`
	PublicAddress string `json:"public_address"`
	Chain         string `json:"chain"`
}
