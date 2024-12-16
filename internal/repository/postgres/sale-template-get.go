package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

func (r *repository) GetSaleTemplatesByStaffId(ctx context.Context, staffId string) ([]domain.SaleTemplate, error) {
	ctx, span := tracer.StartNewSpan(ctx, "repository.GetSaleTemplatesByStaffId")
	defer span.End()

	queryString := fmt.Sprintf("SELECT * FROM %s WHERE staff_id = $1;", TableName_SaleTemplate)
	args := []any{staffId}

	rows, err := r.dbPool.Query(ctx, queryString, args...)
	if err != nil {
		return nil, err
	}

	var saleTemplates []domain.SaleTemplate
	if err := r.scanApi.ScanAll(&saleTemplates, rows); err != nil {
		return nil, err
	}
	return saleTemplates, nil
}
