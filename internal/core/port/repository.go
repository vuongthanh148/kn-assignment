package port

import (
	"context"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

type Repository interface {
	GetCustomersByStaffId(ctx context.Context, staffId string) ([]string, error)
	GetSaleTemplatesByStaffId(ctx context.Context, staffId string) ([]domain.SaleTemplate, error)
	GetStoresByStaffId(ctx context.Context, staffId string) ([]domain.Store, error)
	PostSaleTemplatesByStaffId(ctx context.Context, data domain.SaleTemplate) error
	PutSaleTemplatesByStaffId(ctx context.Context, data domain.SaleTemplate) error
}
