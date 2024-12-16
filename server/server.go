package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bufbuild/protovalidate-go"
	pb "github.com/madhab452/gt-store/pstore/gen-go"
	"github.com/madhab452/gt-store/server/internal"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50090, "The server port")
)

type server struct {
	pb.UnimplementedStoreServer
	ms *internal.MemoryStore
}

func (s *server) Get(_ context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	if err := protovalidate.Validate(in); err != nil {
		return nil, err
	}

	v, err := s.ms.Get(in.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetReply{Value: v}, nil
}

func (s *server) Put(_ context.Context, in *pb.PutRequest) (*pb.GetReply, error) {
	if err := protovalidate.Validate(in); err != nil {
		return nil, err
	}

	err := s.ms.Put(in.Key, in.Value)
	if err != nil {
		return nil, err
	}

	return &pb.GetReply{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := &server{
		ms: internal.NewMemoryStore(),
	}
	pb.RegisterStoreServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
