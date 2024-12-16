package corecustomeradapter

import (
	"context"
	"net/http"

	"github.com/centraldigital/cfw-core-customer-api/pkg/enum"
	"github.com/centraldigital/cfw-core-lib/pkg/adaptor"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	"github.com/centraldigital/cfw-sales-x-ordering-api/property"
)

func (a *adapter) GetCustomersProfilesByIds(ctx context.Context, req domain.GetCustomersByStaffIdRequest, customerIds []string) (*domain.GetCustomersByStaffIdResponse, error) {
	_, span := tracer.StartNewSpan(ctx, "corecustomeradapter.GetCustomersProfilesByIds")
	defer span.End()

	header := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	request := domainadapter.GetCustomerByMemberIdsRequest{}.FromDomain(req, customerIds)
	request.OrderBy = enum.OrderBy_ASC
	request.ShowInactive = true
	request.PageSize = len(customerIds)
	httpResp, err := a.CoreCustomerClient.NewRequest(ctx, "GetCustomersProfilesByIds").
		WithUrlConfig(property.Get().ExternalCoreCustomerApi.CoreCustomerApiGetCustomersPath, nil).
		Post(request, &header, nil)
	if err != nil {
		return nil, err
	}

	response, err := adaptor.HandleReadResponseWithFailError[domainadapter.GetCustomersResponse, error](httpResp).
		HandleSuccessAsJSON(http.StatusOK).
		HandleFailAsJSON(http.StatusBadRequest,
			func(readResponse adaptor.ReadResponse[domainadapter.GetCustomersResponse]) error {
				return errormodel.ClientErrorDefaultCode(string(readResponse.Data))
			}).
		DoHandle()
	if err != nil {
		return nil, err
	}

	return response.ToDomain(), nil
}
