package apikey

import "fmt"

type ApiKeyOnlyKeyPair struct {
	apiKey string
}

func (keypair *ApiKeyOnlyKeyPair) ApiKey() string { return keypair.apiKey }
func (keypair *ApiKeyOnlyKeyPair) Sign(value string) (string, error) {
	return "", fmt.Errorf("no private key present to create a signature")
}
