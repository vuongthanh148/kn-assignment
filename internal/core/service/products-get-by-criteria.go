package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetProductsByCriteria(ctx context.Context, req domain.GetProductsByCriteriaRequest) (*domain.GetProductsResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetProductsByCriteria")
	data, err := s.coreProductAdapter.GetProductByCriteria(ctx, req)
	if err != nil {
		errMsg := "error getting product by criteria data from mock api"
		log.Errorf("%s:%+v", errMsg, err)
		return nil, errormodel.NewTechnicalError(errMsg, err)
	}

	return data, nil
}
