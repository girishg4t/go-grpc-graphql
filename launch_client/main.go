// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/girishg4t/grpc-test/launch_client/service"
	grpc_pb "github.com/girishg4t/grpc-test/launch_grpc"
	"google.golang.org/grpc"
)

var (
	address = os.Getenv("LAUNCH_ADDRESS")
)

func main() {
	log.Printf("Started Launces")

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	processLaunch(ctx, conn)
}

func processLaunch(ctx context.Context, conn *grpc.ClientConn) {

	lc := grpc_pb.NewLaunchServiceClient(conn)
	var ls service.ILaunchService
	srv := service.NewLaunchService(lc, ls)

	limit, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		limit = 5
	}
	err = srv.GetLaunches(ctx, limit)

	if err != nil {
		log.Fatalf("error in processing launches %v", err)
	}

	id, err := strconv.ParseInt(flag.Arg(1), 10, 64)
	if err != nil {
		id = 1
	}
	err = srv.GetLaunch(ctx, id)

	if err != nil {
		log.Fatalf("error in processing launch %v", err)
	}

}
