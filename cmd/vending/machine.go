package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLoggers "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	gRpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/maze1377/manager-vending-machine/config"
	"github.com/maze1377/manager-vending-machine/internal/machine"
	"github.com/maze1377/manager-vending-machine/internal/models"
	"github.com/maze1377/manager-vending-machine/internal/service"
	pb "github.com/maze1377/manager-vending-machine/pkg/vendingMachineService"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "run vending machine :)",
	Run:   RunMachine,
}

func init() {
	rootCmd.AddCommand(machineCmd)
}

func RunMachine(_ *cobra.Command, _ []string) {
	// todo read product from configmap
	products := []*models.Product{
		{Name: "Coke", Price: 50, Quantity: 5},
		{Name: "Pepsi", Price: 60, Quantity: 3},
		{Name: "Sprite", Price: 40, Quantity: 0},
	}
	vm := machine.NewVendingMachine(products)
	machineService := service.NewMachineService(vm)

	addr := config.Instance.Address
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.WithError(err).Fatalf("Failed to listen on %q", addr)
	}

	var opts []grpc.ServerOption
	logEntry := log.WithFields(map[string]interface{}{
		"app": "machine",
	})
	opts = append(
		opts,
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcLoggers.UnaryServerInterceptor(logEntry),
				gRpcPrometheus.UnaryServerInterceptor,
				grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandlerContext(
					func(ctx context.Context, p interface{}) error {
						log.Errorf("[PANIC] %s\n\n%s", p, string(debug.Stack()))
						return status.Errorf(codes.Internal, "%s", p)
					})),
			),
		),
	)
	gRPCServer := grpc.NewServer(opts...)
	pb.RegisterVendingMachineServiceServer(gRPCServer, machineService)
	handleSignals(func() {
		go gRPCServer.GracefulStop()
	})

	if err = gRPCServer.Serve(conn); err != nil {
		log.WithError(err).Fatal("Error while apiHandler.Serve()")
	}
}

func handleSignals(shutdown func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			log.Warnf("Received sig-%s", sig.String())

			// keeping server responsive for a short amount of time
			// after receiving a soft signal. This is due to a known
			// issue of Kubernetes that leads to not marking pod unready
			// in cases of termination.
			<-time.After(3 * time.Second)
			shutdown()
			<-time.After(7 * time.Second)
			break
		}
	}()
}
