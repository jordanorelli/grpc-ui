sample client-server pair using gRPC and Qt in Go. The server simply stores an
integer. The client, when started, automatically connects to the server and,
once a second, requests the next value in an incrementing counter. I made this
for testing the install environment when working with gRPC and Qt together in
Go, since both have a non-trivial install process.

see here to install gRPC: http://www.grpc.io/docs/quickstart/go.html
and here to install Qt: https://github.com/therecipe/qt

### Files:

- `lib`: contains the gRPC protobuf definition of our service. Our service defines one unary endpoint with an input message type and an output message type.  
- `lib/count.proto`: the gRPC definitions, written by a human  
- `lib/count`: the Go package containing our gRPC client and server definitions  
- `lib/count/count.pb.go`: generated from lib/count.proto using the following protoc invokation:  
  `protoc -I count count.proto --go_out=plugins=grpc:count`  
- `cmd`: contains our executable programs
- `cmd/count-client`: a gRPC client with a Qt ui
- `cmd/count-server`: a gRPC server, no graphical ui

### Compiling the client:

Compiling the client is a bit tricky. The way I got this to work was to use MSYS2 and MinGW-64. I installed the pacman package `mingw-w64-x86_64-qt-creator`, which includes a 64-bit gcc. I was able to piece together the install process that worked for me based on the instructions found in the https://github.com/therecipe/qt README, [the MSYS2 wiki](https://sourceforge.net/p/msys2/wiki/MSYS2%20introduction/), and [the MSYS2 page of the Qt wiki](https://wiki.qt.io/MSYS2).

_this has only been tested on Windows 10_
