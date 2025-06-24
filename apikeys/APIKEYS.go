package apikeys

import (
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

const (
	NIL_KEYPAIR_TYPE = iota
	HMAC_KEYPAIR_TYPE
	Ed25519_KEYPAIR_TYPE
	RSA_KEYPAIR_TYPE
)

type KeyPair interface {
	// Returns which APIKEY type, unlikely to be used by the developer, only used by the library itself
	GetType() int

	// Returns the APIKEY for the keypair, "" is returned if the keypair is invalid or empty
	GetAPIKEY() string

	// Creates a signature with the provided APISECRET, returns "" if the keypair is invalid or empty
	Sign(string) (string, error)
}

////

type Nil_KeyPair struct{}

func (*Nil_KeyPair) GetType() int                { return NIL_KEYPAIR_TYPE }
func (*Nil_KeyPair) GetAPIKEY() string           { return "" }
func (*Nil_KeyPair) Sign(string) (string, error) { return "", nil }

////

type HMAC_KeyPair struct {
	apiKey    string
	apiSecret string
}

func (keypair *HMAC_KeyPair) FromString(APIKEY string, APISECRET string) {
	keypair.apiKey = APIKEY
	keypair.apiSecret = APISECRET
}

func (*HMAC_KeyPair) GetType() int              { return HMAC_KEYPAIR_TYPE }
func (keypair *HMAC_KeyPair) GetAPIKEY() string { return keypair.apiKey }
func (keypair *HMAC_KeyPair) Sign(value string) (string, error) {
	h := hmac.New(sha256.New, []byte(keypair.apiSecret))
	_, err := h.Write([]byte(value))
	if err != nil {
		return "", err
	}

	signature := hex.EncodeToString(h.Sum(nil))
	return signature, nil
}

////

type Ed25519_KeyPair struct {
	apiKey     string
	privateKey ed25519.PrivateKey
}

func (keypair *Ed25519_KeyPair) FromFiles(APIKEY string, privateKey_path string) error {
	keypair.apiKey = APIKEY

	pemData, err := os.ReadFile(privateKey_path)
	if err != nil {
		return fmt.Errorf("failed to read private key file: %w", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return fmt.Errorf("invalid PEM block or wrong type, got: %v", block)
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse PKCS8 private key: %w", err)
	}

	edKey, ok := parsedKey.(ed25519.PrivateKey)
	if !ok {
		return fmt.Errorf("not an Ed25519 private key")
	}
	if len(edKey) != ed25519.PrivateKeySize {
		return fmt.Errorf("invalid private key size: expected %d, got %d", ed25519.PrivateKeySize, len(edKey))
	}

	keypair.privateKey = edKey
	return nil
}

func (keypair *Ed25519_KeyPair) FromBase64(APIKEY string, privateKey_base64 string) error {
	keypair.apiKey = APIKEY

	privData, err := base64.StdEncoding.DecodeString(privateKey_base64)
	if err != nil {
		return fmt.Errorf("invalid private base64: %w", err)
	}
	if len(privData) != ed25519.PrivateKeySize {
		return fmt.Errorf("invalid private key length: %d", len(privData))
	}

	keypair.privateKey = ed25519.PrivateKey(privData)

	return nil
}

func (*Ed25519_KeyPair) GetType() int              { return Ed25519_KEYPAIR_TYPE }
func (keypair *Ed25519_KeyPair) GetAPIKEY() string { return keypair.apiKey }
func (keypair *Ed25519_KeyPair) Sign(value string) (string, error) {
	raw_signature := ed25519.Sign(keypair.privateKey, []byte(value))
	signature := base64.StdEncoding.EncodeToString(raw_signature)
	fmt.Println("signature:", signature)
	return signature, nil
}

////

type RSA_KeyPair struct {
	apiKey     string
	privateKey *rsa.PrivateKey
}

func (keypair *RSA_KeyPair) FromFiles(APIKEY string, privateKey_path string) error {
	keypair.apiKey = APIKEY

	// Load private key from PEM file
	pemData, err := os.ReadFile(privateKey_path)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return errors.New("invalid PEM block")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return errors.New("not an RSA private key")
	}

	keypair.privateKey = rsaKey
	return nil
}

func (keypair *RSA_KeyPair) FromBase64(APIKEY string, privateKey_base64 string) error {
	keypair.apiKey = APIKEY

	keyBytes, err := base64.StdEncoding.DecodeString(privateKey_base64)
	if err != nil {
		return err
	}

	key, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return errors.New("not an RSA private key")
	}

	keypair.privateKey = rsaKey
	return nil
}

func (keypair *RSA_KeyPair) GetType() int      { return RSA_KEYPAIR_TYPE }
func (keypair *RSA_KeyPair) GetAPIKEY() string { return keypair.apiKey }
func (keypair *RSA_KeyPair) Sign(payload string) (string, error) {
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
