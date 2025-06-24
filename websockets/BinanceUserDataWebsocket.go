package websockets

type BinanceUserDataWebsocket interface {
	Close()
	RestartUserDataStream() error
	SetOnMessage(func(messageType int, msg []byte))
	SetOnReconnect(func())
	SetOnError(func(error))
	SetOnClose(func())
}
