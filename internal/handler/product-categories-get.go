package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get product categories
// @Tags		Category
// @Security 	ApiKeyAuth
// @Produce		json
// @Success		200	{object}	dto.CategoryGetResp
// @Router 		/v1/categories [get]
func (h *handler) GetProductCategories(ctx context.Context, _ dto.EmptyStruct) (*dto.CategoryGetResp, error) {
	res, err := h.service.GetProductCategories(ctx)
	if err != nil {
		return nil, err
	}

	return dto.CategoryGetResp{}.FromDomain(res), nil
}
