package service

import (
	"context"
	"log"

	grpc_pb "github.com/girishg4t/grpc-test/launch_grpc"
)

type LaunchClient struct {
	launchClient grpc_pb.LaunchServiceClient
	service      ILaunchService
}

type ILaunchService interface {
	GetLaunch(ctx context.Context, id int64) error
	GetLaunches(ctx context.Context, limit int64) error
}

// NewLaunchService get new instance of launch client service
func NewLaunchService(client grpc_pb.LaunchServiceClient, service ILaunchService) LaunchClient {
	return LaunchClient{
		launchClient: client,
		service:      service,
	}
}

// GetLaunch get the launch by id default it to 1
func (l LaunchClient) GetLaunch(ctx context.Context, id int64) error {
	launchResp, err := l.launchClient.GetLaunch(ctx, &grpc_pb.GetLaunchRequest{Id: id})
	if err != nil {
		return err
	}
	log.Printf("Launce received for id %d as %v", id, launchResp.Launch)
	return nil
}

// GetLaunches get the number of launches default it to 5
func (l LaunchClient) GetLaunches(ctx context.Context, limit int64) error {
	lr, err := l.launchClient.GetLaunches(ctx, &grpc_pb.GetLaunchesRequest{Limit: limit})
	if err != nil {
		return err
	}
	log.Printf("total %d launches received\n", limit)
	log.Printf("Launces: %s", lr.Launches)
	return nil
}
