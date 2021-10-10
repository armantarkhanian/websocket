package websocket

import (
	"sync"

	"github.com/centrifugal/centrifuge"
)

type Client struct {
	*centrifuge.Client
	rwm     sync.RWMutex
	session Session
}

func (c *Client) GetSession() Session {
	c.rwm.RLock()
	defer c.rwm.RUnlock()
	return c.session
}

func (c *Client) Authorized() bool {
	c.rwm.RLock()
	defer c.rwm.RUnlock()
	return c.session.Authorized()
}

func (c *Client) SetSession(i Session) {
	c.rwm.Lock()
	defer c.rwm.Unlock()
	c.session = i
}
