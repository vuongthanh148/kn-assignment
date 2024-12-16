package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (r *repository) PutSaleTemplatesByStaffId(ctx context.Context, data domain.SaleTemplate) error {
	ctx, span := tracer.StartNewSpan(ctx, "repository.PutSaleTemplatesByStaffId")
	defer span.End()

	queryString := fmt.Sprintf(`UPDATE %s SET (
		name, description, sku, staff_id, is_active,
		updated_at, updated_by
	) = ($2, $3, $4, $5, $6, $7, $8)
	WHERE id = $1;`, TableName_SaleTemplate)
	args := []interface{}{
		data.Id,
		data.Name,
		data.Description,
		data.Skus,
		data.StaffId,
		data.IsActive,
		data.UpdatedAt,
		data.UpdatedBy,
	}

	_, err := r.dbPool.Exec(ctx, queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
