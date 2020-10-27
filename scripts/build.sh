#!bin/bash
UNAME=$(uname)

if [ "$UNAME" == "Linux" ] ; then
    GOOS=linux GOARCH=amd64 go build -o ./bin/tcptun ./main.go
elif [ "$UNAME" == "Darwin" ] ; then
    GOOS=darwin GOARCH=amd64 go build -o ./bin/tcptun ./main.go
elif [[ "$UNAME" == CYGWIN* || "$UNAME" == MINGW* ]] ; then
    GOOS=windows GOARCH=amd64 go build -o ./bin/tcptun.exe ./main.go
fi

echo "done!"
