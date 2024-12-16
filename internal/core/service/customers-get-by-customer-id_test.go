package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomersByMemberId(t *testing.T) {
	t.Run("01_success", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		tm.coreProductAdapter.On("GetProductCategories", ctx).Return(&domain.CategoryGetResp{
			Category: []domain.CategoryData{
				{
					Code:     "01",
					NameTh:   "เนื้อสัตว์",
					NameEn:   "Meat",
					ImageUrl: "https://storage.googleapis.com/cfw-online-platform-dev_storage/product-category-img/01.jpg",
					SubCategoryData: []domain.SubCategoryData{
						{
							Code:     "01",
							NameTh:   "เนื้อหมู",
							NameEn:   "Pork",
							ImageUrl: "",
						},
					},
				},
				{
					Code:     "03",
					NameTh:   "ปลาและอาหารทะเล",
					NameEn:   "Fish & Seafoods",
					ImageUrl: "https://storage.googleapis.com/cfw-online-platform-dev_storage/product-category-img/03.jpg",
					SubCategoryData: []domain.SubCategoryData{
						{
							Code:     "02",
							NameTh:   "ปลา",
							NameEn:   "Fish",
							ImageUrl: "",
						},
					},
				},
			},
		}, nil)
		_, err := tm.service.GetProductCategories(ctx)
		assert.NoError(t, err)
	})
	t.Run("01_fail_error_getting_data", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		tm.coreProductAdapter.On("GetProductCategories", ctx).Return(nil, errors.New("error getting category data from core api"))
		_, err := tm.service.GetProductCategories(ctx)
		assert.Error(t, err)
	})
}
