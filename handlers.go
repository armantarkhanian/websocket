package websocket

import (
	"github.com/centrifugal/centrifuge"
)

type ClientHandler interface {
	OnAlive(*Client)
	OnDisconect(*Client, centrifuge.DisconnectEvent)
	OnSubscribe(*Client, centrifuge.SubscribeEvent) (centrifuge.SubscribeReply, error)
	OnUnsubscribe(*Client, centrifuge.UnsubscribeEvent)
	OnPublish(*Client, centrifuge.PublishEvent) (centrifuge.PublishReply, error)
	OnRefresh(*Client, centrifuge.RefreshEvent) (centrifuge.RefreshReply, error)
	OnSubRefresh(*Client, centrifuge.SubRefreshEvent) (centrifuge.SubRefreshReply, error)
	OnRPC(*Client, centrifuge.RPCEvent) (centrifuge.RPCReply, error)
	OnMessage(*Client, centrifuge.MessageEvent)
	OnPresence(*Client, centrifuge.PresenceEvent) (centrifuge.PresenceReply, error)
	OnPresenceStats(*Client, centrifuge.PresenceStatsEvent) (centrifuge.PresenceStatsReply, error)
	OnHistory(*Client, centrifuge.HistoryEvent) (centrifuge.HistoryReply, error)
}

type NodeHandler interface {
	OnSurvey(*centrifuge.Node, centrifuge.SurveyEvent) centrifuge.SurveyReply
	OnNotification(*centrifuge.Node, centrifuge.NotificationEvent)
	OnConnecting(*centrifuge.Node, centrifuge.ConnectEvent) (Session, centrifuge.ConnectReply, error)
}
