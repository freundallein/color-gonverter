# color-gonverter
[![Build Status](https://travis-ci.org/freundallein/color-gonverter.svg?branch=master)](https://travis-ci.org/freundallein/color-gonverter)

Simple gRPC service for converting color-points from ARGB to RGBA (32/64bit) 

## Installation
### With docker  
```
$> docker pull freundallein/gonverter
```
### With source
```
$> git clone git@github.com:freundallein/color-gonverter.git
$> cd color-gonverter
$> make build
```

## Usage
Docker-compose

```
version: "3.5"

networks:
  network:
    name: example-network
    driver: bridge

services:
  gonverter:
    image: freundallein/gonverter:latest
    container_name: gonverter
    restart: always
    networks: 
      - network

```


You can do 2 different RPC calls ConvertARGB_RGBA32 and ConvertARGB_RGBA64, based on gRPC stream.
```
0xAABBCCDD -> 0xBBCCDDAA                 32-bit
0xAABBCCDDAABBCCDD -> 0xBBCCDDAABBCCDDAA 64-bit
```

### Protocol

```
service GOnverter {
    rpc ConvertARGB_RGBA32(stream ARGB32) returns (stream RGBA32) {}
    rpc ConvertARGB_RGBA64(stream ARGB64) returns (stream RGBA64) {}
}


message ARGB32 {
    uint32 argb = 1;
}

message RGBA32 {
    uint32 rgba = 1;
}

message ARGB64 {
    uint64 argb = 1;
}

message RGBA64 {
    uint64 rgba = 1;
}
```
