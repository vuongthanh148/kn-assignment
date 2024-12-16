package handler

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/handler/dto"
)

// @Summary		get customers in store by member-id
// @Tags		Customers
// @Security 	ApiKeyAuth
// @Param		customer-id	path	string	true	"customer-id"
// @Produce		json
// @Success		200	{object}	dto.GetCustomerResponse
// @Router 		/v1/customers/{customer-id} [get]
func (h *handler) GetCustomersByMemberId(ctx context.Context, req dto.GetCustomerRequest) (*dto.GetCustomerResponse, error) {
	res, err := h.service.GetCustomersByMemberId(ctx, req.ToDomain())
	if err != nil {
		return nil, err
	}

	return dto.GetCustomerResponse{}.FromDomain(res), nil
}
