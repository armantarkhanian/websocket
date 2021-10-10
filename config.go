package websocket

import (
	"github.com/centrifugal/centrifuge"
)

type contextKeyType int

var (
	tokenContextKey   contextKeyType = 0
	sessionContextKey contextKeyType = 1
)

type Config struct {
	Engine           Engine
	ClientHandler    ClientHandler
	NodeHandler      NodeHandler
	TokenLookup      TokenLookup
	CentrifugeConfig centrifuge.Config
	WebsocketConfig  centrifuge.WebsocketConfig
}

// Priority:
// 1. setToken(token) function of centrifuge-client library
// 2. Header value
// 3. Cookie value
type TokenLookup struct {
	Header       string
	Cookie       string
	HeaderPrefix string
}
