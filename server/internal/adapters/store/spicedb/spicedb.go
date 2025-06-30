package spicedb

import (
	"context"
	"net"

	authzedV1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/server/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type SpiceDB struct {
	Client *authzed.Client
}

func NewClient(cfg Config, clientMetrics *prometheus.ClientMetrics) (*SpiceDB, error) {

	endpoint := net.JoinHostPort(cfg.Host, cfg.Port)
	
	client, err := authzed.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpcutil.WithInsecureBearerToken(cfg.PreSharedKey),
		grpc.WithUnaryInterceptor(clientMetrics.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(clientMetrics.StreamClientInterceptor()),
	)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "spiceDB.NewClient")
	}

	return &SpiceDB{ Client: client }, nil
}


func (s *SpiceDB) Check() error {

	_, err := s.Client.ReadSchema(context.Background(), &authzedV1.ReadSchemaRequest{})

	grpCStatus := status.Convert(err)

	if grpCStatus.Code() == codes.Unavailable {
		return err
	}

	return nil
}

func (s *SpiceDB) Close() error {

	err := s.Client.Close()

	if (err != nil) {
		return utils.WrapErrorf(err, utils.ErrorCodeUnknown, "spiceDB.Close")
	}

	return nil
}

