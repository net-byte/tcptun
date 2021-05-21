#!bin/bash

export GO111MODULE=on
UNAME=$(uname)
ARCH=$(uname -m)

if [[ "$UNAME" == "Linux" && "$ARCH" == "x86_64" ]] ; then
    GOOS=linux GOARCH=amd64 go build -o ./bin/tcptun ./main.go
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o ./bin/tcptun-openwrt-amd64 ./main.go
elif [[ "$UNAME" == "Linux" && "$ARCH" == "aarch64" ]] ; then
    GOOS=linux GOARCH=arm64 go build -o ./bin/tcptun_arm64 ./main.go
elif [ "$UNAME" == "Darwin" ] ; then
    GOOS=darwin GOARCH=amd64 go build -o ./bin/tcptun ./main.go
elif [[ "$UNAME" == CYGWIN* || "$UNAME" == MINGW* ]] ; then
    GOOS=windows GOARCH=amd64 go build -o ./bin/tcptun.exe ./main.go
fi

echo "done!"
