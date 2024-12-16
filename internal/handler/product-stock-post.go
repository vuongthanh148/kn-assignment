package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get products stock
// @Tags		Product
// @Security 	ApiKeyAuth
// @Param		sku		path	string	true	"sku"
// @Param		request	body	dto.GetProductStockRequest true "request body"
// @Produce		json
// @Success		200	{object}	dto.GetProductStockResponse
// @Router 		/v1/products/sku/{sku}/available-stock [post]
func (h *handler) GetProductStock(ctx context.Context, req dto.GetProductStockRequest) (*dto.GetProductStockResponse, error) {
	res, err := h.service.GetProductStock(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetProductStockResponse{}.FromDomain(res), nil
}
