#!bin/bash
#Linux
GOOS=linux GOARCH=amd64 go build -o ./bin/tcptun-linux-amd64 ./main.go
#Linux arm
GOOS=linux GOARCH=arm64 go build -o ./bin/tcptun-linux-arm64 ./main.go
#Mac OS
GOOS=darwin GOARCH=amd64 go build -o ./bin/tcptun-darwin-amd64 ./main.go
#Windows
GOOS=windows GOARCH=amd64 go build -o ./bin/tcptun-windows-amd64.exe ./main.go
#Openwrt
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o ./bin/tcptun-openwrt-amd64 ./main.go

echo "DONE!!!"
