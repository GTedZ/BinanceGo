package apikey

type NilKeyPair struct{}

func (*NilKeyPair) ApiKey() string              { return "" }
func (*NilKeyPair) Sign(string) (string, error) { return "", nil }
