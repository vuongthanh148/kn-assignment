package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/constant"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get product by criteria
// @Tags		Product
// @Security 	ApiKeyAuth
// @Param		request	body	dto.GetProductsByCriteriaRequest true "request body"
// @Produce		json
// @Success		200	{object}	dto.GetProductsResponse
// @Router 		/v1/products/search [post]
func (h *handler) GetProductsByCriteria(ctx context.Context, req dto.GetProductsByCriteriaRequest) (*dto.GetProductsResponse, error) {
	if req.Pagination != nil {
		if req.Pagination.PageId == 0 {
			req.Pagination.PageId = constant.DEFAULT_PAGE_ID
		}
		if req.Pagination.PageSize == 0 {
			req.Pagination.PageSize = constant.DEFAULT_PAGE_SIZE
		}
	}

	res, err := h.service.GetProductsByCriteria(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetProductsResponse{}.FromDomain(res), nil
}
