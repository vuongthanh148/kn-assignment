package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSaleTemplatesByStaffId(t *testing.T) {
	t.Run("01_unable_to_get_sale_template_by_staff_id", func(t *testing.T) {
		ctx := context.Background()
		tm := newTestingModule(t)

		tm.repository.
			On("GetSaleTemplatesByStaffId", ctx, mock.AnythingOfType("string")).
			Return(nil, errors.New("invalid connection string"))

		mockReq := domain.GetSaleTemplatesByStaffIdRequest{StaffId: ""}
		resp, err := tm.service.GetSaleTemplatesByStaffId(ctx, mockReq)

		assert.Nil(t, resp)
		assert.EqualError(t, err, "[9999] unable to get sale-template by staff-id from postgres")
	})

	t.Run("02_success", func(t *testing.T) {
		ctx := context.Background()
		tm := newTestingModule(t)

		mockResp := domain.GetSaleTemplatesByStaffIdResponse{
			SaleTemplates: []domain.SaleTemplate{
				{
					Id:          "",
					StaffId:     "",
					Name:        "",
					Description: "",
					Skus:        []string{},
					IsActive:    true,
				},
			},
		}

		tm.repository.
			On("GetSaleTemplatesByStaffId", ctx, mock.AnythingOfType("string")).
			Return(mockResp.SaleTemplates, nil)

		mockReq := domain.GetSaleTemplatesByStaffIdRequest{StaffId: ""}
		resp, err := tm.service.GetSaleTemplatesByStaffId(ctx, mockReq)

		assert.NoError(t, err)
		assert.Equal(t, *resp, mockResp)
	})
}
