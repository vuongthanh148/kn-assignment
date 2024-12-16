package coreproductadapter

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (a *adapter) GetPricesDetail(ctx context.Context, req domain.GetPricesDetailRequest) (*domain.GetPricesDetailResponse, error) {
	ctx, span := tracer.StartNewSpan(ctx, "coreproductadapter.GetPricesDetail")
	defer span.End()

	ctx, err := a.appendGrpcBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	in := (&domainadapter.GetPricesDetailRequest{}).FromDomain(req)
	out, err := a.serviceClient.PriceDetail(ctx, in)
	if err != nil {
		return nil, err
	}

	return (&domainadapter.GetPricesDetailResponse{}).ToDomain(out), nil
}
