package protocol

import (
	"crypto/tls"
	"fmt"

	"github.com/anima-protocol/anima-go/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var Client AnimaClient

type Config struct {
	Secure bool
}

// Init - Initialize New Client
func Init(config *Config, protocol *models.Protocol) (AnimaClient, error) {
	if Client == nil {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts := []grpc.DialOption{}

		if config.Secure {
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithInsecure())
		}

		cc, err := grpc.Dial(protocol.Network, opts...)
		if err != nil {
			return nil, fmt.Errorf("could not connect to GRPC Server")
		}

		Client = NewAnimaClient(cc)
	}

	return Client, nil
}
