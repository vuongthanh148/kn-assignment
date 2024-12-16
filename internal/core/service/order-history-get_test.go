package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	domainadapter "github.com/centraldigital/cfw-sales-x-ordering-api/internal/adapter/domain-adapter"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderHistory(t *testing.T) {
	t.Run("Happy flow user 2306000227", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()

		req := domain.GetOrderHistoryRequest{CustomerId: "2306000227", PageID: 1, PageSize: 1}
		var mockOrderHistory domainadapter.GetOrderHistoryResponse
		mockData := `{ "orders": [ { "order_id": "WP23080010000016", "order_status": "NEW", "order_datetime": "2023-08-16 20:13:46", "order_creation_datetime": "2023-08-16 13:13:46", "order_updated_datetime": "2023-08-16 13:13:46", "store_code": "001", "customer_segment": "E00,E00E01,E00E01FAS", "sale_channel": "WA", "net_amount": 1932, "total_discount": 0, "total_marketing_discount": 0, "total_merchandize_discount": 0, "mobile_no": null, "email": null, "tax_destination": null, "items": [ { "line_item_no": 1, "sku_code": "10002260", "pr_code": "10002260", "quantity": 1, "is_weight_item": false, "avg_weight": 0, "quantity_weight_item": 0, "sale_unit": "Case", "line_item_price": 1932, "unit_price": 1932, "promotion_code": "", "reference_line_item_no": "" } ], "delivery_method": "PIC", "delivery_address": { "address_id": "", "firstname": "taste like a fish sauce", "lastname": "", "is_company": false, "company_name": "", "tax_id": "", "branch_id": "", "phone_no": "0875551437", "address_no": "", "building": "", "floor": "", "room": "", "moo": "", "soi": "", "road": "", "subdistrict": "", "district": "", "province": "กรุงเทพมหานคร", "zipcode": "10240", "latitude": "", "longitude": "" }, "delivery_cost": 0, "delivery_datetime": "2023-08-17 07:00:00", "total_amount": 1932, "delivery_slot": "เช้า 01 (08:00 - 10:00)", "promotions": [], "coupons": [], "payments": [ { "payment_id": "WP23080010000016-46845", "payment_type": "POST_PAID", "payment_method": "POD", "payment_amount": 1932, "tendor": "WCOD", "creditcard": "", "approve_code": "", "batch_id": "", "trace_no": "", "payment_datetime": null, "payment_status": "UNPAID", "payment_reason": null, "created_at": "2023-08-16 13:13:46", "created_by": "2306000227", "updated_at": "2023-08-16 13:13:46", "updated_by": "2306000227", "paid_at": null, "additional_info": null } ] } ], "pagination": { "page_id": 1, "page_size": 1, "last_page": 72, "total_elements": 72 } }`
		_ = json.Unmarshal([]byte(mockData), &mockOrderHistory)

		tm.coreOrderAdapter.On("GetOrderHistory", ctx, req).Return(mockOrderHistory.ToDomain(), nil)

		_, err := tm.service.GetOrderHistory(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("Happy flow user 2306000227 page size 10", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()

		req := domain.GetOrderHistoryRequest{CustomerId: "2306000227", PageID: 1, PageSize: 10}
		tm.coreOrderAdapter.On("GetOrderHistory", ctx, req).Return(domainadapter.GetOrderHistoryResponse{}.ToDomain(), nil)

		_, err := tm.service.GetOrderHistory(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("Error case", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()

		req := domain.GetOrderHistoryRequest{CustomerId: "2306000227", PageID: 1, PageSize: 2}
		tm.coreOrderAdapter.On("GetOrderHistory", ctx, req).Return(nil, errors.New("Error"))

		_, err := tm.service.GetOrderHistory(ctx, req)
		assert.Error(t, err)
	})
}
