package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetProductCategories(ctx context.Context) (*domain.CategoryGetResp, error) {
	log := loggerutil.GetLogger(ctx, "GetProductCategories")
	catData, err := s.coreProductAdapter.GetProductCategories(ctx)
	if err != nil {
		errMsg := "error getting category data from mock api"
		log.Errorf("%s:%+v", errMsg, err)
		return nil, errormodel.NewTechnicalError(errMsg, err)
	}

	return catData, nil
}
