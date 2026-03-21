package apikey

// KeyPair represents an authentication credential used by the client to
// authenticate requests against the Binance API.
//
// Different Binance endpoints require different authentication schemes.
// Implementations of KeyPair encapsulate the logic required to attach
// authentication information to outgoing requests and, when necessary,
// sign them.
type KeyPair interface {
	ApiKey() string
	Sign(payload string) (string, error)
}
