package main

import (
	"net"
	"net/http"
	"strings"

	"github.com/jonathabp/grpc-leak/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"

	_ "expvar"
	_ "net/http/pprof"
)

//go:generate protoc -I ./../proto ./../proto/leak.proto --go_out=plugins=grpc:./../proto

var data = strings.Repeat("hello", 1024*1024)

type server struct{}

func (s *server) Get(ctx context.Context, in *proto.DataRequest) (*proto.DataReply, error) {
	return &proto.DataReply{Data: data}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	// pprof
	go func() {
		if err := http.ListenAndServe(":8001", nil); err != nil {
			panic(err)
		}
	}()

	// expvars
	go func() {
		if err := http.ListenAndServe(":8002", http.DefaultServeMux); err != nil {
			panic(err)
		}
	}()

	s := grpc.NewServer()
	proto.RegisterDataServer(s, &server{})

	s.Serve(lis)
}
