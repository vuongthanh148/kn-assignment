package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/constant"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get order history
// @Tags		Customers
// @Security 	ApiKeyAuth
// @Param		customer-id		path	string	true	"customer-id"  default(2306000227)
// @Param 		page-id query int false "page-id"
// @Param 		page-size query int false "page-size"
// @Produce		json
// @Success		200	{object}	dto.GetOrderHistoryResponse
// @Router 		/v1/customers/{customer-id}/orders [get]
func (h *handler) GetOrderHistorys(ctx context.Context, req dto.GetOrderHistoryRequest) (*dto.GetOrderHistoryResponse, error) {
	// Verify staff token - space relation ***

	if req.PageID == 0 {
		req.PageID = constant.DEFAULT_PAGE_ID
	}

	if req.PageSize == 0 {
		req.PageSize = constant.DEFAULT_PAGE_SIZE
	}

	res, err := h.service.GetOrderHistory(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetOrderHistoryResponse{}.FromDomain(res), nil

}
