package service

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/loggerutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (s *service) GetProducts(ctx context.Context, req domain.GetProductsRequest) (*domain.GetProductsResponse, error) {
	log := loggerutil.GetLogger(ctx, "GetProducts")

	productResp, err := s.coreProductAdapter.GetProducts(ctx, req)
	if err != nil {
		errMsg := "unable to get product from core"
		log.Errorf("%s: %w", errMsg, err)
		return nil, errormodel.NewTechnicalError(errMsg, err)
	}

	skuMap := map[string]bool{}
	for i := range productResp.Products {
		sku := productResp.Products[i].Sku
		if _, found := skuMap[sku]; !found {
			skuMap[sku] = true
		}
	}

	skus := make([]string, 0, len(skuMap))
	for sku := range skuMap {
		skus = append(skus, sku)
	}

	getPricesDetailReq := domain.GetPricesDetailRequest{
		StoreCode: req.StoreCode,
		Skus:      skus,
	}

	priceResp, err := s.coreProductAdapter.GetPricesDetail(ctx, getPricesDetailReq)
	if err != nil {
		errMsg := "unable to get prices detail from core"
		log.Errorf("%s: %w", errMsg, err)
		return nil, errormodel.NewTechnicalError(errMsg, err)
	}

	return mappaingProductAndPriceDetail(ctx, priceResp, productResp)
}
