package apikey

import "fmt"

type NilKeyPair struct{}

func (*NilKeyPair) ApiKey() string {
	return ""
}

func (*NilKeyPair) Sign(string) (string, error) {
	return "", fmt.Errorf("unable to create a signature without valid API Keys")
}
