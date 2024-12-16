package infrastructure

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	pbmockv1 "github.com/centraldigital/cfw-sales-x-ordering-api/gen/proto"
	"github.com/centraldigital/cfw-sales-x-ordering-api/property"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewMockGrpcClient(ctx context.Context) (pbmockv1.OpenApiServiceClient, oauth2.TokenSource) {
	caCerPool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("unable to get system cert: %+v", err)
	}

	creds := credentials.NewTLS(&tls.Config{RootCAs: caCerPool})

	mockApiHost := property.Get().ExternalSaleMockApi.Host
	hostWirhPort := fmt.Sprintf("%s:443", mockApiHost)

	grpcConn, err := grpc.Dial(hostWirhPort, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("unable to initial grpc connection: %+v", err)
	}

	httpsHost := fmt.Sprintf("https://%s", mockApiHost)
	sourceToken, err := idtoken.NewTokenSource(ctx, httpsHost)
	if err != nil {
		log.Fatalf("unable to get token source: %+v", err)
	}

	return pbmockv1.NewOpenApiServiceClient(grpcConn), sourceToken
}
