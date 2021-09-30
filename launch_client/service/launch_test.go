package service_test

import (
	"context"
	"errors"

	"github.com/girishg4t/grpc-test/launch_client/service"
	launch "github.com/girishg4t/grpc-test/launch_grpc"
	"github.com/girishg4t/grpc-test/launch_grpc/launch_grpcfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Launch", func() {

	var srv service.ILaunchService
	var ctx context.Context
	var fakerClient *launch_grpcfakes.FakeLaunchServiceClient
	var data launch.Launch
	BeforeEach(func() {
		ctx = context.Background()
		fakerClient = new(launch_grpcfakes.FakeLaunchServiceClient)
		srv = service.NewLaunchService(fakerClient, srv)
		data = launch.Launch{
			MissionName: "mangal",
		}
	})
	AfterEach(func() {

	})
	Context("Check launch", func() {
		It("Negative: returns error if no launch are set", func() {
			fakerClient.GetLaunchReturns(nil, errors.New("no launch"))
			err := srv.GetLaunch(ctx, 1)
			Expect(err).Should(HaveOccurred())
		})
		It("Positive: returns launch", func() {
			fakerClient.GetLaunchReturns(&launch.GetLaunchResponse{
				Launch: &data,
			}, nil)
			err := srv.GetLaunch(ctx, 1)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Check launches", func() {
		It("Negative: returns error if no launch are set", func() {
			fakerClient.GetLaunchesReturns(nil, errors.New("no launches"))
			err := srv.GetLaunches(ctx, 2)
			Expect(err).Should(HaveOccurred())
		})
		It("Positive: returns launch", func() {
			fakerClient.GetLaunchesReturns(&launch.GetLaunchesResponse{
				Launches: []*launch.Launch{&data},
			}, nil)
			err := srv.GetLaunches(ctx, 1)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
