package apikey

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"errors"
)

type RSAKeyPair struct {
	apiKey     string
	privateKey *rsa.PrivateKey
}

func (keypair *RSAKeyPair) ApiKey() string { return keypair.apiKey }
func (keypair *RSAKeyPair) Sign(payload string) (string, error) {
	if keypair.privateKey == nil {
		return "", errors.New("private key not loaded")
	}

	hashed := crypto.SHA256.New()
	hashed.Write([]byte(payload))
	digest := hashed.Sum(nil)

	sig, err := rsa.SignPKCS1v15(nil, keypair.privateKey, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sig), nil
}
