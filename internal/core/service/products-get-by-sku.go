package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetProductsBySku(ctx context.Context, req domain.GetProductsDetailRequest) (*domain.GetProductsDetailResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetProductsBySku")

	if len(req.Skus) == 0 {
		return &domain.GetProductsDetailResponse{}, nil
	}

	productResp, err := s.coreProductAdapter.GetProductsDetail(ctx, req)
	if err != nil {
		log.Errorf("Error getting product from Product Master: %v", err)
		return nil, errormodel.NewTechnicalError("Error getting product from Product Master", err)
	}

	return productResp, nil
}
