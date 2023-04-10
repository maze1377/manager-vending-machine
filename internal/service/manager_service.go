package service

import (
	"context"

	"github.com/maze1377/manager-vending-machine/internal/communication"
	"github.com/maze1377/manager-vending-machine/internal/storage/dbrepository"
	"github.com/maze1377/manager-vending-machine/internal/storage/entity"
	pb "github.com/maze1377/manager-vending-machine/pkg/vendingMachineService"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ManagerService struct {
	repository dbrepository.EventLogRepository
	uid        string
}

func NewManagerService(repository dbrepository.EventLogRepository, uid string) *ManagerService {
	return &ManagerService{repository: repository, uid: uid}
}

func (ms *ManagerService) AttachToMachine(ctx context.Context, configGRPC *communication.GRPCCommunicator) error {
	conn, err := configGRPC.EnsureConnection(true)
	if err != nil {
		return err
	}
	client := pb.NewVendingMachineServiceClient(conn)
	eventClient, err := client.NotifyEvent(ctx, &pb.NotifyEventRequest{Uid: ms.uid})
	if err != nil {
		log.WithError(err).Error("can not NotifyEvent")
		return err
	}
	for {
		event, err := eventClient.Recv()
		if err != nil {
			log.WithError(err).Error("error to send command")
			return status.Errorf(codes.Unknown, "error receiving command: %s", err)
		}
		log.WithField("event", event).Info("receive event")
		switch event.GetType() {
		case pb.NotifyEventResponse_EVENT_TYPE_PAYMENT_PROCESSED:
			err := ms.repository.Save(ctx, &entity.EventLog{
				Status:  entity.EventPayment,
				Message: event.GetPaymentResponse().GetMessage(),
				Success: event.GetPaymentResponse().GetSuccess(),
			})
			if err != nil {
				log.WithError(err).Error("can not save Event Payment")
				return err
			}
		case pb.NotifyEventResponse_EVENT_TYPE_PRODUCT_DISPENSED:
			err := ms.repository.Save(ctx, &entity.EventLog{
				Status:  entity.EventDispensed,
				Message: event.GetDispenseResponse().GetMessage(),
				Success: event.GetDispenseResponse().GetSuccess(),
			})
			if err != nil {
				log.WithError(err).Error("can not save Dispense Event")
				return err
			}
		case pb.NotifyEventResponse_EVENT_TYPE_UNSPECIFIED:
			return status.Error(codes.NotFound, "unspecified event type")
		default:
			log.WithField("event", event).Error("unknown event type")
			return status.Error(codes.Unimplemented, "unknown event type")
		}
	}
}
