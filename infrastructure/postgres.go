package infrastructure

import (
	"context"

	"github.com/centraldigital/cfw-core-lib/pkg/util/infrastructureutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/property"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, *pgxscan.API) {
	pgx, scanapi, err := infrastructureutil.NewPostgresWithScanApi(ctx, property.Get().PostgresConfig)
	if err != nil {
		panic(err)
	}

	return pgx, scanapi
}
