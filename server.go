package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/NoahOrberg/prac_grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type UserAgentServer struct {
}

func NewUserAgentServer() *UserAgentServer {
	return &UserAgentServer{}
}

func (s UserAgentServer) GetUA(ctx context.Context, req *proto.UserAgentReq) (*proto.UserAgentRes, error) {
	log.Print(ctx)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &proto.UserAgentRes{}, grpc.Errorf(codes.Internal, "Cannot get user-agent")
	}

	return &proto.UserAgentRes{
		UserAgent: md["user-agent"][0],
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUserAgentServer(s, NewUserAgentServer())
	s.Serve(lis)
}
