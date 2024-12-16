package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/centraldigital/cfw-core-lib/pkg/model/jsonmodel"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsBySku(t *testing.T) {
	t.Run("01_success", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		tm.coreProductAdapter.On("GetProductsDetail", ctx, domain.GetProductsDetailRequest{
			Skus: []string{"2312000105", "2312000106"},
		}).Return(&domain.GetProductsDetailResponse{
			ProductsDetail: []domain.ProductDetail{
				{
					Sku:               "2312000105",
					AvgWeight:         decimal.NewFromFloat(1.23),
					BrandEn:           "Brand A",
					BrandTh:           "แบรนด์ A",
					Barcode:           "1234567890123",
					CategoryCode:      "CAT01",
					CategoryNameEn:    "Category 1",
					CategoryNameTh:    "หมวดหมู่ 1",
					ClassCode:         "CLASS01",
					DescriptionEn:     "This is product 1",
					DescriptionTh:     "นี่คือผลิตภัณฑ์ 1",
					ImageUrl:          []string{"http://example.com/image1.jpg"},
					IsAvailableStock:  true,
					IsVat:             true,
					IsWeightScale:     false,
					IsWeightItem:      false,
					IsActive:          true,
					IsOnline:          true,
					NameEn:            "Product 1",
					NameTh:            "ผลิตภัณฑ์ 1",
					Price:             jsonmodel.Money(decimal.NewFromFloat(99.99)),
					PricePerKg:        nil,
					PricePr:           jsonmodel.Money(decimal.NewFromFloat(109.99)),
					PricePrPerKg:      nil,
					ProductPromotion:  []domain.PromotionDetail{},
					QtyLimitPerDay:    10,
					QtyLimitPerOrder:  5,
					Rank:              1,
					Stock:             decimal.NewFromFloat(100),
					SubCategoryCode:   "SUBCAT01",
					SubCategoryNameEn: "Subcategory 1",
					SubCategoryNameTh: "หมวดหมู่ย่อย 1",
					UnitEn:            "Piece",
					UnitTh:            "ชิ้น",
				},
				{
					Sku:               "2312000106",
					AvgWeight:         decimal.NewFromFloat(2.34),
					BrandEn:           "Brand B",
					BrandTh:           "แบรนด์ B",
					Barcode:           "1234567890124",
					CategoryCode:      "CAT02",
					CategoryNameEn:    "Category 2",
					CategoryNameTh:    "หมวดหมู่ 2",
					ClassCode:         "CLASS02",
					DescriptionEn:     "This is product 2",
					DescriptionTh:     "นี่คือผลิตภัณฑ์ 2",
					ImageUrl:          []string{"http://example.com/image2.jpg"},
					IsAvailableStock:  true,
					IsVat:             true,
					IsWeightScale:     false,
					IsWeightItem:      false,
					IsActive:          true,
					IsOnline:          true,
					NameEn:            "Product 2",
					NameTh:            "ผลิตภัณฑ์ 2",
					Price:             jsonmodel.Money(decimal.NewFromFloat(199.99)),
					PricePerKg:        nil,
					PricePr:           jsonmodel.Money(decimal.NewFromFloat(219.99)),
					PricePrPerKg:      nil,
					ProductPromotion:  []domain.PromotionDetail{},
					QtyLimitPerDay:    20,
					QtyLimitPerOrder:  10,
					Rank:              2,
					Stock:             decimal.NewFromFloat(200),
					SubCategoryCode:   "SUBCAT02",
					SubCategoryNameEn: "Subcategory 2",
					SubCategoryNameTh: "หมวดหมู่ย่อย 2",
					UnitEn:            "Case",
					UnitTh:            "ลัง",
				},
			},
		}, nil)
		_, err := tm.service.GetProductsBySku(ctx, domain.GetProductsDetailRequest{
			Skus: []string{"2312000105", "2312000106"},
		})
		assert.NoError(t, err)
	})

	t.Run("02_success_with_empty_sku_list", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		_, err := tm.service.GetProductsBySku(ctx, domain.GetProductsDetailRequest{
			Skus: []string{},
		})
		assert.NoError(t, err)
	})

	t.Run("03_fail_error_getting_product_from_product_master", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		tm.coreProductAdapter.On("GetProductsDetail", ctx, domain.GetProductsDetailRequest{
			Skus: []string{"2312000105", "2312000106"},
		}).Return(nil, errors.New("Error getting product from Product Master"))
		_, err := tm.service.GetProductsBySku(ctx, domain.GetProductsDetailRequest{
			Skus: []string{"2312000105", "2312000106"},
		})
		assert.Error(t, err)
	})

	// t.Run("04_fail_error_getting_customer_profiles_by_ids", func(t *testing.T) {
	// 	tm := newTestingModule(t)
	// 	ctx := context.Background()
	// 	tm.repository.On("GetCustomersByStaffId", ctx, "47cd52d1-3e09-4143-9371-0178029b9037").Return([]string{"2312000105", "2312000106"}, nil)
	// 	tm.coreCustomerAdapter.On("GetCustomersProfilesByIds", ctx, domain.GetCustomersByStaffIdRequest{
	// 		StaffId:   "47cd52d1-3e09-4143-9371-0178029b9037",
	// 		StoreCode: "001",
	// 	}, []string{"2312000105", "2312000106"}).Return(nil, errors.New("Error getting customer profiles by ids"))
	// 	_, err := tm.service.GetCustomersByStaffId(ctx, domain.GetCustomersByStaffIdRequest{
	// 		StaffId:   "47cd52d1-3e09-4143-9371-0178029b9037",
	// 		StoreCode: "001"})
	// 	assert.Error(t, err)
	// })

}
