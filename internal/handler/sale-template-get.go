package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get sale template by staff id
// @Tags		Template
// @Security 	ApiKeyAuth
// @Param		staff-id		path	string	true	"staff-id"
// @Produce		json
// @Success		200	{object}	dto.GetSaleTemplatesByStaffIdResponse
// @Router 		/v1/sale-templates/staffs/{staff-id} [get]
func (h *handler) GetStaffSaleTemplate(ctx context.Context, req dto.GetSaleTemplatesByStaffIdRequest) (*dto.GetSaleTemplatesByStaffIdResponse, error) {
	res, err := h.service.GetSaleTemplatesByStaffId(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetSaleTemplatesByStaffIdResponse{}.FromDomain(res), nil
}
