// Package websocket ...
package websocket

import (
	"context"

	"github.com/centrifugal/centrifuge"
)

func setNodeHandler(node *centrifuge.Node, handler NodeHandler) {
	node.OnSurvey(func(e centrifuge.SurveyEvent, callback centrifuge.SurveyCallback) {
		callback(handler.OnSurvey(node, e))
	})

	node.OnNotification(func(e centrifuge.NotificationEvent) {
		handler.OnNotification(node, e)
	})

	node.OnConnecting(func(ctx context.Context, e centrifuge.ConnectEvent) (centrifuge.ConnectReply, error) {
		if e.Token == "" {
			e.Token, _ = ctx.Value(tokenContextKey).(string)
		}
		session, connectReply, err := handler.OnConnecting(node, e)

		if session != nil {
			connectReply.Context = context.WithValue(ctx, sessionContextKey, session)
			connectReply.Credentials = session.Credentials()
		}

		return connectReply, err
	})
}

func setClientHandler(node *centrifuge.Node, handler ClientHandler) {
	node.OnConnect(func(c *centrifuge.Client) {
		session, _ := c.Context().Value(sessionContextKey).(Session)

		client := &Client{
			Client:  c,
			session: session,
		}

		client.OnAlive(func() {
			handler.OnAlive(client)
		})
		client.OnDisconnect(func(e centrifuge.DisconnectEvent) {
			handler.OnDisconect(client, e)
		})
		client.OnSubscribe(func(e centrifuge.SubscribeEvent, callback centrifuge.SubscribeCallback) {
			callback(handler.OnSubscribe(client, e))
		})
		client.OnUnsubscribe(func(e centrifuge.UnsubscribeEvent) {
			handler.OnUnsubscribe(client, e)
		})
		client.OnPublish(func(e centrifuge.PublishEvent, callback centrifuge.PublishCallback) {
			callback(handler.OnPublish(client, e))
		})
		client.OnRefresh(func(e centrifuge.RefreshEvent, callback centrifuge.RefreshCallback) {
			callback(handler.OnRefresh(client, e))
		})
		client.OnSubRefresh(func(e centrifuge.SubRefreshEvent, callback centrifuge.SubRefreshCallback) {
			callback(handler.OnSubRefresh(client, e))
		})
		client.OnRPC(func(e centrifuge.RPCEvent, callback centrifuge.RPCCallback) {
			callback(handler.OnRPC(client, e))
		})
		client.OnMessage(func(e centrifuge.MessageEvent) {
			handler.OnMessage(client, e)
		})
		client.OnPresence(func(e centrifuge.PresenceEvent, callback centrifuge.PresenceCallback) {
			callback(handler.OnPresence(client, e))
		})
		client.OnPresenceStats(func(e centrifuge.PresenceStatsEvent, callback centrifuge.PresenceStatsCallback) {
			callback(handler.OnPresenceStats(client, e))
		})
		client.OnHistory(func(e centrifuge.HistoryEvent, callback centrifuge.HistoryCallback) {
			callback(handler.OnHistory(client, e))
		})
	})
}
