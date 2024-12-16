package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (r *repository) GetStoresByStaffId(ctx context.Context, staffId string) ([]domain.Store, error) {
	ctx, span := tracer.StartNewSpan(ctx, "repository.GetSaleTemplatesByStaffId")
	defer span.End()

	queryString := fmt.Sprintf(`SELECT distinct staff_id, store_code as code, 'Store TH - ' || store_code as name_th, 'Store EN - ' || store_code as name_en
		FROM %s
		WHERE staff_id = $1 AND deleted_at is null AND deleted_by is null;
	`, TableName_SpaceSetting)
	args := []any{staffId}

	rows, err := r.dbPool.Query(ctx, queryString, args...)
	if err != nil {
		return nil, err
	}

	var stores []domain.Store
	if err := r.scanApi.ScanAll(&stores, rows); err != nil {
		return nil, err
	}
	return stores, nil
}
