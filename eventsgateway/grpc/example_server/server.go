// protos
// https://github.com/topfreegames/protos
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/topfreegames/protos/eventsgateway/grpc/generated"
	context "golang.org/x/net/context"
)

type server struct{}

func (*server) SendEvent(ctx context.Context, event *pb.Event) (*pb.Response, error) {
	fmt.Println(
		"Received Event",
		event.GetId(),
		event.GetName(),
		event.GetTopic(),
		event.GetProps(),
	)
	return &pb.Response{
		Message: "ack",
		Code:    200,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server listening at :10000")
	grpcServer := grpc.NewServer()
	pb.RegisterGRPCForwarderServer(grpcServer, &server{})

	grpcServer.Serve(lis)
}
