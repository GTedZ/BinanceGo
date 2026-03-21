package apikey

import (
	"crypto/ed25519"
	"encoding/base64"
)

type Ed25519KeyPair struct {
	apiKey     string
	privateKey ed25519.PrivateKey
}

func (keypair *Ed25519KeyPair) ApiKey() string { return keypair.apiKey }
func (keypair *Ed25519KeyPair) Sign(value string) (string, error) {
	raw_signature := ed25519.Sign(keypair.privateKey, []byte(value))
	signature := base64.StdEncoding.EncodeToString(raw_signature)
	return signature, nil
}
