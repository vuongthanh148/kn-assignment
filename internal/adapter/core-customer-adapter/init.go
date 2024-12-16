package corecustomeradapter

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/adaptor"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/port"
)

type adapter struct {
	CoreCustomerClient adaptor.ClientAdaptor
}

func New(ctx context.Context, coreCustomerAdapter adaptor.ClientAdaptor) port.CoreCustomerAdapter {
	return &adapter{
		CoreCustomerClient: coreCustomerAdapter,
	}
}
