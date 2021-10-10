package websocket

import "github.com/centrifugal/centrifuge"

type Session interface {
	Authorized() bool
	Credentials() *centrifuge.Credentials
}
