# Example
```golang
package main

import (
	"fmt"
	"net/http"

	"github.com/armantarkhanian/websocket"
	"github.com/centrifugal/centrifuge"
)

var config = websocket.Config{
	Engine: websocket.RedisEngine{
		RedisBrokerConfig:          centrifuge.RedisBrokerConfig{},
		RedisPresenceManagerConfig: centrifuge.RedisPresenceManagerConfig{},
		RedisShards: []centrifuge.RedisShardConfig{
			{
				Address: "localhost:6379",
			},
			{
				SentinelAddresses: []string{"localhost:6379"},
			},
			{
				ClusterAddresses: []string{"localhost:6379"},
			},
		},
	},

	ClientHandler: &clientHandler{},
	NodeHandler:   &nodeHandler{},

	TokenLookup: websocket.TokenLookup{
		Header:       "Authorization",
		Cookie:       "SSID",
		HeaderPrefix: "Bearer",
	},

	CentrifugeConfig: centrifuge.Config{},

	WebsocketConfig: centrifuge.WebsocketConfig{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		UseWriteBufferPool: true,
	},
}

func main() {
	node, handler, err := websocket.New(config)
	if err != nil {
		panic(err)
	}

	fmt.Println(node) // save node object somewhere

	http.Handle("/websocket", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// client Handler
type clientHandler struct{}

func (h clientHandler) OnAlive(*websocket.Client) {}

func (h clientHandler) OnDisconect(*websocket.Client, centrifuge.DisconnectEvent) {}

func (h clientHandler) OnSubscribe(*websocket.Client, centrifuge.SubscribeEvent) (centrifuge.SubscribeReply, error) {
	return centrifuge.SubscribeReply{}, nil
}

func (h clientHandler) OnUnsubscribe(*websocket.Client, centrifuge.UnsubscribeEvent) {}

func (h clientHandler) OnPublish(*websocket.Client, centrifuge.PublishEvent) (centrifuge.PublishReply, error) {
	return centrifuge.PublishReply{}, nil
}

func (h clientHandler) OnRefresh(*websocket.Client, centrifuge.RefreshEvent) (centrifuge.RefreshReply, error) {
	return centrifuge.RefreshReply{}, nil
}

func (h clientHandler) OnSubRefresh(*websocket.Client, centrifuge.SubRefreshEvent) (centrifuge.SubRefreshReply, error) {
	return centrifuge.SubRefreshReply{}, nil
}

func (h clientHandler) OnMessage(*websocket.Client, centrifuge.MessageEvent) {}

func (h clientHandler) OnPresence(*websocket.Client, centrifuge.PresenceEvent) (centrifuge.PresenceReply, error) {
	return centrifuge.PresenceReply{}, nil
}

func (h clientHandler) OnPresenceStats(*websocket.Client, centrifuge.PresenceStatsEvent) (centrifuge.PresenceStatsReply, error) {
	return centrifuge.PresenceStatsReply{}, nil
}

func (h clientHandler) OnRPC(*websocket.Client, centrifuge.RPCEvent) (centrifuge.RPCReply, error) {
	return centrifuge.RPCReply{}, nil
}

func (h clientHandler) OnHistory(*websocket.Client, centrifuge.HistoryEvent) (centrifuge.HistoryReply, error) {
	return centrifuge.HistoryReply{}, nil
}

// node Handler
type nodeHandler struct{}

func (h nodeHandler) OnSurvey(*centrifuge.Node, centrifuge.SurveyEvent) centrifuge.SurveyReply {
	return centrifuge.SurveyReply{}
}

func (h nodeHandler) OnNotification(*centrifuge.Node, centrifuge.NotificationEvent) {}

func (h nodeHandler) OnConnecting(_ *centrifuge.Node, e centrifuge.ConnectEvent) (websocket.Session, centrifuge.ConnectReply, error) {
	return nil, centrifuge.ConnectReply{}, nil
}
```
