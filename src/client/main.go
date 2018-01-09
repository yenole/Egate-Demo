package main

import (
	_ "client/models"
	"egate"
	"egate/gate"
	"egate/gate/network"
	"github.com/golang/protobuf/proto"
	"server/msg"
)

func main() {
	clt := gate.NewClient()
	egate.ClientMode(clt, egate.IO_MODE_MODEL)
	//agt := network.NewWsClient("127.0.0.1:9091", clt)
	agt := network.NewTCPClient("127.0.0.1:9091", clt)
	clt.Agent(agt)
	clt.WriteAgentMsg(&msg.CS_Login{User: proto.String("Tom"), Pass: proto.String("123")})
	clt.Recv()
}
