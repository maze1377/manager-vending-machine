package integration

import (
	"context"
	"net"
	"testing"

	"github.com/maze1377/manager-vending-machine/config"
	"github.com/maze1377/manager-vending-machine/internal/machine"
	"github.com/maze1377/manager-vending-machine/internal/models"
	"github.com/maze1377/manager-vending-machine/internal/service"
	pb "github.com/maze1377/manager-vending-machine/pkg/vendingMachineService"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

type machineTestSuite struct {
	suite.Suite
}

func TestIntegrationMachine(t *testing.T) {
	suite.Run(t, new(machineTestSuite))
}

func (s *machineTestSuite) TestMachine() {
	// initialize server for testing

	ctx := context.Background()
	err := config.NewConfig("")
	if !s.NoError(err) {
		s.FailNow("can't init config")
	}
	products := []*models.Product{
		{Name: "Coke", Price: 50, Quantity: 5},
		{Name: "Pepsi", Price: 60, Quantity: 3},
		{Name: "Sprite", Price: 40, Quantity: 0},
	}
	vm := machine.NewVendingMachine(products)
	machineService := service.NewMachineService(vm)
	// initialize server for testing
	lis := bufconn.Listen(bufSize)
	gRPCServer := grpc.NewServer()
	pb.RegisterVendingMachineServiceServer(gRPCServer, machineService)
	go func() {
		if err2 := gRPCServer.Serve(lis); err2 != nil {
			s.FailNow("Server exited with error: %v", err2)
		}
	}()

	conn, err := grpc.DialContext(ctx, "buff net", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.Failf("Failed to dial buff net: %v", err.Error())
	}
	defer conn.Close()
	client := pb.NewVendingMachineServiceClient(conn)

	// start testing
	// get product
	resGetProduct, err := client.GetProduct(ctx, &pb.GetProductRequest{})
	s.NoError(err)
	s.Len(resGetProduct.GetProducts(), 3)
	// send cmd payment
	cmdService, err := client.ExecuteCommand(ctx)
	s.NoError(err)
	paymentRequest := &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_PROCESS_PAYMENT,
		Payload: &pb.ExecuteCommandRequest_PaymentRequest{
			PaymentRequest: &pb.PaymentRequest{
				Amount:        50,
				PaymentMethod: "test-payment",
			},
		},
	}
	err = cmdService.Send(paymentRequest)
	s.NoError(err)
	paymentResponse, err := cmdService.Recv()
	s.NoError(err)
	s.True(paymentResponse.GetSuccess())
	// now other user try to pay
	paymentRequest = &pb.ExecuteCommandRequest{
		Uid:  "test-2",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_PROCESS_PAYMENT,
		Payload: &pb.ExecuteCommandRequest_PaymentRequest{
			PaymentRequest: &pb.PaymentRequest{
				Amount:        40,
				PaymentMethod: "test-payment",
			},
		},
	}
	err = cmdService.Send(paymentRequest)
	s.NoError(err)

	paymentResponse, err = cmdService.Recv()
	s.NoError(err)
	s.False(paymentResponse.GetSuccess())
	s.Contains(paymentResponse.GetMessage(), machine.ErrMachineBusyNow.Error())
	// now do dispense product that not valid transaction
	dispenseRequest := &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_DISPENSE_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_DispenseRequest{
			DispenseRequest: &pb.DispenseRequest{
				ProductName: "Coke",
			},
		},
	}
	err = cmdService.Send(dispenseRequest)
	s.NoError(err)

	dispenseResponse, err := cmdService.Recv()
	s.NoError(err)
	s.False(dispenseResponse.GetSuccess())
	s.Contains(dispenseResponse.GetMessage(), machine.ErrTransactionNotValid.Error())
	// now select item that running out
	selectProductRequest := &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_SELECT_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_SelectProductRequest{
			SelectProductRequest: &pb.SelectProductRequest{
				ProductName: "Sprite",
			},
		},
	}
	err = cmdService.Send(selectProductRequest)
	s.NoError(err)

	selectProductResponse, err := cmdService.Recv()
	s.NoError(err)
	s.False(selectProductResponse.GetSuccess())
	s.Contains(selectProductResponse.GetMessage(), machine.ErrProductRunningOut.Error())
	// now select item that not found
	selectProductRequest = &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_SELECT_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_SelectProductRequest{
			SelectProductRequest: &pb.SelectProductRequest{
				ProductName: "test-item",
			},
		},
	}
	err = cmdService.Send(selectProductRequest)
	s.NoError(err)

	selectProductResponse, err = cmdService.Recv()
	s.NoError(err)
	s.False(selectProductResponse.GetSuccess())
	s.Contains(selectProductResponse.GetMessage(), machine.ErrProductNotFound.Error())
	// now select item that not enough money
	selectProductRequest = &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_SELECT_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_SelectProductRequest{
			SelectProductRequest: &pb.SelectProductRequest{
				ProductName: "Pepsi",
			},
		},
	}
	err = cmdService.Send(selectProductRequest)
	s.NoError(err)

	selectProductResponse, err = cmdService.Recv()
	s.NoError(err)
	s.False(selectProductResponse.GetSuccess())
	s.Contains(selectProductResponse.GetMessage(), machine.ErrNotEnoughMoney.Error())
	// now select correct item
	selectProductRequest = &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_SELECT_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_SelectProductRequest{
			SelectProductRequest: &pb.SelectProductRequest{
				ProductName: "Coke",
			},
		},
	}
	err = cmdService.Send(selectProductRequest)
	s.NoError(err)

	selectProductResponse, err = cmdService.Recv()
	s.NoError(err)
	s.True(selectProductResponse.GetSuccess())
	// now dispose correct item
	dispenseRequest = &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_DISPENSE_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_DispenseRequest{
			DispenseRequest: &pb.DispenseRequest{
				ProductName: "Coke",
			},
		},
	}
	err = cmdService.Send(dispenseRequest)
	s.NoError(err)
	selectProductResponse, err = cmdService.Recv()
	s.NoError(err)
	s.True(selectProductResponse.GetSuccess())
}

func (s *machineTestSuite) SetupTest() {}

func (s *machineTestSuite) TearDownTest() {}
