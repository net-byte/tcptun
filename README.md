# tcptun

A simple secure tcp tunnel.  

[![Travis](https://travis-ci.com/net-byte/tcptun.svg?branch=master)](https://github.com/net-byte/tcptun)
[![Go Report Card](https://goreportcard.com/badge/github.com/net-byte/tcptun)](https://goreportcard.com/report/github.com/net-byte/tcptun)
![image](https://img.shields.io/badge/License-MIT-orange)
![image](https://img.shields.io/badge/License-Anti--996-red)

# Features
* Proxying tcp to tcp
* Proxying tcp to tls

# Usage  
## Cmd

```
Usage of ./tcptun:  
  -k string
        encryption key (default "")
  -l string
        local address (default ":2000")
  -s string
        server address (default ":2001")
  -tls  tcp to tls mode
```  

## Docker
### Run client(tcp to tcp)
```
docker run -d --restart=always  \ 
--name tcptun-client -p 2000:2000 netbyte/tcptun -l :2000 -s server-ip:2001 -k 123456
```

### Run server(tcp to tcp)
```
docker run  -d --restart=always  \
--net=host --name tcptun-server -p 2001:2001 netbyte/tcptun -l :2001 -s :1080 -k 123456
```

### Run with TLS(tcp to tls)
```
docker run -d --restart=always  \ 
--name tcptun-client -p 2000:2000 netbyte/tcptun -l :2000 -s server-ip:2001 -tls
```
