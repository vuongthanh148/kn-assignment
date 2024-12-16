package domainadapter

import (
	pbmockv1 "github.com/centraldigital/cfw-sales-x-ordering-api/gen/proto"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

type Store pbmockv1.Stores

func (da *Store) ToDomain() *domain.Store {
	if da == nil {
		return nil
	}

	return &domain.Store{
		Code:      da.Code,
		Channel:   da.Channel,
		NameTh:    da.NameTh,
		NameEn:    da.NameEn,
		AddressTh: da.AddressTh,
		AddressEn: da.AddressEn,
	}
}

type GetStoreResponse pbmockv1.StoreMasterResponse

func (da *GetStoreResponse) ToDomain() *domain.GetStoresByStaffIdResponse {
	stores := make([]domain.Store, 0, len(da.Stores))
	for i := range da.Stores {
		store := (*Store)(da.Stores[i]).ToDomain()
		if store != nil {
			stores = append(stores, *store)
		}
	}

	return &domain.GetStoresByStaffIdResponse{
		Stores: stores,
	}
}
