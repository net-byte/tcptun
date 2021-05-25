#!bin/bash

export GO111MODULE=on
#Linux
GOOS=linux GOARCH=amd64 go build -o ./bin/tcptun-linux-amd64 ./main.go
#Mac OS
GOOS=darwin GOARCH=amd64 go build -o ./bin/tcptun-darwin-amd64 ./main.go
#Windows
GOOS=windows GOARCH=amd64 go build -o ./bin/tcptun-windows-amd64.exe ./main.go
#Openwrt
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o ./bin/tcptun-openwrt-amd64 ./main.go

echo "done!"
