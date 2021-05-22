# tcptun

A simple secure tcp tunnel.  

[![Travis](https://travis-ci.com/net-byte/tcptun.svg?branch=master)](https://github.com/net-byte/tcptun)
[![Go Report Card](https://goreportcard.com/badge/github.com/net-byte/tcptun)](https://goreportcard.com/report/github.com/net-byte/tcptun)
![image](https://img.shields.io/badge/License-MIT-orange)
![image](https://img.shields.io/badge/License-Anti--996-red)


# Usage  
## Cmd

```
Usage of ./tcptun:  
  -S    server mode
  -k string
        encryption key (default "123456")
  -l string
        local address (default ":2000")
  -s string
        server address (default ":2001")
```  

## Docker
### Run client
```
docker run -d --restart=always  \ 
--name tcptun-client -p 2000:2000 netbyte/tcptun -l=:2000 -s=server-ip:2001
```

### Run server
```
docker run  -d --restart=always  \
--net=host --name tcptun-server -p 2001:2001 netbyte/tcptun -S -l=:2001 -s=:1080
```
