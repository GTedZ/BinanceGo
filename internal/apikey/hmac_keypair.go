package apikey

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type HmacKeyPair struct {
	apiKey    string
	apiSecret string
}

func (keypair *HmacKeyPair) ApiKey() string { return keypair.apiKey }
func (keypair *HmacKeyPair) Sign(value string) (string, error) {
	h := hmac.New(sha256.New, []byte(keypair.apiSecret))
	_, err := h.Write([]byte(value))
	if err != nil {
		return "", err
	}

	signature := hex.EncodeToString(h.Sum(nil))
	return signature, nil
}
