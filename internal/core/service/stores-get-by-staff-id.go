package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetStoresByStaffId(ctx context.Context, req domain.GetStoresByStaffIdRequest) (*domain.GetStoresByStaffIdResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetStoresByStaffId")

	staffStores, err := s.repository.GetStoresByStaffId(ctx, req.StaffId)
	if err != nil {
		errMsg := "unable to get stores by staff-id from postgres"
		log.Errorf("%s: %w", errMsg, err)
		return nil, errormodel.NewTechnicalError(errMsg, err)
	}

	staffStoreMap := map[string]bool{}
	for i := range staffStores {
		staffStoreMap[staffStores[i].Code] = true
	}

	allStoreResp, err := s.coreStoreAdapter.GetStores(ctx)
	if err != nil {
		errMsg := "unable to get stores from core"
		log.Errorf("%s: %w", errMsg, err)
		return nil, errormodel.NewTechnicalError(errMsg, err)
	}

	filteredStore := make([]domain.Store, 0, len(allStoreResp.Stores))
	for i := range allStoreResp.Stores {
		store := allStoreResp.Stores[i]
		if _, found := staffStoreMap[store.Code]; found {
			filteredStore = append(filteredStore, store)
		}
	}

	return &domain.GetStoresByStaffIdResponse{
		Stores: filteredStore,
	}, nil
}
