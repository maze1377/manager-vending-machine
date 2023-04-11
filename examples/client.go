//go:build !test

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	pb "github.com/maze1377/manager-vending-machine/pkg/vendingMachineService"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:10000", "address of machine backend")
	flag.Parse()

	fmt.Printf("address of machine backend %s", *backend)

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewVendingMachineServiceClient(conn)

	callExecuteCommandRequestAddProduct(client)
	// callExecuteCommandRequestProcessPayment(client)
	// callExecuteCommandRequestSelectItem(client)
	// callExecuteCommandRequestDispenseProduct(client)
}

func callGetProductRequest(client pb.VendingMachineServiceClient) {
	params := &pb.GetProductRequest{}

	res, err := client.GetProduct(context.Background(), params)
	if err != nil {
		log.Fatalf("could not call GetProductRequest %v", err)
	}
	fmt.Println("input:")
	spew.Dump(params)
	fmt.Println("output:")
	spew.Dump(res)
}

func callExecuteCommandRequestAddProduct(client pb.VendingMachineServiceClient) {
	params := &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_ADD_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_Product{
			Product: &pb.Product{
				Name:     "test-product",
				Quantity: 10,
				Price:    100.0,
			},
		},
	}
	service, err := client.ExecuteCommand(context.Background())
	if err != nil {
		log.Fatalf("could not call ExecuteCommand %v", err)
	}

	err = service.Send(params)
	if err != nil {
		log.Fatalf("could not call send %v", err)
	}

	res, err := service.Recv()
	if err != nil {
		log.Fatalf("could not call recv %v", err)
	}

	fmt.Println("input:")
	spew.Dump(params)
	fmt.Println("output:")
	spew.Dump(res)
}

func callExecuteCommandRequestProcessPayment(client pb.VendingMachineServiceClient) {
	params := &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_PROCESS_PAYMENT,
		Payload: &pb.ExecuteCommandRequest_PaymentRequest{
			PaymentRequest: &pb.PaymentRequest{
				Amount:        50,
				PaymentMethod: "test-payment",
			},
		},
	}
	service, err := client.ExecuteCommand(context.Background())
	if err != nil {
		log.Fatalf("could not call ExecuteCommand %v", err)
	}

	err = service.Send(params)
	if err != nil {
		log.Fatalf("could not call send %v", err)
	}

	res, err := service.Recv()
	if err != nil {
		log.Fatalf("could not call recv %v", err)
	}

	fmt.Println("input:")
	spew.Dump(params)
	fmt.Println("output:")
	spew.Dump(res)
}

func callExecuteCommandRequestSelectItem(client pb.VendingMachineServiceClient) {
	params := &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_SELECT_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_SelectProductRequest{
			SelectProductRequest: &pb.SelectProductRequest{
				ProductName: "Coke",
			},
		},
	}
	service, err := client.ExecuteCommand(context.Background())
	if err != nil {
		log.Fatalf("could not call ExecuteCommand %v", err)
	}

	err = service.Send(params)
	if err != nil {
		log.Fatalf("could not call send %v", err)
	}

	res, err := service.Recv()
	if err != nil {
		log.Fatalf("could not call recv %v", err)
	}

	fmt.Println("input:")
	spew.Dump(params)
	fmt.Println("output:")
	spew.Dump(res)
}

func callExecuteCommandRequestDispenseProduct(client pb.VendingMachineServiceClient) {
	params := &pb.ExecuteCommandRequest{
		Uid:  "test",
		Type: pb.ExecuteCommandRequest_COMMAND_TYPE_DISPENSE_PRODUCT,
		Payload: &pb.ExecuteCommandRequest_DispenseRequest{
			DispenseRequest: &pb.DispenseRequest{
				ProductName: "Coke",
			},
		},
	}
	service, err := client.ExecuteCommand(context.Background())
	if err != nil {
		log.Fatalf("could not call ExecuteCommand %v", err)
	}

	err = service.Send(params)
	if err != nil {
		log.Fatalf("could not call send %v", err)
	}

	res, err := service.Recv()
	if err != nil {
		log.Fatalf("could not call recv %v", err)
	}

	fmt.Println("input:")
	spew.Dump(params)
	fmt.Println("output:")
	spew.Dump(res)
}
