package corestoreadapter

import (
	"context"
	"fmt"

	pbmockv1 "github.com/centraldigital/cfw-sales-x-ordering-api/gen/proto"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/port"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/metadata"
)

type adapter struct {
	serviceClient pbmockv1.OpenApiServiceClient
	sourceToken   oauth2.TokenSource
}

func New(ctx context.Context, grpcClient pbmockv1.OpenApiServiceClient, sourceToken oauth2.TokenSource) port.CoreStoreAdapter {
	return &adapter{
		serviceClient: grpcClient,
		sourceToken:   sourceToken,
	}
}

func (a *adapter) appendGrpcBearerToken(ctx context.Context) (context.Context, error) {
	token, err := a.sourceToken.Token()
	if err != nil {
		return nil, err
	}

	tokenString := fmt.Sprintf("Bearer %s", token.AccessToken)
	return metadata.AppendToOutgoingContext(ctx, "authorization", tokenString), nil
}
