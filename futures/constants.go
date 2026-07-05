package futures

type constants struct {
	BaseUrls baseUrls

	WssBaseUrls wssBaseUrls

	WssApiBaseUrls wssApiBaseUrls
}

type baseUrls struct {
	Api string

	// The testnet URL used to route all requests to a test version of the exchange, balances reset every month
	Testnet string
}

type wssBaseUrls struct {
	Stream string

	Testnet string
}

type wssApiBaseUrls struct {
	WsApi string

	WsApiTestnet string
}

var Constants = constants{
	BaseUrls: baseUrls{
		Api: "https://fapi.binance.com",

		Testnet: "https://demo-fapi.binance.com",
	},

	WssBaseUrls: wssBaseUrls{
		Stream: "wss://fstream.binance.com",

		Testnet: "wss://fstream.binancefuture.com",
	},

	WssApiBaseUrls: wssApiBaseUrls{
		WsApi: "wss://ws-fapi.binance.com/ws-fapi/v1",

		WsApiTestnet: "wss://testnet.binancefuture.com/ws-fapi/v1",
	},
}
