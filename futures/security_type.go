package futures

type SecurityType string

const (
	NONE        SecurityType = "NONE"
	TRADE       SecurityType = "TRADE"
	USER_DATA   SecurityType = "USER_DATA"
	USER_STREAM SecurityType = "USER_STREAM"
)
