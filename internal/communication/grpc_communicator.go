package communication

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type GRPCCommunicator struct {
	Conn        *grpc.ClientConn
	Lock        *sync.Mutex
	URL         string
	CertFile    string
	ServiceName string
	Redial      int
	Insecure    bool
}

func (gc *GRPCCommunicator) createConnection() (*grpc.ClientConn, error) {
	resolver.SetDefaultScheme("dns")
	grpcDialOpts := []grpc.DialOption{
		grpc.WithDisableServiceConfig(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	}
	if gc.Insecure {
		grpcDialOpts = append(grpcDialOpts,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
	} else {
		credential, err := credentials.NewClientTLSFromFile(gc.CertFile, "localhost")
		if err != nil {
			logrus.WithError(err).Errorf("failed to load tls from %s", gc.CertFile)
			return nil, fmt.Errorf("[createConnection] credentials.NewClientTLSFromFile error: %w", err)
		}
		grpcDialOpts = append(grpcDialOpts,
			grpc.WithTransportCredentials(credential),
		)
	}

	conn, err := grpc.Dial(gc.URL, grpcDialOpts...)
	if err != nil {
		return nil, fmt.Errorf("dial %s failed, %v", gc.ServiceName, err)
	}
	return conn, nil
}

func (gc *GRPCCommunicator) shouldRedial() bool {
	return rand.Intn(10000) < gc.Redial
}

func (gc *GRPCCommunicator) EnsureConnection(forceRedial bool) (*grpc.ClientConn, error) {
	if gc.Conn != nil && !forceRedial && !gc.shouldRedial() {
		return gc.Conn, nil
	}
	gc.Lock.Lock()
	defer gc.Lock.Unlock()
	newConn, err := gc.createConnection()
	if err != nil {
		return nil, fmt.Errorf("GRPC connection does not exists %v", err)
	}
	if gc.Conn != nil {
		go gc.closeConn(gc.Conn)
	}
	gc.Conn = newConn
	return newConn, nil
}

func (gc *GRPCCommunicator) closeConn(closingConn *grpc.ClientConn) {
	closingTimeout := 10 * time.Second
	time.Sleep(closingTimeout)

	err := closingConn.Close()
	if err != nil {
		logrus.WithError(err).Error("Error in closing grpc connection")
	}
}
