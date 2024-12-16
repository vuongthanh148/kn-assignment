package repository

import (
	"context"
	"fmt"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/logger"
	utilTypeConvertor "github.com/centraldigital/cfw-core-lib/pkg/util/typeconvertutil"
	"github.com/jackc/pgx/v5"
)

type TxnObjKey struct{}

func (s *repository) Transactional(ctx context.Context, f func(c context.Context) error) error {
	ctx = utilTypeConvertor.CheckandConvertGinContext(ctx)
	log := logger.ExtractLogger(ctx)
	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error on starting transaction: %+v", err)
	}

	ctx = context.WithValue(ctx, TxnObjKey{}, tx)
	defer func(ctx context.Context) {
		if err := tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			log.Errorf("error on rolling back transaction: %+v", err)
		}
	}(ctx)

	if err := f(ctx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error on committing transaction: %+v", err)
	}

	return nil
}
