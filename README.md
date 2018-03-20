# Micro Path Example

This repo is meant to be an example of a suspected bug in `go-micro`'s gRPC implmentation.

A Hello World service is defined in `proto/hello.proto` and compiled using the protoc micro plugin.

A server that implements the service is written in Go using `go-micro` and `micro/go-grpc`.

A client that consumes the service is written in node.js using the `grpc-node` library.

## Set up

It assumed that glide, protobuf, node, and yarn are in your path. If running OS X `make prereqs` will install any missing tools via `brew`

All dependencies should be installed by running `make config-local`. This will:
  - go get micro/micro, micro's server and client grpc plugins, micro's kubernetes registry plugin, and micro's proto compiler 
  - build micro with grpc and kubernetes enabled as detailed [here](https://github.com/micro/go-grpc#build-yourself)
  - install go dependencies with glide
  - install node dependencies with yarn


## Server 

Implements the `Greet` method by always returning a greeting string.

## Client

Instantiates 2 clients by consuming the proto file.

One is instantiated the way the library is intended to be used. The other modifies the path for each method in the service definition by changing `/Service/Method` to `Service.Method`. 

They both then try to consume the service. The default client should error with `malformed method name` while the modified client should receive the greeting as intended.

## Running the example

Build the server with `make build` and run it with `make`. Make a note which port the server is listening on.

Then run the client with `SERVICE_ENDPOINT=localhost:<SERVER PORT> make run-client`.

