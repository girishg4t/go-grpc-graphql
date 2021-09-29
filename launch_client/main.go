// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	grpc_pb "github.com/girishg4t/grpc-test/launch_grpc"
	"google.golang.org/grpc"
)

var (
	address = os.Getenv("LAUNCH_ADDRESS")
)

func main() {
	log.Printf("Started Launces")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	lc := grpc_pb.NewLaunchServiceClient(conn)
	lr, err := lc.GetLaunches(ctx, &grpc_pb.GetLaunchesRequest{})
	if err != nil {
		log.Fatalf("could not get the launches: %v", err)
	}
	log.Printf("Launces: %s", lr.Launches)

	launchResp, err := lc.GetLaunch(ctx, &grpc_pb.GetLaunchRequest{Id: int64(100)})
	if err != nil {
		log.Fatalf("could not get the launch: %v", err)
	}
	log.Printf("Launces: %s", launchResp.Launch)
}
