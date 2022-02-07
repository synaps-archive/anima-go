package protocol

import (
	context "context"

	"github.com/anima-protocol/anima-go/chains/ethereum"
	"github.com/anima-protocol/anima-go/models"
	"google.golang.org/grpc/metadata"
)

func Issue(anima *models.Protocol, req *IssueRequest) error {
	config := &Config{Secure: false}
	client, err := Init(config, anima)
	if err != nil {
		return err
	}

	signature, err := ethereum.SignRequest(anima, req)
	if err != nil {
		return err
	}

	header := metadata.New(map[string]string{"signature": signature, "chain": anima.Chain})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	if _, err = client.Issue(ctx, req); err != nil {
		return err
	}
	return nil
}

func Verify(anima *models.Protocol, req *VerifyRequest) (*VerifyResponse, error) {
	config := &Config{Secure: false}
	client, err := Init(config, anima)
	if err != nil {
		return &VerifyResponse{}, err
	}

	signature, err := ethereum.SignRequest(anima, req)
	if err != nil {
		return &VerifyResponse{}, err
	}

	header := metadata.New(map[string]string{"signature": signature, "chain": anima.Chain})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	res, err := client.Verify(ctx, req)
	if err != nil {
		return &VerifyResponse{}, err
	}

	return res, nil
}
