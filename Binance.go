package Binance

var Utils utils

type Binance struct {
	configs BinanceConfig
	Opts    BinanceOptions

	API APIKEYS

	Spot    Spot
	Futures Futures
}

//

type BinanceConfig struct {
	timestamp_offset int64
}

func (config *BinanceConfig) init() {
	config.timestamp_offset = 0
}

//

type BinanceOptions struct {
	// TODO: need to implement this ASAP!
	updateTimestampOffset bool
	recvWindow            int64
}

func (options *BinanceOptions) init() {
	options.updateTimestampOffset = false
	options.recvWindow = 5000
}

func (options *BinanceOptions) Set_UpdateTimestampOffset(value bool) {
	options.updateTimestampOffset = value
}

func (options *BinanceOptions) Set_recvWindow(recvWindow int64) {
	options.recvWindow = recvWindow
}

// Public funcs

func CreateReadClient() *Binance {
	binance := Binance{}

	binance.configs.init()
	binance.Opts.init()

	binance.Spot.init(&binance)
	binance.Futures.init(&binance)

	return &binance
}

func CreateClient(APIKEY string, APISECRET string) *Binance {
	binance := CreateReadClient()
	binance.API.Set(APIKEY, APISECRET)

	binance.Spot.init(binance)
	binance.Futures.init(binance)

	return binance
}

func CreateClientWithOptions(APIKEY string, APISECRET string, recvWindow int64) *Binance {
	binance := CreateClient(APIKEY, APISECRET)

	binance.Opts.Set_recvWindow(recvWindow)

	return binance
}
