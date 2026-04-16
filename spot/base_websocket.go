package spot

import (
	"time"

	"github.com/GTedZ/binancego/internal/berror"
)

type baseWebsocket interface {
	Close()

	SendRequest(message map[string]interface{}, timeout time.Duration) (data []byte, err berror.Error)
}
