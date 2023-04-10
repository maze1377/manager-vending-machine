package dbrepository

import (
	"context"
)

type Repository[T any] interface {
	Save(ctx context.Context, models ...T) error
	ByField(context.Context, *QueryByField) ([]T, error)
}
