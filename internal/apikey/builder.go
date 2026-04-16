package apikey

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// builder provides helper constructors for creating different types of
// Binance API authentication key pairs.
//
// Binance supports multiple authentication schemes depending on the API
// features being used. The builder exposes convenience methods to create
// the appropriate KeyPair implementation for each supported scheme.
//
// Each method returns a type implementing the KeyPair interface, which is
// later used by the client to sign authenticated requests.
type builder struct{}

var Builder builder

// ApiKeyOnly creates a KeyPair that only contains an API key.
//
// This authentication mode is used for endpoints that require an API key
// but do not require request signing. A common example is accessing
// Binance User Data Streams.
//
// In this mode:
//
//   - The API key is sent via the `X-MBX-APIKEY` header.
//   - No cryptographic signing is performed.
//
// Parameters:
//
//	APIKEY - The Binance API key generated from the Binance API management
//	console.
//
// Returns:
//
//	A KeyPair capable of authenticating requests that only require an API key.
func (builder) ApiKeyOnly(APIKEY string) KeyPair {
	return &ApiKeyOnlyKeyPair{
		apiKey: APIKEY,
	}
}

// HMAC creates a KeyPair using Binance's HMAC-SHA256 authentication.
//
// This is the most common authentication method used by Binance REST APIs.
// Requests are signed using the API secret via HMAC-SHA256.
//
// In this mode:
//
//   - The API key is sent via the `X-MBX-APIKEY` header.
//   - Requests are signed using HMAC-SHA256.
//   - The signature is appended to the query string.
//
// Parameters:
//
//	APIKEY   - Binance API key.
//	APISECRET - Binance API secret used to compute the HMAC signature.
//
// Returns:
//
//	A KeyPair capable of signing requests using HMAC-SHA256.
func (builder) HMAC(APIKEY string, APISECRET string) KeyPair {
	return &HmacKeyPair{
		apiKey:    APIKEY,
		apiSecret: APISECRET,
	}
}

// Ed25519FromFile loads an Ed25519 private key from a PEM-encoded PKCS#8 file.
//
// Binance supports Ed25519 key-based authentication as an alternative to
// HMAC-based authentication. Ed25519 provides modern public-key signatures
// and allows signing requests using an asymmetric key pair.
//
// The provided file must:
//
//   - Be PEM encoded
//   - Contain a PKCS#8 formatted private key
//   - Contain an Ed25519 key
//
// Parameters:
//
//	APIKEY           - Binance API key associated with the Ed25519 key pair.
//	privateKey_path  - Path to a PEM file containing a PKCS#8 Ed25519 private key.
//
// Returns:
//
//	A KeyPair configured to sign requests using the Ed25519 private key.
func (builder) Ed25519FromFile(APIKEY string, privateKey_path string) (KeyPair, error) {
	pemData, err := os.ReadFile(privateKey_path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("invalid PEM block or wrong type, got: %v", block)
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PKCS8 private key: %w", err)
	}

	edKey, ok := parsedKey.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an Ed25519 private key")
	}
	if len(edKey) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("invalid private key size: expected %d, got %d", ed25519.PrivateKeySize, len(edKey))
	}

	return &Ed25519KeyPair{
		apiKey:     APIKEY,
		privateKey: edKey,
	}, nil
}

// Ed25519Base64 creates an Ed25519 KeyPair from a base64-encoded private key.
//
// This method is useful when the private key is stored in environment
// variables, configuration files, or secret managers that store keys in
// base64 form rather than PEM files.
//
// The provided string must be:
//
//   - Base64 encoded
//   - A raw Ed25519 private key
//   - Exactly ed25519.PrivateKeySize bytes once decoded.
//
// Parameters:
//
//	APIKEY            - Binance API key associated with the Ed25519 key.
//	privateKey_base64 - Base64 encoded Ed25519 private key.
//
// Returns:
//
//	A KeyPair capable of signing requests using Ed25519.
func (builder) Ed25519Base64(APIKEY string, privateKey_base64 string) (KeyPair, error) {
	privData, err := base64.StdEncoding.DecodeString(privateKey_base64)
	if err != nil {
		return nil, fmt.Errorf("invalid private base64: %w", err)
	}
	if len(privData) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("invalid private key length: %d", len(privData))
	}

	edKey := ed25519.PrivateKey(privData)

	return &Ed25519KeyPair{
		apiKey:     APIKEY,
		privateKey: edKey,
	}, nil
}

// RSAFromFile loads an RSA private key from a PEM-encoded PKCS#8 file.
//
// Binance also supports RSA-based authentication where requests are signed
// using an RSA private key instead of HMAC.
//
// The file must:
//
//   - Be PEM encoded
//   - Contain a PKCS#8 formatted private key
//   - Contain an RSA private key
//
// Parameters:
//
//	APIKEY          - Binance API key associated with the RSA key.
//	privateKey_path - Path to the PEM file containing the RSA private key.
//
// Returns:
//
//	A KeyPair configured to sign requests using RSA.
func (builder) RSAFromFile(APIKEY string, privateKey_path string) (KeyPair, error) {
	// Load private key from PEM file
	pemData, err := os.ReadFile(privateKey_path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("invalid PEM block")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}

	return &RSAKeyPair{
		apiKey:     APIKEY,
		privateKey: rsaKey,
	}, nil
}

// RSABase64 creates an RSA KeyPair from a base64-encoded PKCS#8 private key.
//
// This method is useful when storing RSA keys in environment variables
// or secret managers.
//
// The base64 string must decode to a PKCS#8 encoded RSA private key.
//
// Parameters:
//
//	APIKEY            - Binance API key associated with the RSA key.
//	privateKey_base64 - Base64 encoded PKCS#8 RSA private key.
//
// Returns:
//
//	A KeyPair configured to sign requests using RSA.
func (builder) RSABase64(APIKEY string, privateKey_base64 string) (KeyPair, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(privateKey_base64)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}

	return &RSAKeyPair{
		apiKey:     APIKEY,
		privateKey: rsaKey,
	}, nil
}
