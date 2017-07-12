package main

import (
	"flag"
	"log"

	"github.com/NoahOrberg/prac_grpc/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	addrFlag = flag.String("addr", "localhost:50052", "server address host:post")
)

func main() {
	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	defer conn.Close()

	c := proto.NewUserAgentClient(conn)

	resp, err := c.GetUA(context.Background(), &proto.UserAgentReq{})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Printf("\nUserAgent: `%s`", resp.UserAgent)
}
