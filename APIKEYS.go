package Binance

type APIKEYS struct {
	key    string
	secret string
}

func (keys *APIKEYS) Set(KEY string, SECRET string) {
	keys.key = KEY
	keys.secret = SECRET
}

func (keys *APIKEYS) Get() (KEY string, SECRET string) {
	return keys.key, keys.secret
}
