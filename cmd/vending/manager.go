package main

import (
	"context"
	"sync"
	"time"

	"github.com/maze1377/manager-vending-machine/config"
	"github.com/maze1377/manager-vending-machine/internal/communication"
	"github.com/maze1377/manager-vending-machine/internal/metrics"
	"github.com/maze1377/manager-vending-machine/internal/service"
	"github.com/maze1377/manager-vending-machine/internal/storage"
	"github.com/maze1377/manager-vending-machine/internal/storage/dbrepository"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "run manager :)",
	Run:   RunManager,
}

func init() {
	rootCmd.AddCommand(managerCmd)
}

func RunManager(_ *cobra.Command, _ []string) {
	simpleMetric, err := metrics.NewSimpleMetric(config.Instance.MetricAddr)
	if err != nil {
		log.Panic("can not initialize metric")
	}

	// Require running migration every time due to use of in-memory database.
	db := storage.Migrate()
	repository := dbrepository.NewEventLog(db,
		simpleMetric.NewCommunicator(metrics.CommunicatorConfig{
			Name: "repository_eventLog",
			Help: "Communicator metric for db repository eventLog",
		}))

	// todo read from config
	uid := "cmd-runner"
	configGRPC := &communication.GRPCCommunicator{
		URL:         "localhost:10000",
		ServiceName: "first machine",
		Redial:      1,
		Insecure:    true,
		Lock:        &sync.Mutex{},
	}
	managerService := service.NewManagerService(repository, uid)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		err := managerService.AttachToMachine(ctx, configGRPC)
		if err != nil {
			log.WithError(err).Error("can not attach to machine")
		}
	}()
	handleSignals(func() {
		done <- struct{}{}
		defer cancel()
	})
	<-done
	<-time.After(10 * time.Second)
}
