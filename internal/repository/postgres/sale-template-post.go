package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (r *repository) PostSaleTemplatesByStaffId(ctx context.Context, data domain.SaleTemplate) error {
	ctx, span := tracer.StartNewSpan(ctx, "repository.PostSaleTemplatesByStaffId")
	defer span.End()

	queryString := fmt.Sprintf(`INSERT INTO %s (
		id, name, description, sku,
		staff_id, is_active, created_at, created_by,
		updated_at, updated_by
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`, TableName_SaleTemplate)
	args := []interface{}{
		data.Id,
		data.Name,
		data.Description,
		data.Skus,
		data.StaffId,
		data.IsActive,
		data.CreatedAt,
		data.CreatedBy,
		data.UpdatedAt,
		data.UpdatedBy,
	}
	_, err := r.dbPool.Exec(ctx, queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
