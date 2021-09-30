package service

import (
	"context"
	"encoding/json"

	grpc_srv "github.com/girishg4t/grpc-test/launch_grpc"
	"github.com/machinebox/graphql"
)

type LaunchServer struct {
	grpc_srv.LaunchServiceServer
	graphqlClient *graphql.Client
}

// NewLaunchService get's the new instance of launch service
func NewLaunchService(client *graphql.Client) LaunchServer {
	return LaunchServer{
		graphqlClient: client,
	}
}

// GetLaunches get the launches by limit
func (s *LaunchServer) GetLaunches(ctx context.Context, in *grpc_srv.GetLaunchesRequest) (*grpc_srv.GetLaunchesResponse, error) {
	req := graphql.NewRequest(`
	    query ($key: Int!) {
			launchesPast(limit: $key) {
				id
				is_tentative
				launch_success
				upcoming
				mission_name
			}
		}
    `)
	req.Var("key", int(in.Limit))
	var graphqlResponse interface{}
	if err := s.graphqlClient.Run(context.Background(), req, &graphqlResponse); err != nil {
		return nil, err
	}
	var launchpast map[string][]*grpc_srv.Launch
	bytes, _ := json.Marshal(graphqlResponse)
	_ = json.Unmarshal(bytes, &launchpast)
	launches := launchpast["launchesPast"]
	return &grpc_srv.GetLaunchesResponse{Launches: launches}, nil
}

// GetLaunch get launch based on id
func (s *LaunchServer) GetLaunch(ctx context.Context, in *grpc_srv.GetLaunchRequest) (*grpc_srv.GetLaunchResponse, error) {
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
		return nil, err
	}
	var launchpast map[string]*grpc_srv.Launch
	bytes, _ := json.Marshal(graphqlResponse)
	_ = json.Unmarshal(bytes, &launchpast)
	launch := launchpast["launch"]
	return &grpc_srv.GetLaunchResponse{Launch: launch}, nil
}
