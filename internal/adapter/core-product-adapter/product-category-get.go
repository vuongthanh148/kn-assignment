package coreproductadapter

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	pbmockv1 "github.com/centraldigital/cfw-sales-x-ordering-api/gen/proto"
	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (a *adapter) GetProductCategories(ctx context.Context) (*domain.CategoryGetResp, error) {
	ctx, span := tracer.StartNewSpan(ctx, "coreproductadapter.GetProductCategories")
	defer span.End()

	ctx, err := a.appendGrpcBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	in := pbmockv1.CategoryRequest{}
	out, err := a.serviceClient.Categories(ctx, &in)
	if err != nil {
		return nil, err
	}

	return (*domainadapter.CategoryGetResp)(out).ToDomain(), nil
}
