package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get stores by staff-id
// @Tags		Space
// @Security 	ApiKeyAuth
// @Param		staff-id		path	string	true	"staff-id"
// @Produce		json
// @Success		200	{object}	dto.GetStoresByStaffIdResponse
// @Router 		/v1/spaces/staffs/{staff-id}/stores [get]
func (h *handler) GetStoresByStaffId(ctx context.Context, req dto.GetStoresByStaffIdRequest) (*dto.GetStoresByStaffIdResponse, error) {
	res, err := h.service.GetStoresByStaffId(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetStoresByStaffIdResponse{}.FromDomain(res), nil
}
