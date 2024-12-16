package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get products by sku from product master
// @Tags		Product
// @Security 	ApiKeyAuth
// @Produce		json
// @Param		request	body dto.GetProductsDetailRequest true "request body"
// @Success		200	{object}	dto.GetProductsResponse
// @Router 		/v1/products/search-sku [post]
func (h *handler) GetProductsBySku(ctx context.Context, req dto.GetProductsDetailRequest) (*dto.GetProductsDetailResponse, error) {
	res, err := h.service.GetProductsBySku(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetProductsDetailResponse{}.FromDomain(res), nil
}
