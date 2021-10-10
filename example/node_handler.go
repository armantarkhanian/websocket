package main

import (
	"github.com/armantarkhanian/websocket"

	"github.com/centrifugal/centrifuge"
)

type nodeHandler struct{}

func (h nodeHandler) OnSurvey(*centrifuge.Node, centrifuge.SurveyEvent) centrifuge.SurveyReply {
	return centrifuge.SurveyReply{}
}

func (h nodeHandler) OnNotification(*centrifuge.Node, centrifuge.NotificationEvent) {}

func (h nodeHandler) OnConnecting(_ *centrifuge.Node, e centrifuge.ConnectEvent) (websocket.Session, centrifuge.ConnectReply, error) {
	return nil, centrifuge.ConnectReply{}, nil
}
