package dto

import (
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	modelv1 "github.com/centraldigital/cfw-sales-x-ordering-api/pkg/model/v1"
)

type Store modelv1.Store

type GetStoresByStaffIdRequest modelv1.GetStoresByStaffIdRequest

func (d GetStoresByStaffIdRequest) ToDomain() domain.GetStoresByStaffIdRequest {
	return domain.GetStoresByStaffIdRequest(d)
}

type GetStoresByStaffIdResponse modelv1.GetStoresByStaffIdResponse

func (GetStoresByStaffIdResponse) FromDomain(d *domain.GetStoresByStaffIdResponse) *GetStoresByStaffIdResponse {
	stores := make([]modelv1.Store, len(d.Stores))
	for i, storeDomain := range d.Stores {
		stores[i] = modelv1.Store(storeDomain)
	}

	return &GetStoresByStaffIdResponse{
		Stores: stores,
	}
}
