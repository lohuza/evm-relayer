package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lohuza/relayer/pkg"
	"github.com/uptrace/bun"
)

type SelectQuery func(*bun.SelectQuery) *bun.SelectQuery
type UpdateQuery func(*bun.UpdateQuery) *bun.UpdateQuery
type DeleteQuery func(*bun.DeleteQuery) *bun.DeleteQuery

type RepositoryBase[T any] interface {
	FindAll(context.Context, ...SelectQuery) ([]T, error)
	FindOne(context.Context, ...SelectQuery) (T, error)
	Save(context.Context, *T) error
	SaveMany(context.Context, *[]T) error
	Update(context.Context, *T) (int64, error)
	UpdateWithQuery(ctx context.Context, model *T, s ...UpdateQuery) (int64, error)
	UpdateMany(context.Context, *[]T, ...string) (int64, error)
	Delete(context.Context, *T) error
	DeleteMany(ctx context.Context, model *[]T) error
	DeleteWithQuery(context.Context, ...DeleteQuery) error
}

type repositoryBase[T any] struct {
	DB bun.IDB
}

func newBaseRepository[T any](db bun.IDB) RepositoryBase[T] {
	return repositoryBase[T]{
		DB: db,
	}
}

func (c repositoryBase[T]) FindAll(ctx context.Context, s ...SelectQuery) ([]T, error) {
	var rows []T

	q := c.DB.NewSelect().Model(&rows)
	for i := range s {
		q.Apply(s[i])
	}

	err := q.Scan(ctx)
	return rows, err
}

func (c repositoryBase[T]) FindOne(ctx context.Context, s ...SelectQuery) (T, error) {
	var row T

	q := c.DB.NewSelect().Model(&row)
	for i := range s {
		q.Apply(s[i])
	}

	err := q.Limit(1).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return row, pkg.ErrRecordNotFound
	}
	return row, err
}

func (c repositoryBase[T]) Save(ctx context.Context, model *T) error {
	_, err := c.DB.NewInsert().Model(model).Exec(ctx)
	return err
}

func (c repositoryBase[T]) SaveMany(ctx context.Context, model *[]T) error {
	_, err := c.DB.NewInsert().Model(model).Exec(ctx)
	return err
}

func (c repositoryBase[T]) Update(ctx context.Context, model *T) (int64, error) {
	res, err := c.DB.NewUpdate().Model(model).WherePK().Exec(ctx)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (c repositoryBase[T]) UpdateWithQuery(ctx context.Context, model *T, s ...UpdateQuery) (int64, error) {
	q := c.DB.NewUpdate().Model(model)
	for i := range s {
		q.Apply(s[i])
	}

	res, err := q.Exec(ctx)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (c repositoryBase[T]) UpdateMany(ctx context.Context, model *[]T, columns ...string) (int64, error) {
	res, err := c.DB.NewUpdate().Model(model).Bulk().Column(columns...).Exec(ctx)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (c repositoryBase[T]) Delete(ctx context.Context, model *T) error {
	_, err := c.DB.NewDelete().Model(model).WherePK().Exec(ctx)
	return err
}

func (c repositoryBase[T]) DeleteMany(ctx context.Context, model *[]T) error {
	_, err := c.DB.NewDelete().Model(model).WherePK().Exec(ctx)
	return err
}

func (c repositoryBase[T]) DeleteWithQuery(ctx context.Context, s ...DeleteQuery) error {
	var model T
	q := c.DB.NewDelete().Model(&model)
	for i := range s {
		q.Apply(s[i])
	}

	_, err := q.Exec(ctx)
	return err
}

func Take(amount int) SelectQuery {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Limit(amount)
	}
}

func Skip(amount int) SelectQuery {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Offset(amount)
	}
}
