package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get products
// @Tags		Product
// @Param		store-code		path	string	true	"store-code"
// @Param		channel-id		path	string	true	"channel-id"
// @Param		 request	body	dto.GetProductsBySkuProductMasterRequest true "request body"
// @Produce		json
// @Success		200	{object}	dto.GetProductsResponse
// @Router 		/v1/products/search-sku/stores/{store-code}/channels/{channel-id} [post]
func (h *handler) GetAllProductsListBySkuProductMaster(ctx context.Context, req dto.GetProductsBySkuProductMasterRequest) (*dto.GetProductsResponse, error) {
	res, err := h.service.GetAllProductsListBySkuProductMaster(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetProductsResponse{}.FromDomain(res), nil
}
