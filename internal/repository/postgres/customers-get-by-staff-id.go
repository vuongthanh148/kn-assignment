package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
)

func (s *repository) GetCustomersByStaffId(ctx context.Context, staffId string) ([]string, error) {
	ctx, span := tracer.StartNewSpan(ctx, "repository.GetCustomersByStaffId")
	defer span.End()

	queryString := fmt.Sprintf(`SELECT customer_id 
		FROM %s 
		WHERE staff_id = $1 AND deleted_at is null AND deleted_by is null;`, TableName_CustomerAssignment)
	args := []interface{}{staffId}

	rows, err := s.dbPool.Query(ctx, queryString, args...)
	if err != nil {
		return nil, err
	}

	var customers []string
	if err := s.scanApi.ScanAll(&customers, rows); err != nil {
		return nil, err
	}

	return customers, nil
}
