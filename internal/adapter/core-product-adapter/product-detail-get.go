package coreproductadapter

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (a *adapter) GetProductsDetail(ctx context.Context, req domain.GetProductsDetailRequest) (*domain.GetProductsDetailResponse, error) {
	ctx, span := tracer.StartNewSpan(ctx, "coreproductadapter.GetProductsDetail")
	defer span.End()

	ctx, err := a.appendGrpcBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	in := (&domainadapter.GetProductsDetailRequest{}).FromDomain(req, "ALL", false)
	out, err := a.serviceClient.ProductDetail(ctx, in)
	if err != nil {
		return nil, err
	}

	return (&domainadapter.GetProductsDetailResponse{}).ToDomain(out), nil
}
