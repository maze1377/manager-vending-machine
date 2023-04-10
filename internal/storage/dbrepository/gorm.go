package dbrepository

import (
	"context"
	"errors"
	"time"

	"github.com/maze1377/manager-vending-machine/internal/metrics"
	"github.com/maze1377/manager-vending-machine/internal/storage"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TableSetting struct {
	ConflictColumn   string
	OnConflictUpdate []string
}

type Gorm[T any] struct {
	Metrics metrics.Communicator
	Client  *gorm.DB
	TableSetting
}

func (g *Gorm[T]) Save(ctx context.Context, models ...T) error {
	start := time.Now()

	if len(models) == 0 {
		return nil
	}

	tr := g.Client.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: g.ConflictColumn}},
		DoUpdates: clause.AssignmentColumns(g.OnConflictUpdate),
	}).WithContext(ctx)

	err := tr.CreateInBatches(models, 100).Error

	g.Metrics.Done(start, "Save", GetStatusByError(err))
	return err
}

func (g *Gorm[T]) ByField(ctx context.Context, config *QueryByField) ([]T, error) {
	start := time.Now()

	var item []T
	qr := g.Client.Where(config.Field+" IN ?", config.Values).WithContext(ctx)

	if config.Limit > 0 {
		qr = qr.Limit(config.Limit)
	}
	if len(config.Joins) > 0 {
		for _, v := range config.Joins {
			qr = qr.Joins(v)
		}
	}
	if len(config.Preloads) > 0 {
		for _, v := range config.Preloads {
			qr = qr.Preload(v)
		}
	}

	err := qr.Find(&item).Error

	g.Metrics.Done(start, "ByField", GetStatusByError(err))
	return item, mapGormToEntityError(err)
}

func mapGormToEntityError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return storage.ErrEntityNotFound
	}
	return err
}

func GetStatusByError(err error) string {
	if err != nil {
		return "error"
	}
	return "ok"
}
