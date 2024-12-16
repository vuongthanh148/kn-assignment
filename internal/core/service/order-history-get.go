package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetOrderHistory(ctx context.Context, req domain.GetOrderHistoryRequest) (*domain.GetOrderHistoryResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetOrderHistory")
	orderHistory, err := s.coreOrderAdapter.GetOrderHistory(ctx, req)
	if err != nil {
		errmsg := "Error get order history data from core api"
		log.Errorf("%s:%+v", errmsg, err)
		return nil, errormodel.NewTechnicalError(errmsg, err)
	}

	return orderHistory, nil
}
