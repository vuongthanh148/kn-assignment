package coreproductadapter

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	pbmockv1 "github.com/centraldigital/cfw-sales-x-ordering-api/gen/proto"
	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (a *adapter) GetAllProductsListBySkuProductMaster(ctx context.Context, req domain.GetProductsBySkuProductMasterRequest) (*domain.GetProductsResponse, error) {
	ctx, span := tracer.StartNewSpan(ctx, "coreproductadapter.GetProducts")
	defer span.End()

	ctx, err := a.appendGrpcBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	// TODO enhance grpc and map values in "req" to "in"
	in := pbmockv1.ProductDetailRequest{}
	out, err := a.serviceClient.ProductDetail(ctx, &in)
	if err != nil {
		return nil, err
	}

	return (&domainadapter.GetProductsResponse{}).ProductDetailToDomain(out, req.ChannelId), nil
}
