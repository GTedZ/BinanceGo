package spot

type constants struct {
	BaseUrls           baseUrls
	BaseUrlsMarketData baseUrlsMarketData

	WssBaseUrlsMarketData wssBaseUrlsMarketData
	WssBaseUrls           wssBaseUrls

	WssApiBaseUrls wssApiBaseUrls
}

type baseUrls struct {
	Api     string
	Api_gcp string
	Api1    string
	Api2    string
	Api3    string
	Api4    string
}

type baseUrlsMarketData struct {
	DataApiVision string
}

type wssBaseUrlsMarketData struct {
	DataStreamBinanceVision string
}

type wssBaseUrls struct {
	Stream443  string
	Stream9443 string
}

type wssApiBaseUrls struct {
	WsApi     string
	WsApi9443 string
}

var Constants = constants{
	BaseUrls: baseUrls{
		Api:     "https://api.binance.com",
		Api_gcp: "https://api-gcp.binance.com",
		Api1:    "https://api1.binance.com",
		Api2:    "https://api2.binance.com",
		Api3:    "https://api3.binance.com",
		Api4:    "https://api4.binance.com",
	},
	BaseUrlsMarketData: baseUrlsMarketData{
		DataApiVision: "https://data-api.binance.vision",
	},

	WssBaseUrlsMarketData: wssBaseUrlsMarketData{
		DataStreamBinanceVision: "wss://data-stream.binance.vision",
	},
	WssBaseUrls: wssBaseUrls{
		Stream443:  "wss://stream.binance.com:443",
		Stream9443: "wss://stream.binance.com:9443",
	},

	WssApiBaseUrls: wssApiBaseUrls{
		WsApi:     "wss://ws-api.binance.com:443/ws-api/v3",
		WsApi9443: "wss://ws-api.binance.com:9443/ws-api/v3",
	},
}
