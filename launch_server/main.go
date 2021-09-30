package main

import (
	"log"
	"net"
	"os"

	grpc_srv "github.com/girishg4t/grpc-test/launch_grpc"
	"github.com/girishg4t/grpc-test/launch_server/service"
	"github.com/machinebox/graphql"
	"google.golang.org/grpc"
)

var (
	port = ":" + os.Getenv("SERVER_PORT")
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	graphqlClient := graphql.NewClient(os.Getenv("GRAPHQL_URL"))

	s := grpc.NewServer()

	launchServerClient := service.NewLaunchService(graphqlClient)

	grpc_srv.RegisterLaunchServiceServer(s, &launchServerClient)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
