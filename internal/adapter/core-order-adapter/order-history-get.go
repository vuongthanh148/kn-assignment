package coreorderadapter

import (
	"context"
	"net/http"
	"strconv"

	"github.com/centraldigital/cfw-core-lib/pkg/adaptor"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	"github.com/centraldigital/cfw-sales-x-ordering-api/property"
)

func (a *adapter) GetOrderHistory(ctx context.Context, req domain.GetOrderHistoryRequest) (*domain.GetOrderHistoryResponse, error) {
	_, span := tracer.StartNewSpan(ctx, "coreorderadapter.GetOrderHistory")
	defer span.End()

	queryParams := []adaptor.QueryParam{
		{
			Key:   "member-id",
			Value: req.CustomerId,
		},
		{
			Key:   "sale-source",
			Value: "WA",
		},
		{
			Key:   "sale-source",
			Value: "EO",
		},
		{
			Key:   "sale-source",
			Value: "OP",
		},
		{
			Key:   "sort-by",
			Value: "order_datetime",
		},
		{
			Key:   "order-by",
			Value: "DESC",
		},
		{
			Key:   "page-id",
			Value: strconv.Itoa(req.PageID),
		},
		{
			Key:   "page-size",
			Value: strconv.Itoa(req.PageSize),
		},
	}

	header := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	httpResp, err := a.httpClient.NewRequest(ctx, "GetOrderHistory").
		WithUrlConfig(property.Get().ExternalCoreOrderApi.CoreOrderApiGetOrderHistoryPath, nil).
		Get(&header, &queryParams)
	if err != nil {
		return nil, err
	}

	response, err := adaptor.HandleReadResponseWithFailError[domainadapter.GetOrderHistoryResponse, error](httpResp).
		HandleSuccessAsJSON(http.StatusOK).
		HandleFailAsJSON(http.StatusConflict,
			func(readResponse adaptor.ReadResponse[domainadapter.GetOrderHistoryResponse]) error {
				return errormodel.ClientErrorDefaultCode(string(readResponse.Data))
			}).
		DoHandle()
	if err != nil {
		return nil, err
	}

	return response.ToDomain(), nil
}
