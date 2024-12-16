package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get customers in store by staff-id
// @Tags		Space
// @Security 	ApiKeyAuth
// @Param		staff-id	path	string	true	"staff-id"
// @Param		store-code	path	string	true	"store-code"
// @Param       q 	 query  dto.GetCustomersByStaffIdRequest true  "query params"
// @Produce		json
// @Success		200	{object}	dto.GetCustomersByStaffIdResponse
// @Router 		/v1/spaces/staffs/{staff-id}/stores/{store-code}/customers [get]
func (h *handler) GetCustomersByStaffId(ctx context.Context, req dto.GetCustomersByStaffIdRequest) (*dto.GetCustomersByStaffIdResponse, error) {
	res, err := h.service.GetCustomersByStaffId(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetCustomersByStaffIdResponse{}.FromDomain(res), nil
}
