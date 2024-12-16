package corecustomeradapter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/centraldigital/cfw-core-lib/pkg/adaptor"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	"github.com/centraldigital/cfw-sales-x-ordering-api/property"
)

func (a *adapter) GetCustomersByMemberId(ctx context.Context, req domain.GetCustomersQuery) (*domain.GetCustomersByCustomerIdResponse, error) {
	_, span := tracer.StartNewSpan(ctx, "corecustomeradapter.GetCustomersByMemberId")
	defer span.End()

	header := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	queryParams := []adaptor.QueryParam{
		{
			Key:   "member-id",
			Value: req.MemberId,
		},
		{
			Key:   "page-size",
			Value: fmt.Sprintf("%d", req.Pagesize),
		},
		{
			Key:   "show-inactive",
			Value: fmt.Sprintf("%t", req.ShowInactive),
		},
		{
			Key:   "order-by",
			Value: req.OrderBy,
		},
	}

	httpResp, err := a.CoreCustomerClient.NewRequest(ctx, "GetCustomersByMemberId").
		WithUrlConfig(property.Get().ExternalCoreCustomerApi.CoreCustomerApiGetCustomersPath, nil).
		Get(&header, &queryParams)
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

	return response.ToDomainCustomer(), nil
}
