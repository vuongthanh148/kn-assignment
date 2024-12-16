package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetProductStock(ctx context.Context, req domain.GetProductStockRequest) (*domain.GetProductStockResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetProductStock")

	storeData, err := s.coreStoreAdapter.GetStores(ctx)
	if err != nil {
		errMsg := "unable to get stores detail from core"
		log.Errorf("%s: %w", errMsg, err)
		return nil, errormodel.NewTechnicalError(errMsg, err)
	}

	storeMap := make(map[string]domain.Store)
	for _, store := range storeData.Stores {
		if _, ok := storeMap[store.Code]; !ok {
			storeMap[store.Code] = store
		}
	}

	stockResp := make([]domain.ProductStockData, 0, len(req.StoreCode))
	for _, storeCode := range req.StoreCode {
		priceResp, err := s.coreProductAdapter.GetPricesDetail(ctx, domain.GetPricesDetailRequest{Skus: []string{req.Sku}, StoreCode: storeCode})
		if err != nil {
			errMsg := "unable to get prices detail from core"
			log.Errorf("%s: %w", errMsg, err)
			return nil, errormodel.NewTechnicalError(errMsg, err)
		}
		if len(priceResp.PricesDetail) != 0 {
			price := priceResp.PricesDetail[0]
			if store, ok := storeMap[price.StoreCode]; ok {
				stockResp = append(stockResp, domain.ProductStockData{
					Sku:       price.Sku,
					Stock:     price.Stock,
					StoreCode: price.StoreCode,
					NameTh:    store.NameTh,
					NameEn:    store.NameEn,
					AddressTh: store.AddressTh,
					AddressEn: store.AddressEn,
				})
			}
		}
	}
	return &domain.GetProductStockResponse{Products: stockResp}, nil
}
