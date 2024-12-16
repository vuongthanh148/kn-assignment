package coreorderadapter

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/adaptor"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/port"
)

type adapter struct {
	httpClient adaptor.ClientAdaptor
}

func New(ctx context.Context, clientAdapter adaptor.ClientAdaptor) port.CoreOrderAdapter {
	return &adapter{
		httpClient: clientAdapter,
	}
}
