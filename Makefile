ws-protoc:
	PATH=`pwd`/tools/protoc/`uname`/ protoc --go_out=. pb/pb.proto
	mv ./pb/pb.pb.go ./src/server/msg/pb.pb.go

ws-win:
	GOPATH=`pwd` go get -d server
	GOPATH=`pwd` GOOS=windows go build -tags "ws" -o bin/Ws.exe server
