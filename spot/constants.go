package spot

type constants struct {
	BaseUrls           baseUrls
	BaseUrlsMarketData baseUrlsMarketData
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
}
