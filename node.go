// Package websocket ...
package websocket

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/centrifugal/centrifuge"
)

func New(config Config) (*centrifuge.Node, http.Handler, error) {
	node, err := centrifuge.New(config.CentrifugeConfig)
	if err != nil {
		return nil, nil, err
	}

	if err = setEngine(node, config.Engine); err != nil {
		return nil, nil, err
	}

	setNodeHandler(node, config.NodeHandler)

	setClientHandler(node, config.ClientHandler)

	if err := node.Run(); err != nil {
		return nil, nil, err
	}

	var handler http.Handler = centrifuge.NewWebsocketHandler(node, config.WebsocketConfig)

	handler = authMiddleware(handler, config.TokenLookup)

	return node, handler, nil
}

func authMiddleware(h http.Handler, tokenLookup TokenLookup) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string

		if tokenLookup.Header != "" {
			token = r.Header.Get(tokenLookup.Header)

			if token != "" && tokenLookup.HeaderPrefix != "" {
				ss := strings.Split(token, tokenLookup.HeaderPrefix)
				if len(ss) == 2 {
					token = strings.TrimSpace(ss[1])
				}
			}

		}

		if token == "" && tokenLookup.Cookie != "" {
			cookie, err := r.Cookie(tokenLookup.Cookie)
			if err == nil {
				token, err = url.QueryUnescape(cookie.Value)
				if err != nil {
					token = cookie.Value
				}
			}
		}

		ctx := context.WithValue(r.Context(), tokenContextKey, token)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func setEngine(node *centrifuge.Node, engine Engine) error {
	if engine == nil {
		return ErrMissingEngine
	}
	return engine.Set(node)
}
