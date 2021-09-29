package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"

	grpc_srv "github.com/girishg4t/grpc-test/launch_grpc"
	launch "github.com/girishg4t/grpc-test/launch_grpc"
	"github.com/machinebox/graphql"
	"google.golang.org/grpc"
)

var (
	port = ":" + os.Getenv("SERVER_PORT")
)

type server struct {
	grpc_srv.LaunchServiceServer
	graphqlClient *graphql.Client
}

func (s *server) GetLaunches(ctx context.Context, in *grpc_srv.GetLaunchesRequest) (*grpc_srv.GetLaunchesResponse, error) {
	graphqlRequest := graphql.NewRequest(`
        {
			launchesPast(limit: 10) {
				id
				is_tentative
				launch_success
				upcoming
				mission_name
			}
		}
    `)
	var graphqlResponse interface{}
	if err := s.graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	var launchpast map[string][]*grpc_srv.Launch
	bytes, _ := json.Marshal(graphqlResponse)
	_ = json.Unmarshal(bytes, &launchpast)
	launches := launchpast["launchesPast"]
	return &grpc_srv.GetLaunchesResponse{Launches: launches}, nil
}

func (s *server) GetLaunch(ctx context.Context, in *launch.GetLaunchRequest) (*launch.GetLaunchResponse, error) {
	req := graphql.NewRequest(`
		query ($key: ID!) {
			launch(id: $key) {
				id
				is_tentative
				launch_success
				upcoming
				mission_name
			}
		}
	`)
	req.Var("key", int(in.Id))
	var graphqlResponse interface{}
	if err := s.graphqlClient.Run(context.Background(), req, &graphqlResponse); err != nil {
		panic(err)
	}
	var launchpast map[string]*grpc_srv.Launch
	bytes, _ := json.Marshal(graphqlResponse)
	_ = json.Unmarshal(bytes, &launchpast)
	launch := launchpast["launch"]
	return &grpc_srv.GetLaunchResponse{Launch: launch}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	graphqlClient := graphql.NewClient(os.Getenv("GRAPHQL_URL"))
	s := grpc.NewServer()

	grpc_srv.RegisterLaunchServiceServer(s, &server{
		graphqlClient: graphqlClient,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
