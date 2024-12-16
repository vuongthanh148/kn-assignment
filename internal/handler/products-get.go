package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get products
// @Tags		Product
// @Security 	ApiKeyAuth
// @Param		store-code		path	string	true	"store-code"
// @Param		channel-id		path	string	true	"channel-id"
// @Produce		json
// @Success		200	{object}	dto.GetProductsResponse
// @Router 		/v1/products/stores/{store-code}/channels/{channel-id} [post]
func (h *handler) GetProducts(ctx context.Context, req dto.GetProductsRequest) (*dto.GetProductsResponse, error) {
	res, err := h.service.GetProducts(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetProductsResponse{}.FromDomain(res), nil
}
