package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
	modelv1 "github.com/centraldigital/cfw-sales-x-ordering-api/pkg/model/v1"
)

// @Summary		update sale template by staff id
// @Tags		Template
// @Security 	ApiKeyAuth
// @Param		staff-id		path	string	true	"staff-id"
// @Param		 request	body	dto.PutSaleTemplatesByStaffIdRequest true "request body"
// @Produce		json
// @Success		200
// @Router 		/v1/sale-templates/staffs/{staff-id} [put]
func (h *handler) PutStaffSaleTemplate(ctx context.Context, req dto.PutSaleTemplatesByStaffIdRequest) (*modelv1.SuccessResponse, error) {
	err := h.service.PutSaleTemplatesByStaffId(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return &modelv1.SuccessResponse_Updated, nil
}
