package requests

type RequestSecurityType string

const (
	NONE     RequestSecurityType = "NONE"
	API_ONLY RequestSecurityType = "API_ONLY"
	SIGNED   RequestSecurityType = "SIGNED"
)
