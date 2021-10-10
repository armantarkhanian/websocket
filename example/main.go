package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/armantarkhanian/websocket"
	"github.com/centrifugal/centrifuge"
)

func main() {
	config := websocket.Config{
		Engine: websocket.RedisEngine{
			RedisBrokerConfig:          centrifuge.RedisBrokerConfig{},
			RedisPresenceManagerConfig: centrifuge.RedisPresenceManagerConfig{},

			RedisShards: []centrifuge.RedisShardConfig{
				{
					Address: "host:port",
				},
				{
					SentinelAddresses: []string{"host1:port1", "host2:port2", "host3:port3"},
				},
				{
					ClusterAddresses: []string{"host1:port1", "host2:port2", "host3:port3"},
				},
			},
		},

		ClientHandler: &clientHandler{}, // put here your own implementation of websocket.ClientHandler interface
		NodeHandler:   &nodeHandler{},   // put here your own implementation of websocket.NodeHandler interface

		// Where to find authorizatoin token
		TokenLookup: websocket.TokenLookup{
			Header:       "Authorization",
			Cookie:       "SSID",
			HeaderPrefix: "Bearer",
		},

		CentrifugeConfig: centrifuge.Config{},

		WebsocketConfig: centrifuge.WebsocketConfig{},
	}

	node, handler, err := websocket.New(config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(node) // save node object somewhere

	http.Handle("/websocket", handler)

	http.ListenAndServe(":8080", nil)
}
