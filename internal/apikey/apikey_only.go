package apikey

type ApiKeyOnlyKeyPair struct {
	apiKey string
}

func (keypair *ApiKeyOnlyKeyPair) ApiKey() string { return keypair.apiKey }
func (keypair *ApiKeyOnlyKeyPair) Sign(value string) (string, error) {
	return "", nil
}
