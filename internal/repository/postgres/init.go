package repository

import (
	"context"
	"errors"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/port"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	dbPool     *pgxpool.Pool
	scanApi    *pgxscan.API
	sqlbuilder sqlbuilder.Flavor
}

func New(dbPool *pgxpool.Pool, scanApi *pgxscan.API, sqlbuilder sqlbuilder.Flavor) port.Repository {

	return &repository{
		dbPool:     dbPool,
		scanApi:    scanApi,
		sqlbuilder: sqlbuilder,
	}
}

type queryKey struct{}

func (r *repository) setPoolTx(ctx context.Context) context.Context {
	if tx, ok := ctx.Value(TxnObjKey{}).(pgx.Tx); ok {
		return context.WithValue(ctx, queryKey{}, tx)
	}
	return context.WithValue(ctx, queryKey{}, r.dbPool)
}

var (
	ErrNoRowsAffected = errors.New("no rows affected")
)
