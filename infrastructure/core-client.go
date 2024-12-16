package infrastructure

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/adaptor"
	infrastructureutil "github.com/centraldigital/cfw-core-lib/pkg/util/infrastructureutil/http-client"
	"github.com/centraldigital/cfw-sales-x-ordering-api/property"
)

func NewCoreCustomerHttpClient(ctx context.Context) adaptor.ClientAdaptor {
	return infrastructureutil.NewIdTokenAdapterMiddlewareNoLoging(ctx, property.Get().ExternalCoreCustomerApi.CoreCustomerApiHost, "core-customer-api-client")
}

func NewCoreOrderHttpClient(ctx context.Context) adaptor.ClientAdaptor {
	return infrastructureutil.NewIdTokenAdapterMiddlewareNoLoging(ctx, property.Get().ExternalCoreOrderApi.CoreOrderApiHost, "core-order-api-client")
}
