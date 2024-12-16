package dto

import (
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	modelv1 "github.com/centraldigital/cfw-sales-x-ordering-api/pkg/model/v1"
)

type ProductPromotion modelv1.ProductPromotion

type Product modelv1.Product

func (ProductPromotion) FromDomain(d domain.ProductPromotion) modelv1.ProductPromotion {
	return modelv1.ProductPromotion(d)
}

func (Product) FromDomain(d domain.Product) modelv1.Product {
	productPromotions := make([]modelv1.ProductPromotion, len(d.Promotions))
	for i, productPromotion := range d.Promotions {
		productPromotions[i] = ProductPromotion{}.FromDomain(productPromotion)
	}

	return modelv1.Product{
		Sku:               d.Sku,
		Barcode:           d.Barcode,
		NameEn:            d.NameEn,
		NameTh:            d.NameTh,
		DescriptionEn:     d.DescriptionEn,
		DescriptionTh:     d.DescriptionTh,
		Brand:             d.Brand,
		ImageUrls:         d.ImageUrls,
		UnitEn:            d.UnitEn,
		UnitTh:            d.UnitTh,
		AvgWeight:         d.AvgWeight,
		Stock:             d.Stock,
		QtyLimitPerOrder:  d.QtyLimitPerOrder,
		QtyLimitPerDay:    d.QtyLimitPerDay,
		IsAvailableStock:  d.IsAvailableStock,
		IsWeightItem:      d.IsWeightItem,
		IsWeightScale:     d.IsWeightScale,
		IsVat:             d.IsVat,
		Price:             d.Price,
		PricePr:           d.PricePr,
		PricePerKg:        d.PricePerKg,
		PricePrPerKg:      d.PricePrPerKg,
		Promotions:        productPromotions,
		ClassCode:         d.ClassCode,
		CategoryCode:      d.CategoryCode,
		CategoryNameEn:    d.CategoryNameEn,
		CategoryNameTh:    d.CategoryNameTh,
		SubCategoryCode:   d.SubCategoryCode,
		SubCategoryNameEn: d.SubCategoryNameEn,
		SubCategoryNameTh: d.SubCategoryNameTh,
		Rank:              d.Rank,
	}
}

type GetProductsRequest modelv1.GetProductsRequest

func (d GetProductsRequest) ToDomain() domain.GetProductsRequest {
	return domain.GetProductsRequest(d)
}

type GetProductsBySkuProductMasterRequest modelv1.GetProductsBySkuProductMasterRequest

func (d GetProductsBySkuProductMasterRequest) ToDomain() domain.GetProductsBySkuProductMasterRequest {
	return domain.GetProductsBySkuProductMasterRequest(d)
}

type GetProductsResponse modelv1.GetProductsResponse

func (GetProductsResponse) FromDomain(d *domain.GetProductsResponse) *GetProductsResponse {
	productsPkg := make([]modelv1.Product, len(d.Products))
	for i, product := range d.Products {
		productsPkg[i] = Product{}.FromDomain(product)
	}

	return &GetProductsResponse{
		Products:   productsPkg,
		Pagination: d.Pagination,
	}
}

type GetProductsDetailRequest modelv1.GetProductsDetailRequest

func (d GetProductsDetailRequest) ToDomain() domain.GetProductsDetailRequest {
	return domain.GetProductsDetailRequest(d)
}

type GetProductsDetailResponse modelv1.GetProductsDetailResponse

type ProductDetail modelv1.ProductDetail

func (ProductDetail) FromDomain(d domain.ProductDetail) modelv1.ProductDetail {
	return modelv1.ProductDetail{
		AvgWeight:      d.AvgWeight,
		BrandEn:        d.BrandEn,
		BrandTh:        d.BrandTh,
		Barcode:        d.Barcode,
		CategoryCode:   d.CategoryCode,
		CategoryNameEn: d.CategoryNameEn,
		CategoryNameTh: d.CategoryNameTh,
		ClassCode:      d.ClassCode,
		DescriptionEn:  d.DescriptionEn,
		DescriptionTh:  d.DescriptionTh,
		ImageUrl:       d.ImageUrl,
		// IsAvailableStock: d.IsAvailableStock,
		IsVat:         d.IsVat,
		IsWeightScale: d.IsWeightScale,
		IsWeightItem:  d.IsWeightItem,
		IsActive:      d.IsActive,
		IsOnline:      d.IsOnline,
		NameEn:        d.NameEn,
		NameTh:        d.NameTh,
		Price:         d.Price,
		PricePerKg:    d.PricePerKg,
		// PricePr: d.PricePr,
		// PricePrPerKg: d.PricePrPerKg,
		// ProductPromotion: d.Promotions,
		QtyLimitPerDay:   d.QtyLimitPerDay,
		QtyLimitPerOrder: d.QtyLimitPerDay,
		Rank:             d.Rank,
		Sku:              d.Sku,
		// Stock: d.Stock,
		SubCategoryCode:   d.SubCategoryCode,
		SubCategoryNameEn: d.SubCategoryNameEn,
		SubCategoryNameTh: d.SubCategoryNameTh,
		UnitEn:            d.UnitTh,
		UnitTh:            d.UnitEn,
	}
}

func (GetProductsDetailResponse) FromDomain(d *domain.GetProductsDetailResponse) *GetProductsDetailResponse {
	productsPkg := make([]modelv1.ProductDetail, len(d.ProductsDetail))
	for i, product := range d.ProductsDetail {
		productsPkg[i] = ProductDetail{}.FromDomain(product)
	}

	return &GetProductsDetailResponse{
		Products: productsPkg,
	}
}

type GetProductStockRequest modelv1.GetProductStockRequest

func (d GetProductStockRequest) ToDomain() domain.GetProductStockRequest {
	return domain.GetProductStockRequest(d)
}

type GetProductStockResponse modelv1.GetProductStockResponse

func (GetProductStockResponse) FromDomain(d *domain.GetProductStockResponse) *GetProductStockResponse {
	productsPkg := make([]modelv1.ProductStockData, len(d.Products))
	for i, productStock := range d.Products {
		productsPkg[i] = modelv1.ProductStockData{
			Sku:       productStock.Sku,
			Stock:     productStock.Stock,
			StoreCode: productStock.StoreCode,
			NameTh:    productStock.NameTh,
			NameEn:    productStock.NameEn,
			AddressTh: productStock.AddressTh,
			AddressEn: productStock.AddressEn,
		}
	}
	return &GetProductStockResponse{
		Products: productsPkg,
	}
}

type GetProductsByCriteriaRequest modelv1.GetProductsByCriteriaRequest

func (d GetProductsByCriteriaRequest) ToDomain() domain.GetProductsByCriteriaRequest {
	return domain.GetProductsByCriteriaRequest(d)
}
