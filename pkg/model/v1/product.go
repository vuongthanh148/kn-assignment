package modelv1

import (
	"github.com/centraldigital/cfw-core-lib/pkg/model/basemodel"
	"github.com/centraldigital/cfw-core-lib/pkg/model/jsonmodel"
	"github.com/shopspring/decimal"
)

type ProductPromotion struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type Product struct {
	Sku               string             `json:"sku"`
	Barcode           string             `json:"barcode"`
	NameEn            string             `json:"name_en"`
	NameTh            string             `json:"name_th"`
	DescriptionEn     string             `json:"description_en"`
	DescriptionTh     string             `json:"description_th"`
	Brand             string             `json:"brand"`
	ImageUrls         []string           `json:"image_url"`
	UnitEn            string             `json:"unit_en"`
	UnitTh            string             `json:"unit_th"`
	AvgWeight         decimal.Decimal    `json:"avg_weight"`
	Stock             decimal.Decimal    `json:"stock"`
	QtyLimitPerOrder  int                `json:"qty_limit_per_order"`
	QtyLimitPerDay    int                `json:"qty_limit_per_day"`
	IsAvailableStock  bool               `json:"is_available_stock"`
	IsWeightItem      bool               `json:"is_weight_item"`
	IsWeightScale     bool               `json:"is_weight_scale"`
	IsVat             bool               `json:"is_vat"`
	Price             jsonmodel.Money    `json:"price"`
	PricePr           jsonmodel.Money    `json:"price_pr"`
	PricePerKg        *jsonmodel.Money   `json:"price_per_kg"`
	PricePrPerKg      *jsonmodel.Money   `json:"price_pr_per_kg"`
	Promotions        []ProductPromotion `json:"promotion"`
	ClassCode         string             `json:"class_code"`
	CategoryCode      string             `json:"category_code"`
	CategoryNameEn    string             `json:"category_name_en"`
	CategoryNameTh    string             `json:"category_name_th"`
	SubCategoryCode   string             `json:"sub_category_code"`
	SubCategoryNameEn string             `json:"sub_category_name_en"`
	SubCategoryNameTh string             `json:"sub_category_name_th"`
	Rank              int                `json:"rank"`
}

type GetProductsRequest struct {
	StoreCode string `json:"-" uri:"store-code"`
	ChannelId string `json:"-" uri:"channel-id"`

	SearchType      string `json:"search_type"`
	Keyword         string `json:"keyword"`
	CategoryCode    string `json:"category_code"`
	SubCategoryCode string `json:"sub_category_code"`
	SortAttribute   string `json:"sort_attr" default:"default"`

	Pagination basemodel.PaginationOffsetRequest `json:"pagination"`
}
type GetProductsBySkuProductMasterRequest struct {
	StoreCode string   `json:"-" uri:"store-code"`
	ChannelId string   `json:"-" uri:"channel-id"`
	Sku       []string `json:"sku"`
}

type GetProductsByCriteriaRequest struct {
	SearchType      string `json:"search_type"`
	Keyword         string `json:"keyword"`
	CategoryCode    string `json:"category_code"`
	SubCategoryCode string `json:"sub_category_code"`
	SortAttribute   string `json:"sort_attr" default:"default"`

	Pagination *basemodel.PaginationOffsetRequest `json:"pagination"`
}

type GetProductsResponse struct {
	Products   []Product                          `json:"products"`
	Pagination basemodel.PaginationOffsetResponse `json:"pagination"`
}

type GetProductsDetailRequest struct {
	Skus []string `json:"sku"`
}

type PromotionDetail struct {
	ImageUrlEn      string `json:"image_url_en"`
	ImageUrlTh      string `json:"image_url_th"`
	PrefixPromotion string `json:"prefix_promotion"`
	PromotionNameEn string `json:"promotion_name_en"`
	PromotionNameTh string `json:"promotion_name_th"`
}

type ProductDetail struct {
	AvgWeight         decimal.Decimal   `json:"avg_weight"`
	BrandEn           string            `json:"brand_en"`
	BrandTh           string            `json:"brand_th"`
	Barcode           string            `json:"barcode"`
	CategoryCode      string            `json:"category_code"`
	CategoryNameEn    string            `json:"category_name_en"`
	CategoryNameTh    string            `json:"category_name_th"`
	ClassCode         string            `json:"class_code"`
	DescriptionEn     string            `json:"description_en"`
	DescriptionTh     string            `json:"description_th"`
	ImageUrl          []string          `json:"image_url"`
	IsAvailableStock  bool              `json:"is_available_stock"`
	IsVat             bool              `json:"is_vat"`
	IsWeightScale     bool              `json:"is_weight_scale"`
	IsWeightItem      bool              `json:"is_weight_item"`
	IsActive          bool              `json:"is_active"`
	IsOnline          bool              `json:"is_online"`
	NameEn            string            `json:"name_en"`
	NameTh            string            `json:"name_th"`
	Price             jsonmodel.Money   `json:"price"`
	PricePerKg        *jsonmodel.Money  `json:"price_per_kg"`
	PricePr           jsonmodel.Money   `json:"price_pr"`
	PricePrPerKg      *jsonmodel.Money  `json:"price_pr_per_kg"`
	ProductPromotion  []PromotionDetail `json:"promotion"`
	QtyLimitPerDay    int               `json:"qty_limit_per_day"`
	QtyLimitPerOrder  int               `json:"qty_limit_per_order"`
	Rank              int               `json:"rank"`
	Sku               string            `json:"sku"`
	Stock             decimal.Decimal   `json:"stock"`
	SubCategoryCode   string            `json:"sub_category_code"`
	SubCategoryNameEn string            `json:"sub_category_name_en"`
	SubCategoryNameTh string            `json:"sub_category_name_th"`
	UnitEn            string            `json:"unit_en"`
	UnitTh            string            `json:"unit_th"`
}

type GetProductsDetailResponse struct {
	Products []ProductDetail `json:"products"`
}

type GetProductStockRequest struct {
	StoreCode []string `json:"store_code"`
	Sku       string   `json:"-" uri:"sku"`
}

type ProductStockData struct {
	Sku       string          `json:"sku"`
	Stock     decimal.Decimal `json:"stock"`
	StoreCode string          `json:"store_code"`
	NameTh    string          `json:"name_th"`
	NameEn    string          `json:"name_en"`
	AddressTh string          `json:"address_th"`
	AddressEn string          `json:"address_en"`
}

type GetProductStockResponse struct {
	Products []ProductStockData `json:"products"`
}
