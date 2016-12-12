package main

import (
    "net"
    "sync"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
    "github.com/jordanorelli/grpc-ui/lib/count"
)

type countServer struct {
    sync.Mutex
    last int64
}

func (c *countServer) Next(ctx context.Context, r *count.NextRequest) (*count.NextReply, error) {
    return &count.NextReply{Val: c.incr()}, nil
}

func (c *countServer) incr() int64 {
    c.Lock()
    defer c.Unlock()

    c.last += 1
    return c.last
}

func main() {
    lis, err := net.Listen("tcp", "localhost:9001")
    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()
    count.RegisterCountServer(s, &countServer{})
    s.Serve(lis)
}