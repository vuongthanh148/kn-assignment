package port

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

type CoreCustomerAdapter interface {
	GetCustomersProfilesByIds(ctx context.Context, req domain.GetCustomersByStaffIdRequest, customerIds []string) (*domain.GetCustomersByStaffIdResponse, error)
	GetCustomersByMemberId(ctx context.Context, req domain.GetCustomersQuery) (*domain.GetCustomersByCustomerIdResponse, error)
}

type CoreOrderAdapter interface {
	GetOrderHistory(ctx context.Context, req domain.GetOrderHistoryRequest) (*domain.GetOrderHistoryResponse, error)
}

type CoreProductAdapter interface {
	GetProductCategories(ctx context.Context) (*domain.CategoryGetResp, error)
	GetProducts(ctx context.Context, req domain.GetProductsRequest) (*domain.GetProductsResponse, error)
	GetPricesDetail(ctx context.Context, req domain.GetPricesDetailRequest) (*domain.GetPricesDetailResponse, error)
	GetAllProductsListBySkuProductMaster(ctx context.Context, req domain.GetProductsBySkuProductMasterRequest) (*domain.GetProductsResponse, error)
	GetProductsDetail(ctx context.Context, req domain.GetProductsDetailRequest) (*domain.GetProductsDetailResponse, error)
	GetProductByCriteria(ctx context.Context, req domain.GetProductsByCriteriaRequest) (*domain.GetProductsResponse, error)
}

type CoreStoreAdapter interface {
	GetStores(ctx context.Context) (*domain.GetStoresByStaffIdResponse, error)
}
