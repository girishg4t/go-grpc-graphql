package service_test

import (
	"context"

	grpc_srv "github.com/girishg4t/grpc-test/launch_grpc"
	"github.com/girishg4t/grpc-test/launch_server/service"
	"github.com/machinebox/graphql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Notification", func() {
	var srv service.LaunchServer
	var ctx context.Context
	var fakerGraphqlClient *graphql.Client
	// var fakerClient *launch_grpcfakes.FakeLaunchServiceClient
	BeforeEach(func() {
		ctx = context.Background()

		fakerGraphqlClient = graphql.NewClient("")
		srv = service.NewLaunchService(fakerGraphqlClient)
	})
	AfterEach(func() {

	})
	Context("Check launch", func() {
		It("Negative: returns error if no launch are set", func() {
			_, err := srv.GetLaunch(ctx, &grpc_srv.GetLaunchRequest{Id: 1})
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Check launches", func() {
		It("Negative: returns error if no launch are set", func() {
			_, err := srv.GetLaunches(ctx, &grpc_srv.GetLaunchesRequest{Limit: 1})
			Expect(err).Should(HaveOccurred())
		})
	})
})
