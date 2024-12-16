package corestoreadapter

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	pbmockv1 "github.com/centraldigital/cfw-sales-x-ordering-api/gen/proto"
	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (a *adapter) GetStores(ctx context.Context) (*domain.GetStoresByStaffIdResponse, error) {
	ctx, span := tracer.StartNewSpan(ctx, "corestoreadapter.GetStoresByStaffId")
	defer span.End()

	ctx, err := a.appendGrpcBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	in := pbmockv1.StoreMasterRequest{}
	out, err := a.serviceClient.Stores(ctx, &in)
	if err != nil {
		return nil, err
	}

	return (*domainadapter.GetStoreResponse)(out).ToDomain(), nil
}
