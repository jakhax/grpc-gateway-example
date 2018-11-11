# GRPC GOLANG EXAMPLE
## Description
- In this project we will build a simple golang application that uses grpc to enable a server/client applications to communicate to each other.

## Application requirements
1. Uses GPRC to communicate between client & server.
2. TLS support.
3. Authentication.
4. Support for existing REST clients. 

## Requirements
- An understanding of golang programming syntax.

## Why use GPRC
- Reasons for using grpc (instead of REST) can be best explained in the following articles, we will focus on building a production ready GRPC app.
    * [https://grpc.io/blog/vendastagrpc](https://grpc.io/blog/vendastagrpc)
    * [https://stackoverflow.com/questions/43682366/how-is-grpc-different-from-rest](https://stackoverflow.com/questions/43682366/how-is-grpc-different-from-rest)
    * [https://thenewstack.io/grpc-lean-mean-communication-protocol-microservices/](https://thenewstack.io/grpc-lean-mean-communication-protocol-microservices/)

## Getting Started 

- You can refer to this [official documentation to setting up grpc & protocol buffers](https://grpc.io/docs/quickstart/go.html) 

```bash
go get -u google.golang.org/grpc
# install protoc buffers 3 from  https://github.com/google/protobuf/releases
go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin
# install
protoc -I api/ -I${GOPATH}/src --go_out=plugins=grpc:api api/api.

```

## Defining the protocol using proto buffers.
- First we need to define the protocol, we will use proto buffs for this, at the root of the directory create a directory `api` with a file `api.proto`.
- Protobuffs enable us to define what can be communicated between the client and the server, this is done using `services` and `messages`.
- A `service` is a collection of actions that the server can perform at the clients request, a `message` is the content of the request.
- write the following in the `api/api.proto` file
```proto3
syntax="proto3";

package api;

service Ping {
    rpc SayHello(PingMessage) returns (PingMessage);
}

message PingMessage {
    string greeting=1;
}
``` 



