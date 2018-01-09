package models

import (
	"egate/elog"
	"egate/gate"
	"egate/gate/route"
	"egate/model"
	"server/msg"
)

func init() {
	model.MsgParse(model.MSG_PARSER_PROTO)
	model.MsgIn(new(msg.CS_Login))
	model.MsgOut(new(msg.SC_Account), new(msg.SC_Msg))
	model.MsgModelIn(new(Account), "ReqSendMsg")

	route.Handle(0, func(m *msg.CS_Login, agent gate.Agent) {
		elog.Debug("Login:%v-%v", *m.User, *m.Pass)
		model := new(Account)
		model.ModelSetup(model, agent)
	})
}
