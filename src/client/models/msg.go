package models

import (
	"egate/elog"
	"egate/gate"
	"egate/gate/route"
	"egate/model"
	"github.com/golang/protobuf/proto"
	"server/msg"
)

func init() {
	model.MsgParse(model.MSG_PARSER_PROTO)
	model.MsgIn(new(msg.SC_Account), new(msg.SC_Msg))
	model.MsgOut(new(msg.CS_Login), new(msg.CS_SendMsg))

	route.Handle(0, func(m *msg.SC_Account, agent gate.Agent) {
		elog.Debug(">>>>%v:%v", *m.Id, *m.Nick)
		agent.AgentWriteMsg(&msg.CS_SendMsg{Msg: proto.String("hello world")})
	})

	route.Handle(1, func(m *msg.SC_Msg) {
		elog.Debug(">>>>Msg:%v", *m.Msg)
	})
}
