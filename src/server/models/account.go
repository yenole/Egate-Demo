package models

import (
	"egate/elog"
	"egate/model"
	"github.com/golang/protobuf/proto"
	"server/msg"
)

type Account struct {
	model.Model
	Id uint64
}

func (a *Account) ID() uint64 {
	return a.Id
}
func (a *Account) SetId(v uint64) {
	a.Id = v
}
func (a *Account) OnAgentOpen() {
	elog.Debug("OnAgentOpen")
	a.AgentWriteMsg(&msg.SC_Account{Id: proto.Uint64(1), Nick: proto.String("Tom")})
}

func (a *Account) ReqSendMsg(m *msg.CS_SendMsg) {
	elog.Debug("SendMsg:%v", *m.Msg)
	a.AgentWriteMsg(&msg.SC_Msg{Msg: m.Msg})
}
