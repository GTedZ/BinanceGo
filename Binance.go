package Binance

import (
	"reflect"

	"github.com/GTedZ/binancego/apikeys"
	"github.com/GTedZ/binancego/lib"
)

// Globals

var Utils lib.Utils
var Binary lib.BinaryUtils

type Error = lib.Error // type alias re-export

// Config

type config struct {
}

//

type Options struct {
	// This sets the default recvWindow of the Binance instance
	RecvWindow int64

	// This is the timestamp offset that will be used throughout the Binance instance.
	//
	// Any signed request that requires a timestamp will add the Timestamp_offset value
	//
	// Used in the case of out-of-sync time clocks
	Timestamp_offset int64
}

type options struct {
	recvWindow       int64
	timestamp_offset int64
}

func (options *options) init(opt_params ...Options) {
	if len(opt_params) != 0 {
		params := opt_params[0]

		options.recvWindow = params.RecvWindow
		options.timestamp_offset = params.Timestamp_offset
	}
}

func (options *options) Set_recvWindow(recvWindow int64) {
	options.recvWindow = recvWindow
}

func (options *options) Set_timestampOffset(timestamp_offset int64) {
	options.timestamp_offset = timestamp_offset
}

//

type Binance struct {
	Opts options

	API apikeys.KeyPair

	Spot    Spot
	Futures Futures
}

// Public Functions

// Creates a Read-Only Client
func NewReadClient(opt_params ...Options) *Binance {
	binance := Binance{}

	binance.Opts.init(opt_params...)

	binance.Spot.init(&binance)
	binance.Futures.init(&binance)

	binance.API = &apikeys.Nil_KeyPair{}

	return &binance
}

// Create a Client with HMAC Key Pair
//
// HMAC API Keys are the system generated symmetric keys that are issued to you from the (!https://www.binance.com/en/my/settings/api-management)[API Management Page]
func NewClient(APIKEY string, APISECRET string, opt_params ...Options) *Binance {
	binance := NewReadClient(opt_params...)

	keyPair := apikeys.HMAC_KeyPair{}
	keyPair.FromString(APIKEY, APISECRET)
	binance.API = &keyPair

	return binance
}

////

// Create a Client with Ed25519 Key Pair by providing the file paths
//
// Ed25519 Keys must be user created (Binance doesn't issue them like the traditional HMAC Keys)
func NewClient_Ed25519(APIKEY string, privKey_path string, opt_params ...Options) (*Binance, error) {
	binance := NewReadClient(opt_params...)

	keyPair := apikeys.Ed25519_KeyPair{}
	err := keyPair.FromFiles(APIKEY, privKey_path)
	if err != nil {
		return nil, err
	}

	binance.API = &keyPair

	return binance, nil
}

// Create a Client with Ed25519 Key Pair by providing them in Base64 format
//
// Ed25519 Keys must be user created (Binance doesn't issue them like the traditional HMAC Keys)
func NewClient_Ed25519_fromBase64(APIKEY string, privKey_base64 string, opt_params ...Options) (*Binance, error) {
	binance := NewReadClient(opt_params...)

	keyPair := apikeys.Ed25519_KeyPair{}
	err := keyPair.FromBase64(APIKEY, privKey_base64)
	if err != nil {
		return nil, err
	}

	binance.API = &keyPair

	return binance, nil
}

////

// Create a Client with Ed25519 Key Pair by providing the file paths
//
// Ed25519 Keys must be user created (Binance doesn't issue them like the traditional HMAC Keys)
func NewClient_RSA(APIKEY string, privKey_path string, opt_params ...Options) (*Binance, error) {
	binance := NewReadClient(opt_params...)

	keyPair := apikeys.RSA_KeyPair{}
	err := keyPair.FromFiles(APIKEY, privKey_path)
	if err != nil {
		return nil, err
	}

	binance.API = &keyPair

	return binance, nil
}

// Create a Client with RSA Key Pair by providing them in Base64 format
//
// RSA Keys must be user created (Binance doesn't issue them like the traditional HMAC Keys)
func NewClient_RSA_fromBase64(APIKEY string, privKey_base64 string, opt_params ...Options) (*Binance, error) {
	binance := NewReadClient(opt_params...)

	keyPair := apikeys.RSA_KeyPair{}
	err := keyPair.FromBase64(APIKEY, privKey_base64)
	if err != nil {
		return nil, err
	}

	binance.API = &keyPair

	return binance, nil
}

////

// Needs to be moved to somewhere else
// Checks if a value is different from its default value.
func isDifferentFromDefault(value any) bool {
	// Get the reflect.Value of the input
	val := reflect.ValueOf(value)

	// Get the default value of the type
	defaultValue := reflect.Zero(val.Type()).Interface()

	// Compare the input value with the default value
	return !reflect.DeepEqual(value, defaultValue)
}
