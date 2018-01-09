package main

import (
	"egate"
	"egate/gate"
	_ "server/models"
)

func main() {
	egt := gate.NewEgate()

	egate.Mode(egt, egate.IO_MODE_MODEL)

	//egate.WsAccpet(egt, "127.0.0.1:9091")
	egate.TcpAccpet(egt, "127.0.0.1:9091")

	egate.Work(egt)
}
