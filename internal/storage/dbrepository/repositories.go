package dbrepository

import (
	"context"

	"github.com/maze1377/manager-vending-machine/internal/storage/entity"
)

type Repository[T any] interface {
	Save(ctx context.Context, models ...T) error
	ByField(context.Context, *QueryByField) ([]T, error)
}

type EventLogRepository interface {
	Repository[*entity.EventLog]
}
