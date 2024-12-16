package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetCustomersByStaffId(ctx context.Context, req domain.GetCustomersByStaffIdRequest) (*domain.GetCustomersByStaffIdResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetCustomersByStaffId")

	customerLists, err := s.repository.GetCustomersByStaffId(ctx, req.StaffId)
	if err != nil {
		log.Errorf("Error getting customers by staff id: %v", err)
		return nil, errormodel.NewTechnicalError("Error getting customers by staff id", err)
	}

	if len(customerLists) == 0 {
		return &domain.GetCustomersByStaffIdResponse{}, nil
	}

	customerDetailsByIds, err := s.coreCustomerAdapter.GetCustomersProfilesByIds(ctx, req, customerLists)
	if err != nil {
		log.Errorf("Error getting customer profiles by ids: %v", err)
		return nil, errormodel.NewTechnicalError("Error getting customer profiles by ids", err)
	}

	return customerDetailsByIds, nil
}
