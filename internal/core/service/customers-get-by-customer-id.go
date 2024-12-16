package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetCustomersByMemberId(ctx context.Context, req domain.GetCustomerRequest) (*domain.GetCustomerResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetCustomers")

	query := domain.GetCustomersQuery{
		MemberId:     req.MemberId,
		Pagesize:     1,
		ShowInactive: true,

		OrderBy: "ASC",
	}

	customerDetailsByIds, err := s.coreCustomerAdapter.GetCustomersByMemberId(ctx, query)
	if err != nil || len(customerDetailsByIds.Customers) != 1 {
		log.Errorf("Error getting customer profiles by ids: %v", err)
		return nil, errormodel.NewTechnicalError("Error getting customer profiles by ids", err)
	}
	customerInfo := customerDetailsByIds.Customers[0]

	return &customerInfo, nil
}
