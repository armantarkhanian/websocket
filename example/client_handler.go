package main

import (
	"github.com/armantarkhanian/websocket"

	"github.com/centrifugal/centrifuge"
)

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
