package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetSaleTemplatesByStaffId(ctx context.Context, req domain.GetSaleTemplatesByStaffIdRequest) (*domain.GetSaleTemplatesByStaffIdResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetSaleTemplatesByStaffId")

	saleTemplates, err := s.repository.GetSaleTemplatesByStaffId(ctx, req.StaffId)
	if err != nil {
		errMsg := "unable to get sale-template by staff-id from postgres"
		log.Errorf("%s: %w", errMsg, err)
		return nil, errormodel.NewTechnicalError(errMsg, err)
	}

	return &domain.GetSaleTemplatesByStaffIdResponse{
		SaleTemplates: saleTemplates,
	}, nil
}
