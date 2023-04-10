package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/maze1377/manager-vending-machine/internal/command"
	"github.com/maze1377/manager-vending-machine/internal/machine"
	"github.com/maze1377/manager-vending-machine/internal/models"
	pb "github.com/maze1377/manager-vending-machine/pkg/vendingMachineService"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MachineService struct {
	vm *machine.VendingMachine
}

func (m *MachineService) GetProduct(ctx context.Context, _ *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	select {
	case <-ctx.Done():
		st, ok := status.FromError(ctx.Err())
		if ok {
			return nil, st.Err()
		}
		return nil, status.Error(codes.Canceled, "ctx canceled")
	default:
		products := m.vm.GetProducts()
		serializedProduct := make([]*pb.Product, len(products))
		for i, product := range products {
			serializedProduct[i] = &pb.Product{
				Name:     product.Name,
				Quantity: product.Quantity,
				Price:    product.Price,
			}
		}
		return &pb.GetProductResponse{
			Products: serializedProduct,
		}, nil
	}
}

func (m *MachineService) ExecuteCommand(server pb.VendingMachineService_ExecuteCommandServer) error {
	for {
		cmd, err := server.Recv()
		if err != nil {
			return status.Errorf(codes.Unknown, "error receiving command: %s", err)
		}
		uid := cmd.GetUid()
		switch cmd.GetType() {
		case pb.ExecuteCommandRequest_COMMAND_TYPE_ADD_PRODUCT:
			product := cmd.GetProduct()
			err := command.NewAddProductCommand(uid, &models.Product{
				Name:     product.GetName(),
				Price:    product.GetPrice(),
				Quantity: product.GetQuantity(),
			}).Execute(m.vm)
			if err != nil {
				err2 := server.Send(&pb.ExecuteCommandResponse{
					Message: fmt.Sprintf("can not add new product :%s", err),
					Success: false,
				})
				if err2 != nil {
					log.WithError(err2).Error("error to send command")
					return status.Errorf(codes.Unknown, "error to send command: %s", err2)
				}
				continue
			}
			err2 := server.Send(&pb.ExecuteCommandResponse{
				Message: "new product added",
				Success: true,
			})
			if err2 != nil {
				log.WithError(err2).Error("error to send command")
				return status.Errorf(codes.Unknown, "error to send command: %s", err2)
			}
		case pb.ExecuteCommandRequest_COMMAND_TYPE_SELECT_PRODUCT:
			productName := cmd.GetSelectProductRequest().GetProductName()
			err := command.NewSelectProductCommand(uid, productName).Execute(m.vm)
			if err != nil { // todo handle more clear and better such as ErrProductRunningOut and...
				log.WithError(err).Warning("can select product")
				err2 := server.Send(&pb.ExecuteCommandResponse{
					Message: fmt.Sprintf("can select product :%s", err),
					Success: false,
				})
				if err2 != nil {
					log.WithError(err2).Error("error to send command")
					return status.Errorf(codes.Unknown, "error to send command: %s", err2)
				}
				continue
			}
			err2 := server.Send(&pb.ExecuteCommandResponse{
				Message: "product selected",
				Success: true,
			})
			if err2 != nil {
				log.WithError(err2).Error("error to send command")
				return status.Errorf(codes.Unknown, "error to send command: %s", err2)
			}
		case pb.ExecuteCommandRequest_COMMAND_TYPE_DISPENSE_PRODUCT:
			productName := cmd.GetDispenseRequest().GetProductName()
			err := command.NewDispenseProductCommand(uid, productName).Execute(m.vm)
			if err != nil { // todo error handle
				log.WithError(err).Warning("can dispense product")
				err2 := server.Send(&pb.ExecuteCommandResponse{
					Message: fmt.Sprintf("can dispense product :%s", err),
					Success: false,
				})
				if err2 != nil {
					log.WithError(err2).Error("error to send command")
					return status.Errorf(codes.Unknown, "error to send command: %s", err2)
				}
				continue
			}
			err2 := server.Send(&pb.ExecuteCommandResponse{
				Message: "dispense product",
				Success: true,
			})
			if err2 != nil {
				log.WithError(err2).Error("error to send command")
				return status.Errorf(codes.Unknown, "error to send command: %s", err2)
			}
		case pb.ExecuteCommandRequest_COMMAND_TYPE_PROCESS_PAYMENT:
			paymentRequest := cmd.GetPaymentRequest()
			err := command.NewProcessPaymentCommand(uid,
				paymentRequest.GetPaymentMethod(),
				paymentRequest.GetAmount()).Execute(m.vm)
			if err != nil { // todo error handle
				log.WithError(err).Warning("can process payment")
				err2 := server.Send(&pb.ExecuteCommandResponse{
					Message: fmt.Sprintf("can process payment :%s", err),
					Success: false,
				})
				if err2 != nil {
					log.WithError(err2).Error("error to send command")
					return status.Errorf(codes.Unknown, "error to send command: %s", err2)
				}
				continue
			}
			err2 := server.Send(&pb.ExecuteCommandResponse{
				Message: "process payment",
				Success: true,
			})
			if err2 != nil {
				log.WithError(err2).Error("error to send command")
				return status.Errorf(codes.Unknown, "error to send command: %s", err2)
			}
		case pb.ExecuteCommandRequest_COMMAND_TYPE_UNSPECIFIED:
			return status.Error(codes.NotFound, "unspecified command type")
		default:
			return status.Error(codes.Unimplemented, "unknown command type")
		}
	}
}

func (m *MachineService) NotifyEvent(request *pb.NotifyEventRequest, server pb.VendingMachineService_NotifyEventServer) error {
	uid := request.GetUid()
	done := make(chan error)
	m.vm.AddObserver(uid, func(event models.Event, date ...interface{}) {
		switch {
		case event == models.Payment:
			err := server.Send(&pb.NotifyEventResponse{
				Type: pb.NotifyEventResponse_EVENT_TYPE_PAYMENT_PROCESSED,
				Payload: &pb.NotifyEventResponse_PaymentResponse{
					PaymentResponse: &pb.PaymentResponse{ // todo clean this part
						Success: date[0].(bool),
						Message: date[1].(string),
					},
				},
			})
			if err != nil {
				done <- err
				return
			}
		case event == models.Dispensed:
			err := server.Send(&pb.NotifyEventResponse{
				Type: pb.NotifyEventResponse_EVENT_TYPE_PRODUCT_DISPENSED,
				Payload: &pb.NotifyEventResponse_DispenseResponse{
					DispenseResponse: &pb.DispenseResponse{ // todo clean this part
						Success: date[0].(bool),
						Message: date[1].(string),
					},
				},
			})
			if err != nil {
				done <- err
				return
			}
		default:
			done <- errors.New("not support event type")
		}
	})
	defer func() {
		m.vm.RemoveObserver(uid)
	}()
	err := <-done
	return status.Errorf(codes.Unknown, "error to send notification: %s", err)
}

func NewMachineService(vm *machine.VendingMachine) *MachineService {
	return &MachineService{vm: vm}
}
