package futures

import (
	"github.com/GTedZ/binancego/internal/apikey"
	"github.com/GTedZ/binancego/internal/berror"
	"github.com/GTedZ/binancego/internal/requests"
)

////
// Re-exported types
////

// KeyPair represents an authentication credential used by the client to
// authenticate requests against the Binance API.
//
// Different Binance endpoints require different authentication schemes.
// Implementations of KeyPair encapsulate the logic required to attach
// authentication information to outgoing requests and, when necessary,
// sign them.
//
// The library provides multiple built-in KeyPair implementations covering
// all supported Binance authentication methods:
//
//   - API Key only
//     Used for endpoints that require the API key header but do not require
//     request signing (for example User Data Stream management).
//
//   - HMAC-SHA256
//     The most common authentication method used by Binance REST APIs.
//     Requests are signed using the API secret and the signature is appended
//     to the query string.
//
//   - RSA
//     Asymmetric request signing using an RSA private key. The corresponding
//     public key must be registered with Binance.
//
//   - Ed25519
//     Modern asymmetric signing using an Ed25519 key pair. Like RSA, the
//     public key must be registered with Binance.
//
// A KeyPair is passed to the client during initialization and is used
// internally whenever an endpoint requires authenticated access.
//
// Users typically do not implement this interface themselves; instead,
// they construct a KeyPair using the helper functions provided in the
// apikey package.
type KeyPair = apikey.KeyPair

// APIKEY provides helper constructors for creating different types of
// Binance API authentication key pairs.
//
// Binance supports multiple authentication schemes depending on the API
// features being used. The builder exposes convenience methods to create
// the appropriate KeyPair implementation for each supported scheme.
//
// Each method returns a type implementing the KeyPair interface, which is
// later used by the client to sign authenticated requests.
var APIKEY = apikey.Builder

type Error = berror.Error

type RequestMethod = requests.Method

var Methods = requests.Methods

type Response = requests.Response
