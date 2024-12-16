package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/madhab452/gt-store/pstore/gen-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", ":50090", "the address to connect to")
	key  = flag.String("key", "", "the key to lookfor")
	val  = flag.String("value", "", "the value to store")
	rpc  = flag.String("rpc", "get", "alloweed are get or put ")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	switch *rpc {
	case "get":
		r, err := c.Get(ctx, &pb.GetRequest{
			Key: *key,
		})
		if err != nil {
			log.Fatalf("could not Get: %v", err)
		}
		fmt.Println("value received ", r.GetValue())
	case "put":
		_, err := c.Put(ctx, &pb.PutRequest{
			Key:   *key,
			Value: *val,
		})
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("value set:", *val)
		return
	default:
		fmt.Print("unknown rpc!!")
	}

}
