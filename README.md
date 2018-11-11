# GRPC Gateway Example

a simple client/server system, in Go, with TLS, authentication & REST support for non-grpc compliant clients.
#### Note: I am not yet finished documenting the branches below, the source code is complete though.

## Table Of Contents
- [Introduction](https://github.com/jakhax/grpc-gateway-example/tree/1.introduction)
    * Setup, introduction to grpc in golang & protocol buffers
- [Simple GRPC server](https://github.com/jakhax/grpc-gateway-example/tree/2.simple-grpc-server)
    * Build a simple GRPC server
- [Simple GRPC client](https://github.com/jakhax/grpc-gateway-example/tree/3.simple-grpc-client)
    * Create a simple GRPC client
- [Secure the communication with SSL](https://github.com/jakhax/grpc-gateway-example/tree/4.secure-the-communication)
    * Secure the client/server communication using SSL.
- [Simple Client Authentication](https://github.com/jakhax/grpc-gateway-example/tree/5.authenticate-client)
    * Authenticating the client to the server.
- [Adding REST support](https://github.com/jakhax/grpc-gateway-example/tree/6.add-rest-support)
    * Create a REST proxy server to add http/1.1 support, enables non-grpc compliant clients to communicate with the grpc server
- [Using a Makefile](https://github.com/jakhax/grpc-gateway-example/tree/7.using-a-makefile)
    * Using a makefile to managed the project tasks such as building binaries & installing dependencies.
- Adding SSL to REST clients
    * To add SSL between rest clients and REST proxy server to secure the communication.
