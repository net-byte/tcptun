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
  -k string
        Encrypt key (default "6da62287-979a-4eb4-a5ab-8b3d89da134b")
  -l string
        Local address (default ":2000")
  -s string
        Server address (default ":2001")
  -S Server mode
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
--net=host --name tcptun-server -p 2001:2001 netbyte/tcptun -S -l=:2001 -s=mysql-server-ip:3306
```
